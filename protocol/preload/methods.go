package preload

import (
	"github.com/ecwid/control/protocol"
)

/*
 */
func Enable(c protocol.Caller) error {
	return c.Call("Preload.enable", nil, nil)
}

/*
 */
func Disable(c protocol.Caller) error {
	return c.Call("Preload.disable", nil, nil)
}
