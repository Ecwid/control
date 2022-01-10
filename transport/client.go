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

type RoundTripper interface {
	RoundTrip(*Call) (*Reply, error)
}

func (call *Call) done(r Reply) {
	select {
	case call.Reply <- r:
	default:
		// We don't want to block here.
	}
}

type Client struct {
	*Publisher
	conn      *websocket.Conn
	sendMutex sync.Mutex
	seq       uint64
	pending   map[uint64]*Call
	mutex     sync.Mutex
	closing   bool // user has called Close
	shutdown  bool // server has told us to stop
	Timeout   time.Duration
	Logger    io.Writer
}

func Dial(url string) (*Client, error) {
	var dialer = websocket.Dialer{
		ReadBufferSize:  defaultReadBufferSize,
		WriteBufferSize: defaultWriteBufferSize,
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
	go client.input()
	return client, nil
}

func (c *Client) Close() error {
	if err := c.Call("", "Browser.close", nil, nil); err != nil {
		return err
	}
	c.closing = true
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
	if c.shutdown || c.closing {
		c.mutex.Unlock()
		return ErrShutdown
	}
	seq := c.seq
	c.seq++
	call.ID = seq
	c.pending[seq] = call
	c.mutex.Unlock()

	if err := c.write(call); err != nil {
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return err
	}
	return nil
}

func (c *Client) write(call *Call) (err error) {
	var body []byte
	body, err = json.Marshal(call)
	if err != nil {
		return err
	}
	log(c.Logger, fmt.Sprintf("send -> %s", string(body)))
	return c.conn.WriteMessage(websocket.TextMessage, body)
}

func (c *Client) read(r *Reply) (err error) {
	var body []byte
	if _, body, err = c.conn.ReadMessage(); err != nil {
		return err
	}
	log(c.Logger, fmt.Sprintf("recv <- %s", string(body)))
	return json.Unmarshal(body, r)
}

func (c *Client) input() {
	var (
		res Reply
		err error
	)
	for {
		res = Reply{}
		if err = c.read(&res); err != nil {
			break
		}
		id := res.ID
		if id == 0 {
			c.Notify(res.SessionID, Event{Method: res.Method, Params: res.Params})
		} else {
			c.mutex.Lock()
			call := c.pending[id]
			delete(c.pending, id)
			c.mutex.Unlock()
			if call == nil {
				err = errors.New("reading error body")
				break
			}
			call.done(res)
		}
	}
	// Terminate pending calls.
	c.sendMutex.Lock()
	c.mutex.Lock()
	c.shutdown = true
	for _, call := range c.pending {
		call.done(Reply{Error: &Error{Message: err.Error()}})
	}
	c.mutex.Unlock()
	c.sendMutex.Unlock()
}

func log(std io.Writer, v interface{}) {
	if std != nil {
		_, _ = std.Write([]byte(fmt.Sprint(v) + "\n"))
	}
}
