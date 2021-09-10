package log

import (
	"github.com/ecwid/control/protocol"
)

/*
	Clears the log.
*/
func Clear(c protocol.Caller) error {
	return c.Call("Log.clear", nil, nil)
}

/*
	Disables log domain, prevents further log entries from being reported to the client.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Log.disable", nil, nil)
}

/*
	Enables log domain, sends the entries collected so far to the client by means of the
`entryAdded` notification.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Log.enable", nil, nil)
}

/*
	start violation reporting.
*/
func StartViolationsReport(c protocol.Caller, args StartViolationsReportArgs) error {
	return c.Call("Log.startViolationsReport", args, nil)
}

/*
	Stop violation reporting.
*/
func StopViolationsReport(c protocol.Caller) error {
	return c.Call("Log.stopViolationsReport", nil, nil)
}
