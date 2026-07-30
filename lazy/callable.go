package lazy

import (
	"time"

	"github.com/ecwid/control/retry"
)

type resolver[T any] interface {
	Resolve() (T, error)
}

type Deadline = retry.Timing

type Callable1[T any] func() (T, error)
type Callable func() error

func caller[E any](r resolver[E], request func(E) error) Callable {
	return func() error {
		node, err := r.Resolve()
		if err != nil {
			return err
		}
		return request(node)
	}
}

func caller1[E, T any](r resolver[E], request func(E) (T, error)) Callable1[T] {
	return func() (T, error) {
		e, err := r.Resolve()
		if err != nil {
			var zero T
			return zero, err
		}
		return request(e)
	}
}

func call(dl Deadline, function func() error) error {
	var (
		err      error
		retry    = 0
		start    = time.Now()
		deadline = dl.GetTimeout()
	)
	for {
		if retry > 0 && time.Since(start) >= deadline {
			break
		}
		dl.Before(retry)
		if err = function(); err == nil {
			return nil
		}
		retry++
	}
	return err
}

func (c Callable) Call(deadline Deadline) error {
	return call(deadline, c)
}

func (c Callable) MustCall(deadline Deadline) {
	if err := c.Call(deadline); err != nil {
		panic(err)
	}
}

func (c Callable1[T]) Call(deadline Deadline) (value T, err error) {
	err = call(deadline, func() error {
		value, err = c()
		return err
	})
	return
}

func (c Callable1[T]) MustCall(deadline Deadline) T {
	value, err := c.Call(deadline)
	if err != nil {
		panic(err)
	}
	return value
}
