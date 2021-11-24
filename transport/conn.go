package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ecwid/control/transport/observe"
	"github.com/gorilla/websocket"
)

var ErrReceiveTimeout = errors.New("websocket response timeout reached")

type RoundTripper interface {
	RoundTrip(context.Context, *Request) (*Response, error)
}

type Client struct {
	conn         *websocket.Conn
	seq          uint64
	recv         chan *Response
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
	c.seq++
	r, err := c.RoundTripper.RoundTrip(sessionCtx, &Request{
		ID:        c.seq,
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
	c.stdout("\033[1;36msend -> %s\033[0m", string(data))
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
		return nil, ErrReceiveTimeout
	}
}

func Connect(ctx context.Context, webSocketURL string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:       conn,
		seq:        0,
		recv:       make(chan *Response, 1),
		observable: observe.New(),
		Timeout:    time.Minute,
		Stderr:     os.Stderr,
	}
	c.Ctx, c.exit = context.WithCancel(ctx)
	c.RoundTripper = c
	go c.reader()
	return c, nil
}

func (c *Client) stdout(format string, v ...interface{}) {
	if c.Stdout != nil {
		_, _ = c.Stdout.Write([]byte(fmt.Sprintf(format, v...) + "\n"))
	}
}

func (c *Client) stderr(format string, v ...interface{}) {
	if c.Stderr != nil {
		_, _ = c.Stderr.Write([]byte(fmt.Sprintf(format, v...) + "\n"))
	}
}

func (c *Client) read() error {
	_, body, err := c.conn.ReadMessage()
	if err != nil {
		return err
	}
	var r = new(Response)
	if err = json.Unmarshal(body, r); err != nil {
		return err
	}
	if r.ID == c.seq {
		if r.isError() {
			c.stderr("\033[1;31mrecv_err <- %s\033[0m", string(body))
		} else {
			c.stdout("\033[1;34mrecv <- %s\033[0m", string(body))
		}
		c.recv <- r
	} else {
		c.stdout("\033[1;30mevent <- %s\033[0m", string(body))
		c.observable.Notify(r.SessionID, observe.Value{Method: r.Method, Params: r.Params})
	}
	return nil
}

func (c *Client) reader() {
	defer c.exit()
	for {
		if err := c.read(); err != nil {
			c.exitCode = err
			return
		}
	}
}
