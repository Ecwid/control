package witness

import "errors"

// cdp errors
var (
	ErrStaleElementReference  = errors.New("referenced element is no longer attached to the DOM")
	ErrNoSuchContext          = errors.New("cannot find context with specified id")
	ErrNoSuchElement          = errors.New("no such element")
	ErrNoSuchFrame            = errors.New("no such frame")
	ErrFrameDetached          = errors.New("frame you working on was detached")
	ErrNoPageTarget           = errors.New("no one page target")
	ErrDevtoolTimeout         = errors.New("devtool response reached timeout")
	ErrNavigateTimeout        = errors.New("navigation reached timeout")
	ErrSessionClosed          = errors.New("session closed")
	ErrElementInvisible       = errors.New("element invisible")
	ErrElementIsOutOfViewport = errors.New("element is out of viewport")
	ErrElementNotFocusable    = errors.New("element is not focusable")
	ErrElementMissClick       = errors.New("miss click - element is overlapping or changing its coordinates")
	ErrInvalidString          = errors.New("object type is not string")
	ErrInvalidElementFrame    = errors.New("specified element is not a IFRAME")
	ErrInvalidElementSelect   = errors.New("specified element is not a SELECT")
	ErrInvalidElementOption   = errors.New("specified element has no options")
)

// v8 inspector devtool errors
const (
	ProtocolParseError     = -32700
	ProtocolInvalidRequest = -32600
	ProtocolMethodNotFound = -32601
	ProtocolInvalidParams  = -32602
	ProtocolInternalError  = -32603
	ProtocolServerError    = -32000
)

func (e rpcError) known() error {
	switch e.Message {
	case "Cannot find context with specified id":
		return ErrNoSuchContext
	case "Could not compute content quads.":
		return ErrElementInvisible
	case "Element is not focusable":
		return ErrElementNotFocusable
	}
	return e
}
