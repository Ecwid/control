package control

import "errors"

var (
	ErrElementInvisible       = errors.New("element not visible")
	ErrElementIsOutOfViewport = errors.New("element is out of viewport")
	ErrElementMissClick       = errors.New("element miss click")
	ErrStaleElementReference  = errors.New("ErrStaleElementReference")
	ErrAlreadyNavigated       = errors.New("page already navigated to this address - nothing done")
	ErrSessionClosed          = errors.New("this session already closed")
)
