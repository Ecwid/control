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
	closed    bool
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
	c.terminate(ErrShutdown)
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

	if c.closed {
		return ErrShutdown
	}

	c.mutex.Lock()
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
	c.log(fmt.Sprintf("send -> %s", string(body)))
	return nil
}

func (c *Client) terminate(err error) {
	c.sendMutex.Lock()
	c.mutex.Lock()
	c.closed = true
	for _, call := range c.pending {
		call.done(Reply{Error: &Error{Message: err.Error()}})
	}
	c.mutex.Unlock()
	c.sendMutex.Unlock()
}

func (c *Client) read() error {
	reply := Reply{}
	_, body, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	c.log(fmt.Sprintf("recv <- %s", string(body)))
	if err = json.Unmarshal(body, &reply); err != nil {
		return err
	}
	if reply.ID == 0 {
		c.Notify(reply.SessionID, Event{Method: reply.Method, Params: reply.Params})
	} else {
		c.mutex.Lock()
		call := c.pending[reply.ID]
		delete(c.pending, reply.ID)
		c.mutex.Unlock()
		if call == nil {
			return errors.New("reading error body")
		}
		call.done(reply)
	}
	return nil
}

func (c *Client) reader() {
	var err error
	for {
		if err = c.read(); err != nil {
			c.terminate(err)
			return
		}
	}
}

func (c *Client) log(v interface{}) {
	if c.Logger != nil {
		_, _ = c.Logger.Write([]byte(fmt.Sprint(v) + "\n"))
	}
}
