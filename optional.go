package control

import "fmt"

type Optional[T any] struct {
	value T
	err   error
}

func (op Optional[T]) Unwrap() (T, error) {
	return op.value, op.err
}

func (op Optional[T]) Err() error {
	return op.err
}

func (op Optional[T]) MustGetValue() T {
	if op.err != nil {
		panic(op.err)
	}
	return op.value
}

func (op Optional[T]) Then(f func(T) error) error {
	if op.err == nil {
		return f(op.value)
	}
	return op.err
}

func (op Optional[T]) Catch(f func(error) error) error {
	if op.err != nil {
		return f(op.err)
	}
	return nil
}

func (op Optional[T]) IfPresent(f func(T)) {
	if op.err == nil {
		f(op.value)
	}
}

func conv[T any](value any, err error) Optional[T] {
	if err != nil {
		return Optional[T]{err: err}
	}
	if value != nil {
		if v, ok := value.(T); ok {
			return Optional[T]{value: v}
		}
	}
	var zero T
	return Optional[T]{err: fmt.Errorf("interface conversion failed: got %T, want %T", value, zero)}
}
