package control

import (
	"sync"
	"time"

	"github.com/ecwid/control/key"
	"github.com/ecwid/control/protocol"
	"github.com/ecwid/control/protocol/input"
)

const (
	MouseNone    input.MouseButton = "none"
	MouseLeft    input.MouseButton = "left"
	MouseRight   input.MouseButton = "right"
	MouseMiddle  input.MouseButton = "middle"
	MouseBack    input.MouseButton = "back"
	MouseForward input.MouseButton = "forward"
)

func NewMouse(caller protocol.Caller) Mouse {
	return Mouse{
		caller: caller,
		mutex:  &sync.Mutex{},
	}
}

type Mouse struct {
	caller protocol.Caller
	mutex  *sync.Mutex
}

func (m Mouse) Move(button input.MouseButton, point Point) error {
	return input.DispatchMouseEvent(m.caller, input.DispatchMouseEventArgs{
		X:      point.X,
		Y:      point.Y,
		Type:   "mouseMoved",
		Button: button,
	})
}

func (m Mouse) Press(button input.MouseButton, point Point) error {
	return input.DispatchMouseEvent(m.caller, input.DispatchMouseEventArgs{
		X:          point.X,
		Y:          point.Y,
		Type:       "mousePressed",
		Button:     button,
		ClickCount: 1,
	})
}

func (m Mouse) Release(button input.MouseButton, point Point) error {
	return input.DispatchMouseEvent(m.caller, input.DispatchMouseEventArgs{
		X:          point.X,
		Y:          point.Y,
		Type:       "mouseReleased",
		Button:     button,
		ClickCount: 1,
	})
}

func (m Mouse) Down(button input.MouseButton, point Point) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if err = m.Move(MouseNone, point); err != nil {
		return err
	}
	if err = m.Press(button, point); err != nil {
		return err
	}
	return
}

func (m Mouse) Click(button input.MouseButton, point Point, delay time.Duration) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if err = m.Move(MouseNone, point); err != nil {
		return err
	}
	if err = m.Press(button, point); err != nil {
		return err
	}
	time.Sleep(delay)
	if err = m.Release(button, point); err != nil {
		return err
	}
	return
}

type Keyboard struct {
	caller protocol.Caller
}

func NewKeyboard(caller protocol.Caller) Keyboard {
	return Keyboard{caller: caller}
}

func (k Keyboard) Down(key key.Definition) error {
	if key.Text == "" && len(key.Key) == 1 {
		key.Text = key.Key
	}
	return input.DispatchKeyEvent(k.caller, input.DispatchKeyEventArgs{
		Type:                  "keyDown",
		WindowsVirtualKeyCode: key.KeyCode,
		Code:                  key.Code,
		Key:                   key.Key,
		Text:                  key.Text,
		Location:              key.Location,
	})
}

func (k Keyboard) Up(key key.Definition) error {
	return input.DispatchKeyEvent(k.caller, input.DispatchKeyEventArgs{
		Type:                  "keyUp",
		WindowsVirtualKeyCode: key.KeyCode,
		Code:                  key.Code,
		Key:                   key.Key,
	})
}

func (k Keyboard) Insert(text string) error {
	return input.InsertText(k.caller, input.InsertTextArgs{Text: text})
}

func (k Keyboard) Press(key key.Definition, delay time.Duration) (err error) {
	if err = k.Down(key); err != nil {
		return err
	}
	if delay > 0 {
		time.Sleep(delay)
	}
	return k.Up(key)
}

type Touch struct {
	caller protocol.Caller
	mutex  *sync.Mutex
}

func NewTouch(caller protocol.Caller) Touch {
	return Touch{
		caller: caller,
		mutex:  &sync.Mutex{},
	}
}

func (t Touch) Start(x, y, radiusX, radiusY, force float64) error {
	return input.DispatchTouchEvent(t.caller, input.DispatchTouchEventArgs{
		Type: "touchStart",
		TouchPoints: []*input.TouchPoint{
			{
				X:       x,
				Y:       y,
				RadiusX: radiusX,
				RadiusY: radiusY,
				Force:   force,
			},
		},
	})
}

func (t Touch) Swipe(from, to Point) (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if err = t.Start(from.X, from.Y, 1, 1, 1); err != nil {
		return err
	}
	if err = t.Move(to.X, to.Y, 1, 1, 1); err != nil {
		return err
	}
	if err = t.End(); err != nil {
		return err
	}
	return nil
}

func (t Touch) Move(x, y, radiusX, radiusY, force float64) error {
	return input.DispatchTouchEvent(t.caller, input.DispatchTouchEventArgs{
		Type: "touchMove",
		TouchPoints: []*input.TouchPoint{
			{
				X:       x,
				Y:       y,
				RadiusX: radiusX,
				RadiusY: radiusY,
				Force:   force,
			},
		},
	})
}

func (t Touch) End() error {
	return input.DispatchTouchEvent(t.caller, input.DispatchTouchEventArgs{
		Type:        "touchEnd",
		TouchPoints: []*input.TouchPoint{},
	})
}
