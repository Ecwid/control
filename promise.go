package control

import (
	"context"
	"sync"
	"time"

	"github.com/ecwid/control/transport"
)

const (
	pending = iota + 0
	resolved
	rejected
)

type Future struct {
	promise *promise
}

type promise struct {
	mutex      sync.Mutex
	context    context.Context
	unregister func()
	cancel     func()
	state      int32
	value      interface{}
	error      error
}

func (u *promise) resolve(val interface{}) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.state == pending {
		u.state = resolved
		u.value = val
		u.cancel()
	}
}

func (u *promise) reject(err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.state == pending {
		u.state = rejected
		u.error = err
		u.cancel()
	}
}

func (u *promise) isPending() bool {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return u.state == pending
}

func (u Future) Cancel() {
	u.promise.cancel()
	u.promise.unregister()
}

func (u Future) IsPending() bool {
	return u.promise.isPending()
}

func (u Future) Get(timeout time.Duration) (interface{}, error) {
	var ctx, cancel = context.WithTimeout(u.promise.context, timeout)
	defer cancel()
	<-ctx.Done()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, FutureTimeoutError{timeout: timeout}
	}
	return u.promise.value, u.promise.error
}

func (s Session) Observe(method string, condition func(transport.Event, func(interface{}), func(error))) Future {
	u := &promise{
		state: pending,
		mutex: sync.Mutex{},
		value: nil,
		error: nil,
	}
	u.context, u.cancel = context.WithCancel(s.context)
	u.unregister = s.Subscribe(method, func(e transport.Event) error {
		if u.isPending() {
			condition(e, u.resolve, u.reject)
		}
		return nil
	})
	return Future{u}
}
