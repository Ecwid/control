package control

import (
	"context"
	"sync"
	"time"

	"github.com/ecwid/control/transport/observe"
)

type Promise struct {
	once        *sync.Once
	ctx         context.Context
	done        chan struct{}
	err         chan error
	unsubscribe func()
}

func NewPromise(ctx context.Context) *Promise {
	return &Promise{
		ctx:  ctx,
		once: &sync.Once{},
		done: make(chan struct{}, 1),
		err:  make(chan error, 1),
	}
}

func (u Promise) Resolve() {
	select {
	case u.done <- struct{}{}:
	default:
	}
}

func (u Promise) Reject(err error) {
	select {
	case u.err <- err:
	default:
	}
	return
}

func (u Promise) Stop() {
	u.once.Do(func() {
		close(u.done)
		close(u.err)
		if u.unsubscribe != nil {
			u.unsubscribe()
		}
	})
}

func (u Promise) WaitWithTimeout(timeout time.Duration) error {
	defer u.Stop()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-u.done:
		return nil
	case err := <-u.err:
		return err
	case <-u.ctx.Done():
		return u.ctx.Err()
	case <-timer.C:
		return WaitTimeoutError{timeout: timeout}
	}
}

func (s Session) NewEventCondition(method string, condition func(value observe.Value) (bool, error)) *Promise {
	u := NewPromise(s.Ctx)
	u.unsubscribe = s.Subscribe(method, false, func(e observe.Value) {
		v, err := condition(e)
		if err != nil {
			u.Reject(err)
		}
		if v {
			u.Resolve()
		}
	})
	return u
}
