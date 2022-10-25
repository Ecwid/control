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
	cancelled
)

// a read-only view of promise
type Future struct {
	promise *promise
}

type promise struct {
	onDefer    sync.Once
	mutex      sync.Mutex
	context    context.Context
	cancelFunc func()
	state      int32
	done       chan struct{}
	value      interface{}
	error      error
}

func (u *promise) resolve(val interface{}) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.state == pending {
		u.state = resolved
		u.value = val
		u.done <- struct{}{}
	}
}

func (u *promise) reject(err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.state == pending {
		u.state = rejected
		u.error = err
		u.done <- struct{}{}
	}
}

func (u *promise) isPending() bool {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return u.state == pending
}

func (u *promise) cancel() {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if u.state == pending {
		u.state = cancelled
	}
	u.onDefer.Do(func() {
		close(u.done)
		if u.cancelFunc != nil {
			u.cancelFunc()
		}
	})
}

func (u Future) Cancel() {
	u.promise.cancel()
}

func (u Future) Get(timeout time.Duration) (interface{}, error) {
	defer u.Cancel()
	var timer = time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-u.promise.done:
		return u.promise.value, u.promise.error
	case <-u.promise.context.Done():
		return nil, u.promise.context.Err()
	case <-timer.C:
		return nil, FutureTimeoutError{timeout: timeout}
	}
}

func (s Session) Observe(method string, condition func(transport.Event, func(interface{}), func(error))) Future {
	u := &promise{
		context: s.context,
		state:   pending,
		onDefer: sync.Once{},
		mutex:   sync.Mutex{},
		done:    make(chan struct{}, 1),
		value:   nil,
		error:   nil,
	}

	u.cancelFunc = s.Subscribe(method, func(e transport.Event) {
		if u.isPending() {
			condition(e, u.resolve, u.reject)
		}
	})
	return Future{u}
}
