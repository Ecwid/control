package control

import (
	"errors"
	"fmt"
	"time"

	"github.com/ecwid/control/protocol/common"

	"github.com/ecwid/control/protocol/target"
)

var (
	ErrElementInvisible       = errors.New("element not visible")
	ErrElementIsOutOfViewport = errors.New("element is out of viewport")
	ErrElementMissClick       = errors.New("element miss click")
	ErrAlreadyNavigated       = errors.New("page already navigated to this address - nothing done")
	ErrTargetDestroyed        = errors.New("this session was destroyed")
	ErrDetachedFromTarget     = errors.New("detached from target")
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

type WaitTimeoutError struct {
	timeout time.Duration
}

func (e WaitTimeoutError) Error() string {
	return fmt.Sprintf("wait condition timeout reached out after %s", e.timeout)
}
