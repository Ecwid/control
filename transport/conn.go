package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/transport/observe"
	"github.com/gorilla/websocket"
)

var AnsiColor = true

type RoundTripper interface {
	RoundTrip(context.Context, *Request) (*Response, error)
}

type Client struct {
	conn         *websocket.Conn
	sequence     uint64
	recv         chan *Response
	mx           *sync.Mutex
	Ctx          context.Context
	exit         func()
	exitCode     error
	observable   *observe.Observable
	RoundTripper RoundTripper
	Timeout      time.Duration
	Stdout       io.Writer
	Stderr       io.Writer
}

func (c *Client) Call(sessionCtx context.Context, sessionID, method string, args, value interface{}) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	r, err := c.RoundTripper.RoundTrip(sessionCtx, &Request{
		ID:        atomic.AddUint64(&c.sequence, 1),
		SessionID: sessionID,
		Method:    method,
		Params:    args,
	})
	if err != nil {
		return err
	}
	if value != nil {
		return json.Unmarshal(r.Result, value)
	}
	return nil
}

func (c *Client) Register(val observe.Observer) {
	c.observable.Register(val)
}

func (c *Client) Unregister(val observe.Observer) {
	c.observable.Unregister(val)
}

func (c *Client) RoundTrip(context context.Context, req *Request) (*Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	c.log(c.Stdout, cyan, fmt.Sprintf("send -> %s", string(data)))
	if err = c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, err
	}
	timeout := time.NewTimer(c.Timeout)
	defer timeout.Stop()
	select {
	case r := <-c.recv:
		if r.isError() {
			r.Error.Request = data
			return nil, r.Error
		}
		return r, nil
	case <-context.Done():
		if c.exitCode != nil {
			return nil, c.exitCode
		}
		return nil, context.Err()
	case <-timeout.C:
		return nil, ReceiveTimeoutError{
			Value:   data,
			Timeout: c.Timeout,
		}
	}
}

func Connect(ctx context.Context, webSocketURL string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:       conn,
		sequence:   0,
		recv:       make(chan *Response, 1),
		observable: observe.New(),
		Timeout:    time.Minute,
		Stderr:     os.Stderr,
		mx:         &sync.Mutex{},
	}
	c.Ctx, c.exit = context.WithCancel(ctx)
	c.RoundTripper = c
	go c.reader()
	return c, nil
}

func (c *Client) log(std io.Writer, color ansiColor, v interface{}) {
	if std != nil {
		val := fmt.Sprint(v)
		if AnsiColor {
			val = color + val + reset
		}
		_, _ = std.Write([]byte(val + "\n"))
	}
}

func (c *Client) readNext() error {
	_, body, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	var r = new(Response)
	if err = json.Unmarshal(body, r); err != nil {
		return err
	}
	if r.ID == 0 {
		c.log(c.Stdout, black, fmt.Sprintf("event <- %s", string(body)))
		c.observable.Notify(r.SessionID, observe.Value{Method: r.Method, Params: r.Params})
	} else if r.ID == atomic.LoadUint64(&c.sequence) {
		if r.isError() {
			c.log(c.Stderr, red, fmt.Sprintf("recv-error <- %s", string(body)))
		} else {
			c.log(c.Stdout, blue, fmt.Sprintf("recv <- %s", string(body)))
		}
		c.recv <- r
	}
	return nil
}

func (c *Client) reader() {
	defer c.exit()
	for {
		if err := c.readNext(); err != nil {
			c.exitCode = err
			return
		}
	}
}
