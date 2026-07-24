package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

var DefaultDialer = websocket.Dialer{
	ReadBufferSize:   8192,
	WriteBufferSize:  8192,
	HandshakeTimeout: 45 * time.Second,
	Proxy:            http.ProxyFromEnvironment,
}

var ErrClosed = errors.New("cdp: closed")

const DefaultEventBuffer = 1024

type Transport struct {
	// Root lifecycle context for all transport goroutines and operations.
	ctx context.Context
	// Cancels ctx with a cause and initiates transport shutdown.
	cancel context.CancelCauseFunc
	// Active websocket connection to the browser endpoint.
	conn *websocket.Conn
	// Optional structured logger used by send/receive paths.
	log *slog.Logger

	// Single-writer queue for outbound websocket writes.
	writeCh chan writeRequest

	// Protects pending map against concurrent read/write loop access.
	pendingMu sync.Mutex
	// In-flight request waiters keyed by CDP request ID.
	pending map[uint64]chan callResult
	// Monotonic generator for request IDs.
	nextID atomic.Uint64

	// Fan-out hub for event subscribers.
	subs *subscriberHub

	// Ensures close/fail sequence runs exactly once.
	closeOnce sync.Once
}

type writeRequest struct {
	request Request
	ack     chan error
}

type callResult struct {
	response Response
	err      error
}

func DefaultDial(parent context.Context, url string, logger *slog.Logger) (*Transport, error) {
	return Dial(parent, DefaultDialer, url, logger)
}

func Dial(parent context.Context, dialer websocket.Dialer, url string, logger *slog.Logger) (*Transport, error) {
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancelCause(parent)
	c := &Transport{
		ctx:     ctx,
		cancel:  cancel,
		conn:    conn,
		log:     logger,
		writeCh: make(chan writeRequest),
		pending: make(map[uint64]chan callResult),
		subs:    newSubscriberHub(),
	}
	c.nextID.Store(1)

	go c.readLoop()
	go c.writeLoop()

	return c, nil
}

func (c *Transport) Context() context.Context {
	return c.ctx
}

func (c *Transport) Close() error {
	c.closeOnce.Do(func() {
		c.cancel(ErrClosed)
		_ = c.conn.Close()
	})
	return context.Cause(c.ctx)
}

func (c *Transport) Shutdown(ctx context.Context) error {
	_, err := c.Do(ctx, Request{Method: "Browser.close"})
	closeErr := c.Close()
	return errors.Join(err, closeErr)
}

func (c *Transport) Do(ctx context.Context, req Request) (Response, error) {
	var zero Response

	if err := context.Cause(c.ctx); err != nil {
		return zero, err
	}

	id := c.nextID.Add(1) - 1
	req.ID = id
	resultCh := make(chan callResult, 1)

	c.pendingMu.Lock()
	c.pending[id] = resultCh
	c.pendingMu.Unlock()

	defer func() {
		c.pendingMu.Lock()
		delete(c.pending, id)
		c.pendingMu.Unlock()
	}()

	ack := make(chan error, 1)
	wr := writeRequest{request: req, ack: ack}

	select {
	case <-ctx.Done():
		return zero, context.Cause(ctx)
	case <-c.ctx.Done():
		return zero, context.Cause(c.ctx)
	case c.writeCh <- wr:
	}

	select {
	case err := <-ack:
		if err != nil {
			return zero, err
		}
	case <-ctx.Done():
		return zero, context.Cause(ctx)
	case <-c.ctx.Done():
		return zero, context.Cause(c.ctx)
	}

	select {
	case result := <-resultCh:
		if result.err != nil {
			return zero, result.err
		}
		if result.response.Error != nil {
			return zero, result.response.Error
		}
		return result.response, nil
	case <-ctx.Done():
		return zero, context.Cause(ctx)
	case <-c.ctx.Done():
		return zero, context.Cause(c.ctx)
	}
}

func (c *Transport) Call(ctx context.Context, sessionID, method string, params any, out any) error {
	res, err := c.Do(ctx, Request{
		SessionID: sessionID,
		Method:    method,
		Params:    params,
	})
	if err != nil {
		return err
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(res.Result, out)
}

func (c *Transport) Subscribe(sessionID string, buffer int) (<-chan Message, func()) {
	if buffer <= 0 {
		buffer = DefaultEventBuffer
	}
	return c.subs.subscribe(sessionID, buffer)
}

func (c *Transport) Log(level slog.Level, msg string, args ...any) {
	if c.log != nil {
		c.log.Log(c.ctx, level, msg, args...)
	}
}

func (c *Transport) writeLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case wr := <-c.writeCh:
			err := c.conn.WriteJSON(wr.request)
			if err != nil {
				wr.ack <- err
				c.fail(err)
				continue
			}
			wr.ack <- nil
			c.Log(slog.LevelDebug, "send ->", "request", wr.request.String())
		}
	}
}

func (c *Transport) readLoop() {
	for {
		var response Response
		if err := c.conn.ReadJSON(&response); err != nil {
			c.fail(err)
			return
		}

		c.Log(slog.LevelDebug, "recv <-", "response", response.String())

		if response.ID == 0 {
			if response.Message != nil {
				c.subs.publish(*response.Message)
			}
			continue
		}

		c.pendingMu.Lock()
		resultCh, ok := c.pending[response.ID]
		if ok {
			delete(c.pending, response.ID)
		}
		c.pendingMu.Unlock()

		if !ok {
			c.Log(slog.LevelWarn, "unexpected response", "response", response.String())
			continue
		}

		resultCh <- callResult{response: response}
	}
}

func (c *Transport) fail(err error) {
	c.closeOnce.Do(func() {
		c.cancel(err)
		_ = c.conn.Close()

		c.pendingMu.Lock()
		for id, resultCh := range c.pending {
			resultCh <- callResult{err: err}
			delete(c.pending, id)
		}
		c.pendingMu.Unlock()

		c.subs.closeAll()
	})
}

type subscriberHub struct {
	mu     sync.RWMutex
	nextID uint64
	subs   map[uint64]subscriber
}

type subscriber struct {
	sessionID string
	ch        chan Message
}

func newSubscriberHub() *subscriberHub {
	return &subscriberHub{subs: make(map[uint64]subscriber)}
}

func (h *subscriberHub) subscribe(sessionID string, buffer int) (<-chan Message, func()) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.nextID++
	id := h.nextID
	ch := make(chan Message, buffer)
	h.subs[id] = subscriber{sessionID: sessionID, ch: ch}

	unsubscribe := func() {
		h.mu.Lock()
		defer h.mu.Unlock()
		sub, ok := h.subs[id]
		if !ok {
			return
		}
		delete(h.subs, id)
		close(sub.ch)
	}

	return ch, unsubscribe
}

func (h *subscriberHub) publish(msg Message) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, sub := range h.subs {
		if msg.SessionID != "" && sub.sessionID != "" && msg.SessionID != sub.sessionID {
			continue
		}
		select {
		case sub.ch <- msg:
		default:
			// Drop when subscriber is slow; transport stays healthy.
		}
	}
}

func (h *subscriberHub) closeAll() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for id, sub := range h.subs {
		close(sub.ch)
		delete(h.subs, id)
	}
}
