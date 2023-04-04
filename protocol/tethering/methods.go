package tethering

import (
	"github.com/ecwid/control/protocol"
)

/*
Request browser port binding.
*/
func Bind(c protocol.Caller, args BindArgs) error {
	return c.Call("Tethering.bind", args, nil)
}

/*
Request browser port unbinding.
*/
func Unbind(c protocol.Caller, args UnbindArgs) error {
	return c.Call("Tethering.unbind", args, nil)
}
