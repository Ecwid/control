package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
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
}

func Dial(url string) (*Client, error) {
	var dialer = websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
		Proxy:            http.ProxyFromEnvironment,
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
	go client.reading()
	return client, nil
}

func (c *Client) Close() error {
	if err := c.Call("", "Browser.close", nil, nil); err != nil {
		return err
	}
	c.conn.Close()
	c.terminate(ErrShutdown)
	return nil
}

func (c *Client) Call(sessionID, method string, args, value interface{}) error {
	var call = &Call{
		SessionID: sessionID,
		Method:    method,
		Args:      args,
		reply:     make(chan Reply, 1),
	}
	if err := c.send(call); err != nil {
		return err
	}
	var timeout = time.NewTimer(c.Timeout)
	defer timeout.Stop()

	var r Reply
	select {
	case r = <-call.reply:
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

	if err := c.conn.WriteJSON(call); err != nil {
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return err
	}
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

func (c *Client) nextReply() error {
	reply := Reply{}
	if err := c.conn.ReadJSON(&reply); err != nil {
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

func (c *Client) reading() {
	var err error
	for ; err == nil; err = c.nextReply() {
	}
	c.terminate(err)
}
