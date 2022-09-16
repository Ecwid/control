package control

import (
	"errors"
	"fmt"
	"time"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/target"
)

var (
	ErrNodeIsNotVisible          = errors.New("node is not visible")
	ErrNodeIsOutOfViewport       = errors.New("node is out of viewport")
	ErrAlreadyNavigated          = errors.New("page already navigated to this address - nothing done")
	ErrTargetDestroyed           = errors.New("this session was destroyed")
	ErrDetachedFromTarget        = errors.New("detached from target")
	ErrClickTimeout              = errors.New("no click registered")
	ErrExecutionContextDestroyed = errors.New("execution context was destroyed")
)

type ErrTargetCrashed target.TargetCrashed

func (e ErrTargetCrashed) Error() string {
	return fmt.Sprintf("TargetID = %s, ErrorCode = %d, Status = %s", e.TargetId, e.ErrorCode, e.Status)
}

type NoSuchElementError struct {
	Selector string
}

func (n NoSuchElementError) Error() string {
	return fmt.Sprintf("no such element `%s`", n.Selector)
}

type NoSuchFrameError struct {
	id common.FrameId
}

func (n NoSuchFrameError) Error() string {
	return fmt.Sprintf("no such frame `%s`", n.id)
}

type RemoteObjectCastError struct {
	object primitiveRemoteObject
	cast   string
}

func (r RemoteObjectCastError) Error() string {
	return fmt.Sprintf("cast to `%s` failed for value `%s`", r.cast, r.object.Type)
}

type FutureTimeoutError struct {
	timeout time.Duration
}

func (e FutureTimeoutError) Error() string {
	return fmt.Sprintf("future timeout has expired (%s)", e.timeout)
}

type ClickTargetOverlappedError struct {
	X, Y      float64
	outerHTML string
}

func (e ClickTargetOverlappedError) Error() string {
	return fmt.Sprintf("click at target is overlapped by `%s`", e.outerHTML)
}
