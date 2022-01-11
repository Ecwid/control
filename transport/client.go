package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	defaultReadBufferSize  = 32 * 1024
	defaultWriteBufferSize = 32 * 1024
)

type Client struct {
	*Publisher
	conn      *websocket.Conn
	sendMutex sync.Mutex
	seq       uint64
	pending   map[uint64]*Call
	mutex     sync.Mutex
	close     bool
	Timeout   time.Duration
	Logger    io.Writer
}

func Dial(url string) (*Client, error) {
	var dialer = websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
		ReadBufferSize:   defaultReadBufferSize,
		WriteBufferSize:  defaultWriteBufferSize,
	}
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Publisher: NewPublisher(),
		conn:      conn,
		seq:       1,
		pending:   map[uint64]*Call{},
		Timeout:   time.Second * 60,
	}
	go client.reader()
	return client, nil
}

func (c *Client) Close() error {
	if err := c.Call("", "Browser.close", nil, nil); err != nil {
		return err
	}
	c.close = true
	return nil
}

func (c *Client) Call(sessionID, method string, args, value interface{}) error {
	var call = &Call{
		SessionID: sessionID,
		Method:    method,
		Args:      args,
		Reply:     make(chan Reply, 1),
	}
	if err := c.send(call); err != nil {
		return err
	}
	var timeout = time.NewTimer(c.Timeout)
	defer timeout.Stop()

	var r Reply
	select {
	case r = <-call.Reply:
		if r.Error != nil {
			return r.Error
		}
	case <-timeout.C:
		return CallTimeoutError{
			Call:    call,
			Timeout: c.Timeout,
		}
	}
	if value != nil {
		return json.Unmarshal(r.Result, value)
	}
	return nil
}

func (c *Client) send(call *Call) error {
	c.sendMutex.Lock()
	defer c.sendMutex.Unlock()

	c.mutex.Lock()
	if c.close {
		c.mutex.Unlock()
		return ErrShutdown
	}
	seq := c.seq
	c.seq++
	call.ID = seq
	c.pending[seq] = call
	c.mutex.Unlock()

	body, err := json.Marshal(call)
	if err != nil {
		return err
	}
	err = c.conn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return err
	}
	log(c.Logger, fmt.Sprintf("send -> %s", string(body)))
	return nil
}

func (c *Client) terminate(err error) {
	c.sendMutex.Lock()
	c.mutex.Lock()
	log(c.Logger, fmt.Sprintf("terminate -> %v", err))
	c.close = true
	for _, call := range c.pending {
		call.done(Reply{Error: &Error{Message: err.Error()}})
	}
	c.sendMutex.Unlock()
	c.mutex.Unlock()
}

func (c *Client) reader() {
	var (
		reply Reply
		body  []byte
		err   error
	)
	for {
		reply = Reply{}
		if _, body, err = c.conn.ReadMessage(); err != nil {
			break
		}
		log(c.Logger, fmt.Sprintf("recv <- %s", string(body)))
		if err = json.Unmarshal(body, &reply); err != nil {
			break
		}
		if reply.ID == 0 {
			c.Notify(reply.SessionID, Event{Method: reply.Method, Params: reply.Params})
		} else {
			c.mutex.Lock()
			call := c.pending[reply.ID]
			delete(c.pending, reply.ID)
			c.mutex.Unlock()
			if call == nil {
				err = errors.New("reading error body")
				break
			}
			call.done(reply)
		}
	}
	c.terminate(err)
}

func log(std io.Writer, v interface{}) {
	if std != nil {
		_, _ = std.Write([]byte(fmt.Sprint(v) + "\n"))
	}
}
