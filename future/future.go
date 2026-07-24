package future

import (
	"context"
	"errors"
	"sync"
)

var ErrFutureCanceled = errors.New("future canceled")

type Future[T any] interface {
	Get(context.Context) (T, error)
	Cancel()
}

func Execute[T any](executor func(resolve func(T), reject func(error), canceled <-chan struct{})) Future[T] {
	value := &future[T]{
		fulfilled: make(chan struct{}),
	}
	go executor(value.resolve, value.reject, value.fulfilled)
	return value
}

type future[T any] struct {
	once      sync.Once
	fulfilled chan struct{}
	value     T
	err       error
}

func (u *future[T]) Get(ctx context.Context) (T, error) {
	defer u.Cancel()
	select {
	case <-ctx.Done():
		return u.value, context.Cause(ctx)
	case <-u.fulfilled:
		return u.value, u.err
	}
}

func (u *future[T]) Cancel() {
	u.reject(ErrFutureCanceled)
}

func (u *future[T]) resolve(value T) {
	u.done(value, nil)
}

func (u *future[T]) reject(err error) {
	var z T
	u.done(z, err)
}

func (u *future[T]) done(value T, err error) {
	u.once.Do(func() {
		u.value = value
		u.err = err
		close(u.fulfilled)
	})
}
