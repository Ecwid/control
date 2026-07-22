package fedcm

import (
	"github.com/ecwid/control/protocol"
)

/*
 */
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("FedCm.enable", args, nil)
}

/*
 */
func Disable(c protocol.Caller) error {
	return c.Call("FedCm.disable", nil, nil)
}

/*
 */
func SelectAccount(c protocol.Caller, args SelectAccountArgs) error {
	return c.Call("FedCm.selectAccount", args, nil)
}

/*
 */
func ClickDialogButton(c protocol.Caller, args ClickDialogButtonArgs) error {
	return c.Call("FedCm.clickDialogButton", args, nil)
}

/*
 */
func OpenUrl(c protocol.Caller, args OpenUrlArgs) error {
	return c.Call("FedCm.openUrl", args, nil)
}

/*
 */
func DismissDialog(c protocol.Caller, args DismissDialogArgs) error {
	return c.Call("FedCm.dismissDialog", args, nil)
}

/*
	Resets the cooldown time, if any, to allow the next FedCM call to show

a dialog even if one was recently dismissed by the user.
*/
func ResetCooldown(c protocol.Caller) error {
	return c.Call("FedCm.resetCooldown", nil, nil)
}
