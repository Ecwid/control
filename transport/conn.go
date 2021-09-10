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
	RoundTrip(*Request) (*Response, error)
}

type T interface {
	Call(sessionID, method string, args, value interface{}) error
	Add(val observe.Observer)
	Remove(val observe.Observer)
}

type WS struct {
	ctx          context.Context
	conn         *websocket.Conn
	seq          uint64
	recv         chan *Response
	fatal        chan error
	observable   *observe.Observable
	RoundTripper RoundTripper
	Timeout      time.Duration
	Stdout       io.Writer
	Stderr       io.Writer
}

func (c *WS) Call(sessionID, method string, args, value interface{}) error {
	c.seq++
	r, err := c.RoundTripper.RoundTrip(&Request{
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

func (c *WS) Add(val observe.Observer) {
	c.observable.Add(val)
}

func (c *WS) Remove(val observe.Observer) {
	c.observable.Remove(val)
}

func (c *WS) RoundTrip(req *Request) (*Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	c.stdout("\033[1;36msend -> %s\033[0m", string(data))
	if err = c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, err
	}
	select {
	case r := <-c.recv:
		if r.isError() {
			r.Error.Request = data
			return nil, r.Error
		}
		return r, nil
	case err = <-c.fatal:
		return nil, err
	case <-time.After(c.Timeout):
		return nil, ErrReceiveTimeout
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

func Connect(ctx context.Context, webSocketURL string) (T, error) {
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		return nil, err
	}
	c := &WS{
		ctx:        ctx,
		conn:       conn,
		seq:        0,
		recv:       make(chan *Response, 1),
		fatal:      make(chan error, 1),
		observable: observe.New(),
		Timeout:    time.Minute,
		Stderr:     os.Stderr,
	}
	c.RoundTripper = c
	go c.reader()
	return c, nil
}

func (c *WS) stdout(format string, v ...interface{}) {
	if c.Stdout != nil {
		_, _ = c.Stdout.Write([]byte(fmt.Sprintf(format, v...) + "\n"))
	}
}

func (c *WS) stderr(format string, v ...interface{}) {
	if c.Stderr != nil {
		_, _ = c.Stderr.Write([]byte(fmt.Sprintf(format, v...) + "\n"))
	}
}

func (c *WS) stop(err error) {
	if err != nil {
		c.stderr("\033[1;31m%s\033[0m", err.Error())
	}
	c.fatal <- err
}

func (c *WS) reader() {
	for {
		_, body, err := c.conn.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				err = nil // do nothing, connection was closed gracefully
			}
			c.stop(err)
			return
		}
		var r = new(Response)
		if err = json.Unmarshal(body, r); err != nil {
			c.stop(err)
			return
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
	}
}
