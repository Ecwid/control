package control

import (
	"errors"
	"fmt"

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
