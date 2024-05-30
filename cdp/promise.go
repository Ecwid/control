package cdp

import (
	"context"
	"errors"
	"sync"
)

var ErrPromiseCanceled = errors.New("promise canceled")

type Future[T any] interface {
	Get(context.Context) (T, error)
	Cancel()
}

func NewPromise[T any](executor func(resolve func(T), reject func(error)), finally func()) Future[T] {
	value := &promise[T]{
		finally:   finally,
		fulfilled: make(chan struct{}, 1),
	}
	go executor(value.resolve, value.reject)
	return value
}

type promise[T any] struct {
	once      sync.Once
	fulfilled chan struct{}
	value     T
	err       error
	finally   func()
}

func (u *promise[T]) Get(parent context.Context) (T, error) {
	defer u.Cancel()
	select {
	case <-parent.Done():
		return u.value, context.Cause(parent)
	case <-u.fulfilled:
		return u.value, u.err
	}
}

func (u *promise[T]) Cancel() {
	u.reject(ErrPromiseCanceled)
}

func (u *promise[T]) resolve(value T) {
	u.once.Do(func() {
		u.value = value
		close(u.fulfilled)
		if u.finally != nil {
			u.finally()
		}
	})
}

func (u *promise[T]) reject(err error) {
	u.once.Do(func() {
		u.err = err
		close(u.fulfilled)
		if u.finally != nil {
			u.finally()
		}
	})
}
