package eventbreakpoints

import (
	"github.com/ecwid/control/protocol"
)

/*
Sets breakpoint on particular native event.
*/
func SetInstrumentationBreakpoint(c protocol.Caller, args SetInstrumentationBreakpointArgs) error {
	return c.Call("EventBreakpoints.setInstrumentationBreakpoint", args, nil)
}

/*
Removes breakpoint on particular native event.
*/
func RemoveInstrumentationBreakpoint(c protocol.Caller, args RemoveInstrumentationBreakpointArgs) error {
	return c.Call("EventBreakpoints.removeInstrumentationBreakpoint", args, nil)
}
