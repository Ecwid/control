package webmcp

import (
	"github.com/ecwid/control/protocol"
)

/*
	Enables the WebMCP domain, allowing events to be sent. Enabling the domain will trigger a toolsAdded event for

all currently registered tools.
*/
func Enable(c protocol.Caller) error {
	return c.Call("WebMCP.enable", nil, nil)
}

/*
Disables the WebMCP domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("WebMCP.disable", nil, nil)
}

/*
Invokes a registered tool.
*/
func InvokeTool(c protocol.Caller, args InvokeToolArgs) (*InvokeToolVal, error) {
	var val = &InvokeToolVal{}
	return val, c.Call("WebMCP.invokeTool", args, val)
}

/*
Cancels a pending tool invocation.
*/
func CancelInvocation(c protocol.Caller, args CancelInvocationArgs) error {
	return c.Call("WebMCP.cancelInvocation", args, nil)
}
