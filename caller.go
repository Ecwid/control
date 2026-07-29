package control

import (
	"context"
	"time"

	"github.com/ecwid/control/transport"
)

const DefaultEventBuffer = 1024

type CdpCaller struct {
	ctx       context.Context
	cancel    context.CancelCauseFunc
	timeout   time.Duration
	transport *transport.Transport
	sessionID string
}

func (c CdpCaller) Context() context.Context {
	return c.ctx
}

func (c CdpCaller) Cancel(err error) {
	if c.cancel != nil {
		c.cancel(err)
	}
}

func (c CdpCaller) Call(method string, send, recv any) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()
	return c.transport.Call(ctx, c.sessionID, method, send, recv)
}

func (c CdpCaller) Subscribe() (channel <-chan transport.Message, cancel func()) {
	return c.transport.Subscribe(c.sessionID, DefaultEventBuffer)
}

func (c CdpCaller) IsDone() bool {
	select {
	case <-c.ctx.Done():
		return true
	default:
		return false
	}
}
