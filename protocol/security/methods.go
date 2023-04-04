package security

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables tracking security state changes.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Security.disable", nil, nil)
}

/*
Enables tracking security state changes.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Security.enable", nil, nil)
}

/*
Enable/disable whether all certificate errors should be ignored.
*/
func SetIgnoreCertificateErrors(c protocol.Caller, args SetIgnoreCertificateErrorsArgs) error {
	return c.Call("Security.setIgnoreCertificateErrors", args, nil)
}
