package control

import (
	"fmt"
	"time"

	"github.com/ecwid/control/transport/observe"
)

type ErrConditionTimeout struct {
	name     string
	deadline time.Duration
}

func (e ErrConditionTimeout) Error() string {
	return fmt.Sprintf("condition timeout reached out after %s", e.deadline)
}

type Condition struct {
	session   *Session
	condition func(value observe.Value) (bool, error)
	deadline  time.Duration
}

func NewCondition(session *Session, deadline time.Duration, v func(observe.Value) (bool, error)) Condition {
	return Condition{
		session:   session,
		condition: v,
		deadline:  deadline,
	}
}

func (con Condition) Do(initialAction func() error) error {
	var (
		isDone      = make(chan bool, 1)
		internalErr = make(chan error, 1)
	)
	defer close(isDone)
	defer close(internalErr)
	unsubscribe := con.session.Subscribe("*", false, func(e observe.Value) {
		v, err := con.condition(e)
		if err != nil {
			select {
			case internalErr <- err:
			default:
			}
			return
		}
		if v {
			select {
			case isDone <- true:
			default:
			}
		}
	})
	defer unsubscribe()
	if initialAction != nil {
		if err := initialAction(); err != nil {
			return err
		}
	}
	select {
	case <-isDone:
		return nil
	case err := <-internalErr:
		return err
	case <-con.session.exited:
		return con.session.exitCode
	case <-time.After(con.deadline):
		return ErrConditionTimeout{deadline: con.deadline}
	}
}
