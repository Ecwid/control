package control

import "github.com/ecwid/control/protocol/input"

const (
	MouseNone    input.MouseButton = "none"
	MouseLeft    input.MouseButton = "left"
	MouseRight   input.MouseButton = "right"
	MouseMiddle  input.MouseButton = "middle"
	MouseBack    input.MouseButton = "back"
	MouseForward input.MouseButton = "forward"
)

type Mouse struct {
	s *Session
}

func (m Mouse) Move(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(m.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mouseMoved",
		Button:     button,
		ClickCount: 1,
	})
}

func (m Mouse) Press(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(m.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mousePressed",
		Button:     button,
		ClickCount: 1,
	})
}

func (m Mouse) Release(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(m.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mouseReleased",
		Button:     button,
		ClickCount: 1,
	})
}

type Keyboard struct {
	s *Session
}
