package headlessexperimental

import (
	"github.com/ecwid/control/protocol"
)

/*
	Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a
screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also
https://goo.gl/3zHXhB for more background.
*/
func BeginFrame(c protocol.Caller, args BeginFrameArgs) (*BeginFrameVal, error) {
	var val = &BeginFrameVal{}
	return val, c.Call("HeadlessExperimental.beginFrame", args, val)
}

/*
	Disables headless events for the target.
*/
func Disable(c protocol.Caller) error {
	return c.Call("HeadlessExperimental.disable", nil, nil)
}

/*
	Enables headless events for the target.
*/
func Enable(c protocol.Caller) error {
	return c.Call("HeadlessExperimental.enable", nil, nil)
}
