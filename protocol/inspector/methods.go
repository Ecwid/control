package inspector

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables inspector domain notifications.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Inspector.disable", nil, nil)
}

/*
Enables inspector domain notifications.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Inspector.enable", nil, nil)
}
