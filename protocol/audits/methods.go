package audits

import (
	"github.com/ecwid/control/protocol"
)

/*
	Returns the response body and size if it were re-encoded with the specified settings. Only

applies to images.
*/
func GetEncodedResponse(c protocol.Caller, args GetEncodedResponseArgs) (*GetEncodedResponseVal, error) {
	var val = &GetEncodedResponseVal{}
	return val, c.Call("Audits.getEncodedResponse", args, val)
}

/*
Disables issues domain, prevents further issues from being reported to the client.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Audits.disable", nil, nil)
}

/*
	Enables issues domain, sends the issues collected so far to the client by means of the

`issueAdded` event.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Audits.enable", nil, nil)
}

/*
	Runs the contrast check for the target page. Found issues are reported

using Audits.issueAdded event.
*/
func CheckContrast(c protocol.Caller, args CheckContrastArgs) error {
	return c.Call("Audits.checkContrast", args, nil)
}
