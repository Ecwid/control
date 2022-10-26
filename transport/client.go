package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	*Publisher
	conn    *websocket.Conn
	seq     uint64
	queue   map[uint64]*Request
	queueMu sync.Mutex
	sendMu  sync.Mutex
	context context.Context
	Timeout time.Duration
	err     error
	cancel  func()
}

func Dial(ctx context.Context, url string) (*Client, error) {
	var dialer = websocket.Dialer{
		ReadBufferSize:   8192,
		WriteBufferSize:  8192,
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
		queue:     map[uint64]*Request{},
		Timeout:   time.Second * 60,
	}
	client.context, client.cancel = context.WithCancel(ctx)
	go client.reading()
	return client, nil
}

func (c *Client) Context() context.Context {
	return c.context
}

func (c *Client) Close() error {
	if err := c.Call("", "Browser.close", nil, nil); err != nil {
		return err
	}
	_ = c.conn.Close()
	c.finalize(errors.New("connection is shut down"))
	return nil
}

func (c *Client) Call(sessionID, method string, args, value interface{}) error {
	var request = &Request{
		SessionID: sessionID,
		Method:    method,
		Args:      args,
		response:  make(chan Response, 1),
	}
	if err := c.send(request); err != nil {
		return err
	}
	var ctx, cancel = context.WithTimeout(c.context, c.Timeout)
	defer cancel()

	var r Response
	select {
	case r = <-request.response:
		if r.Error != nil {
			return r.Error
		}
	case <-ctx.Done():
		return DeadlineExceededError{Request: request, Timeout: c.Timeout}
	}
	if value != nil {
		return json.Unmarshal(r.Result, value)
	}
	return nil
}

func (c *Client) send(request *Request) error {
	c.sendMu.Lock()
	defer c.sendMu.Unlock()

	select {
	case <-c.context.Done():
		return c.err
	default:
	}

	c.queueMu.Lock()
	seq := c.seq
	c.seq++
	request.ID = seq
	c.queue[seq] = request
	c.queueMu.Unlock()

	if err := c.conn.WriteJSON(request); err != nil {
		c.queueMu.Lock()
		delete(c.queue, seq)
		c.queueMu.Unlock()
		return err
	}
	return nil
}

func (c *Client) finalize(err error) {
	c.sendMu.Lock()
	c.queueMu.Lock()
	defer c.queueMu.Unlock()
	defer c.sendMu.Unlock()
	c.err = err
	c.cancel()
	for _, request := range c.queue {
		_ = request.received(Response{Error: &Error{Message: err.Error()}})
	}
}

func (c *Client) read() error {
	response := Response{}
	if err := c.conn.ReadJSON(&response); err != nil {
		return err
	}
	if response.ID == 0 {
		return c.Notify(response.SessionID, Event{
			Method: response.Method,
			Params: response.Params,
		})
	}
	c.queueMu.Lock()
	request := c.queue[response.ID]
	delete(c.queue, response.ID)
	c.queueMu.Unlock()
	if request == nil {
		return errors.New("no request for response")
	}
	return request.received(response)
}

func (c *Client) reading() {
	var err error
	for ; err == nil; err = c.read() {
	}
	c.finalize(err)
}
