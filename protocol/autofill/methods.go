package autofill

import (
	"github.com/ecwid/control/protocol"
)

/*
	Trigger autofill on a form identified by the fieldId.

If the field and related form cannot be autofilled, returns an error.
*/
func Trigger(c protocol.Caller, args TriggerArgs) error {
	return c.Call("Autofill.trigger", args, nil)
}

/*
Set addresses so that developers can verify their forms implementation.
*/
func SetAddresses(c protocol.Caller, args SetAddressesArgs) error {
	return c.Call("Autofill.setAddresses", args, nil)
}

/*
Disables autofill domain notifications.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Autofill.disable", nil, nil)
}

/*
Enables autofill domain notifications.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Autofill.enable", nil, nil)
}
