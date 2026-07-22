package deviceaccess

import (
	"github.com/ecwid/control/protocol"
)

/*
Enable events in this domain.
*/
func Enable(c protocol.Caller) error {
	return c.Call("DeviceAccess.enable", nil, nil)
}

/*
Disable events in this domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("DeviceAccess.disable", nil, nil)
}

/*
Select a device in response to a DeviceAccess.deviceRequestPrompted event.
*/
func SelectPrompt(c protocol.Caller, args SelectPromptArgs) error {
	return c.Call("DeviceAccess.selectPrompt", args, nil)
}

/*
Cancel a prompt in response to a DeviceAccess.deviceRequestPrompted event.
*/
func CancelPrompt(c protocol.Caller, args CancelPromptArgs) error {
	return c.Call("DeviceAccess.cancelPrompt", args, nil)
}
