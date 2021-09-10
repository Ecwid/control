package input

import (
	"github.com/ecwid/control/protocol"
)

/*
	Dispatches a drag event into the page.
*/
func DispatchDragEvent(c protocol.Caller, args DispatchDragEventArgs) error {
	return c.Call("Input.dispatchDragEvent", args, nil)
}

/*
	Dispatches a key event to the page.
*/
func DispatchKeyEvent(c protocol.Caller, args DispatchKeyEventArgs) error {
	return c.Call("Input.dispatchKeyEvent", args, nil)
}

/*
	This method emulates inserting text that doesn't come from a key press,
for example an emoji keyboard or an IME.
*/
func InsertText(c protocol.Caller, args InsertTextArgs) error {
	return c.Call("Input.insertText", args, nil)
}

/*
	Dispatches a mouse event to the page.
*/
func DispatchMouseEvent(c protocol.Caller, args DispatchMouseEventArgs) error {
	return c.Call("Input.dispatchMouseEvent", args, nil)
}

/*
	Dispatches a touch event to the page.
*/
func DispatchTouchEvent(c protocol.Caller, args DispatchTouchEventArgs) error {
	return c.Call("Input.dispatchTouchEvent", args, nil)
}

/*
	Emulates touch event from the mouse event parameters.
*/
func EmulateTouchFromMouseEvent(c protocol.Caller, args EmulateTouchFromMouseEventArgs) error {
	return c.Call("Input.emulateTouchFromMouseEvent", args, nil)
}

/*
	Ignores input events (useful while auditing page).
*/
func SetIgnoreInputEvents(c protocol.Caller, args SetIgnoreInputEventsArgs) error {
	return c.Call("Input.setIgnoreInputEvents", args, nil)
}

/*
	Prevents default drag and drop behavior and instead emits `Input.dragIntercepted` events.
Drag and drop behavior can be directly controlled via `Input.dispatchDragEvent`.
*/
func SetInterceptDrags(c protocol.Caller, args SetInterceptDragsArgs) error {
	return c.Call("Input.setInterceptDrags", args, nil)
}

/*
	Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
*/
func SynthesizePinchGesture(c protocol.Caller, args SynthesizePinchGestureArgs) error {
	return c.Call("Input.synthesizePinchGesture", args, nil)
}

/*
	Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
*/
func SynthesizeScrollGesture(c protocol.Caller, args SynthesizeScrollGestureArgs) error {
	return c.Call("Input.synthesizeScrollGesture", args, nil)
}

/*
	Synthesizes a tap gesture over a time period by issuing appropriate touch events.
*/
func SynthesizeTapGesture(c protocol.Caller, args SynthesizeTapGestureArgs) error {
	return c.Call("Input.synthesizeTapGesture", args, nil)
}
