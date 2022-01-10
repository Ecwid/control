package control

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/transport"
)

const (
	promisePending  = 0
	promiseResolved = 1
	promiseRejected = 2
)

// a read-only view of promise
type Future struct {
	promise *promise
}

type promise struct {
	once       *sync.Once
	context    context.Context
	done       chan interface{}
	err        chan error
	cancelFunc func()
	state      *int32
}

func (u promise) resolve(val interface{}) {
	select {
	case u.done <- val:
		atomic.StoreInt32(u.state, promiseResolved)
	default:
	}
}

func (u promise) reject(err error) {
	select {
	case u.err <- err:
		atomic.StoreInt32(u.state, promiseRejected)
	default:
	}
}

func (u promise) isPending() bool {
	return atomic.LoadInt32(u.state) == promisePending
}

func (u promise) cancel() {
	u.once.Do(func() {
		close(u.done)
		close(u.err)
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
	case val := <-u.promise.done:
		u.promise.done <- val
		return val, nil
	case err := <-u.promise.err:
		u.promise.err <- err
		return nil, err
	case <-u.promise.context.Done():
		return nil, u.promise.context.Err()
	case <-timer.C:
		return nil, FutureTimeoutError{timeout: timeout}
	}
}

func (s Session) Observe(method string, condition func(transport.Event, func(interface{}), func(error))) Future {
	var state int32 = promisePending
	u := &promise{
		context: s.context,
		state:   &state,
		once:    &sync.Once{},
		done:    make(chan interface{}, 1),
		err:     make(chan error, 1),
	}
	u.cancelFunc = s.Subscribe(method, func(e transport.Event) {
		if u.isPending() {
			condition(e, u.resolve, u.reject)
		}
	})
	return Future{u}
}
