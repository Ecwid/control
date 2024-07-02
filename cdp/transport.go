package cdp

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var DefaultDialer = websocket.Dialer{
	ReadBufferSize:   8192,
	WriteBufferSize:  8192,
	HandshakeTimeout: 45 * time.Second,
	Proxy:            http.ProxyFromEnvironment,
}

var ErrGracefullyClosed = errors.New("gracefully closed")

type Transport struct {
	context context.Context
	cancel  func(error)
	conn    *websocket.Conn
	seq     uint64
	pending map[uint64]*promise[Response]
	mutex   sync.Mutex
	broker  broker
	logger  *slog.Logger
}

func DefaultDial(context context.Context, url string, logger *slog.Logger) (*Transport, error) {
	return Dial(context, DefaultDialer, url, logger)
}

func Dial(parent context.Context, dialer websocket.Dialer, url string, logger *slog.Logger) (*Transport, error) {
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancelCause(parent)
	transport := &Transport{
		context: ctx,
		cancel:  cancel,
		conn:    conn,
		seq:     1,
		broker:  makeBroker(),
		pending: make(map[uint64]*promise[Response]),
		logger:  logger,
	}
	go transport.broker.run()
	go func() {
		var readerr error
		for ; readerr == nil; readerr = transport.read() {
		}
		transport.cancel(readerr)
		transport.gracefullyClose()
	}()
	return transport, nil
}

func (t *Transport) Log(level slog.Level, msg string, args ...any) {
	if t.logger != nil {
		t.logger.Log(t.context, level, msg, args...)
	}
}

func (t *Transport) Context() context.Context {
	return t.context
}

func (t *Transport) Close() error {
	select {
	case <-t.context.Done():
		return context.Cause(t.context)
	default:
		_, err := t.Send(&Request{Method: "Browser.close"}).Get(t.context)
		if err != nil {
			return err
		}
		t.cancel(ErrGracefullyClosed)
		return t.conn.Close()
	}
}

func (t *Transport) gracefullyClose() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.broker.Cancel()
	err := context.Cause(t.context)
	for key, value := range t.pending {
		value.reject(err)
		delete(t.pending, key)
	}
}

func (t *Transport) Subscribe(sessionID string) (chan Message, func()) {
	channel := t.broker.subscribe(sessionID)
	return channel, func() {
		if channel != nil {
			t.broker.unsubscribe(channel)
		}
	}
}

func (t *Transport) Send(request *Request) Future[Response] {
	var promise = &promise[Response]{fulfilled: make(chan struct{}, 1)}
	select {
	case <-t.context.Done():
		promise.reject(context.Cause(t.context))
		return promise
	default:
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	seq := t.seq
	t.seq++
	t.pending[seq] = promise
	request.ID = seq
	t.Log(slog.LevelDebug, "send ->", "request", request.String())

	if err := t.conn.WriteJSON(request); err != nil {
		delete(t.pending, seq)
		promise.reject(err)
	}
	return promise
}

func (t *Transport) read() (err error) {
	var response = Response{}
	if err = t.conn.ReadJSON(&response); err != nil {
		return err
	}
	t.Log(slog.LevelDebug, "recv <-", "response", response.String())

	if response.ID == 0 && response.Message != nil {
		t.broker.publish(*response.Message)
		return nil
	}

	t.mutex.Lock()
	value, ok := t.pending[response.ID]
	delete(t.pending, response.ID)
	t.mutex.Unlock()

	if !ok {
		return errors.New("unexpected response " + response.String())
	}
	if response.Error != nil {
		value.reject(response.Error)
		return nil
	}
	value.resolve(response)
	return nil
}
