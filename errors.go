package witness

import "errors"

// cdp errors
var (
	ErrStaleElementReference = errors.New("referenced element is no longer attached to the DOM")
	ErrNoSuchElement         = errors.New("no such element")
	ErrNoPageTarget          = errors.New("no one page target")
	ErrDevtoolTimeout        = errors.New("devtool response reached timeout")
	ErrNavigateTimeout       = errors.New("navigation reached timeout")
	ErrSessionClosed         = errors.New("session closed")
	ErrElementInvisible      = errors.New("element invisible")
	ErrElementOverlapped     = errors.New("element overlapped")
	ErrElementNotFocusable   = errors.New("element is not focusable")
	ErrClickNotTriggered     = errors.New("click not confirmed")
	ErrHoverNotTriggered     = errors.New("mouseover not confirmed")
	ErrInvalidString         = errors.New("object type is not string")
	ErrInvalidElementFrame   = errors.New("specified element is not a IFRAME")
	ErrInvalidElementSelect  = errors.New("specified element is not a SELECT")
	ErrInvalidElementOption  = errors.New("specified element has no options")
	ErrCannotFindContext     = errors.New("cannot find context with specified id")
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
		return ErrCannotFindContext
	case "Could not compute content quads.":
		return ErrElementInvisible
	case "Element is not focusable":
		return ErrElementNotFocusable
	}
	return e
}
