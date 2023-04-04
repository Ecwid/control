package domdebugger

import (
	"github.com/ecwid/control/protocol"
)

/*
Returns event listeners of the given object.
*/
func GetEventListeners(c protocol.Caller, args GetEventListenersArgs) (*GetEventListenersVal, error) {
	var val = &GetEventListenersVal{}
	return val, c.Call("DOMDebugger.getEventListeners", args, val)
}

/*
Removes DOM breakpoint that was set using `setDOMBreakpoint`.
*/
func RemoveDOMBreakpoint(c protocol.Caller, args RemoveDOMBreakpointArgs) error {
	return c.Call("DOMDebugger.removeDOMBreakpoint", args, nil)
}

/*
Removes breakpoint on particular DOM event.
*/
func RemoveEventListenerBreakpoint(c protocol.Caller, args RemoveEventListenerBreakpointArgs) error {
	return c.Call("DOMDebugger.removeEventListenerBreakpoint", args, nil)
}

/*
Removes breakpoint on particular native event.
*/
func RemoveInstrumentationBreakpoint(c protocol.Caller, args RemoveInstrumentationBreakpointArgs) error {
	return c.Call("DOMDebugger.removeInstrumentationBreakpoint", args, nil)
}

/*
Removes breakpoint from XMLHttpRequest.
*/
func RemoveXHRBreakpoint(c protocol.Caller, args RemoveXHRBreakpointArgs) error {
	return c.Call("DOMDebugger.removeXHRBreakpoint", args, nil)
}

/*
Sets breakpoint on particular CSP violations.
*/
func SetBreakOnCSPViolation(c protocol.Caller, args SetBreakOnCSPViolationArgs) error {
	return c.Call("DOMDebugger.setBreakOnCSPViolation", args, nil)
}

/*
Sets breakpoint on particular operation with DOM.
*/
func SetDOMBreakpoint(c protocol.Caller, args SetDOMBreakpointArgs) error {
	return c.Call("DOMDebugger.setDOMBreakpoint", args, nil)
}

/*
Sets breakpoint on particular DOM event.
*/
func SetEventListenerBreakpoint(c protocol.Caller, args SetEventListenerBreakpointArgs) error {
	return c.Call("DOMDebugger.setEventListenerBreakpoint", args, nil)
}

/*
Sets breakpoint on particular native event.
*/
func SetInstrumentationBreakpoint(c protocol.Caller, args SetInstrumentationBreakpointArgs) error {
	return c.Call("DOMDebugger.setInstrumentationBreakpoint", args, nil)
}

/*
Sets breakpoint on XMLHttpRequest.
*/
func SetXHRBreakpoint(c protocol.Caller, args SetXHRBreakpointArgs) error {
	return c.Call("DOMDebugger.setXHRBreakpoint", args, nil)
}
