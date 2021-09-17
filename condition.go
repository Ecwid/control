package control

import (
	"time"

	"github.com/ecwid/control/transport/observe"
)

type Condition struct {
	session     *Session
	isDone      chan struct{}
	error       chan error
	unsubscribe func()
}

func (s *Session) NewCondition(condition func(value observe.Value) (bool, error)) *Condition {
	u := &Condition{
		session: s,
		isDone:  make(chan struct{}, 1),
		error:   make(chan error, 1),
	}
	u.unsubscribe = s.Subscribe("*", false, func(e observe.Value) {
		v, err := condition(e)
		if err != nil {
			select {
			case u.error <- err:
			default:
			}
			return
		}
		if v {
			select {
			case u.isDone <- struct{}{}:
			default:
			}
		}
	})
	return u
}

func (u Condition) Wait(initial func() error) error {
	return u.WaitWithTimeout(initial, u.session.Timeout)
}

func (u Condition) WaitWithTimeout(initial func() error, timeout time.Duration) error {
	defer close(u.isDone)
	defer close(u.error)
	defer u.unsubscribe()

	if initial != nil {
		if err := initial(); err != nil {
			return err
		}
	}

	select {
	case <-u.isDone:
		return nil
	case err := <-u.error:
		return err
	case <-u.session.Ctx.Done():
		return u.session.ExitCode()
	case <-time.After(timeout):
		return WaitTimeoutError{timeout: timeout}
	}
}
