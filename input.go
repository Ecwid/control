package control

import (
	"sync"
	"time"

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

// Definition ...
type keyDefinition struct {
	keyCode      int
	shiftKeyCode int
	key          string
	shiftKey     string
	code         string
	text         string
	shiftText    string
	location     int
}

func isKey(r rune) bool {
	_, ok := keyDefinitions[r]
	return ok
}

var keyDefinitions = map[rune]keyDefinition{
	'0':  {keyCode: 48, key: "0", code: "Digit0"},
	'1':  {keyCode: 49, key: "1", code: "Digit1"},
	'2':  {keyCode: 50, key: "2", code: "Digit2"},
	'3':  {keyCode: 51, key: "3", code: "Digit3"},
	'4':  {keyCode: 52, key: "4", code: "Digit4"},
	'5':  {keyCode: 53, key: "5", code: "Digit5"},
	'6':  {keyCode: 54, key: "6", code: "Digit6"},
	'7':  {keyCode: 55, key: "7", code: "Digit7"},
	'8':  {keyCode: 56, key: "8", code: "Digit8"},
	'9':  {keyCode: 57, key: "9", code: "Digit9"},
	'\r': {keyCode: 13, code: "Enter", key: "Enter", text: "\r"},
	'\n': {keyCode: 13, code: "Enter", key: "Enter", text: "\r"},
	' ':  {keyCode: 32, key: " ", code: "Space"},
	'a':  {keyCode: 65, key: "a", code: "KeyA"},
	'b':  {keyCode: 66, key: "b", code: "KeyB"},
	'c':  {keyCode: 67, key: "c", code: "KeyC"},
	'd':  {keyCode: 68, key: "d", code: "KeyD"},
	'e':  {keyCode: 69, key: "e", code: "KeyE"},
	'f':  {keyCode: 70, key: "f", code: "KeyF"},
	'g':  {keyCode: 71, key: "g", code: "KeyG"},
	'h':  {keyCode: 72, key: "h", code: "KeyH"},
	'i':  {keyCode: 73, key: "i", code: "KeyI"},
	'j':  {keyCode: 74, key: "j", code: "KeyJ"},
	'k':  {keyCode: 75, key: "k", code: "KeyK"},
	'l':  {keyCode: 76, key: "l", code: "KeyL"},
	'm':  {keyCode: 77, key: "m", code: "KeyM"},
	'n':  {keyCode: 78, key: "n", code: "KeyN"},
	'o':  {keyCode: 79, key: "o", code: "KeyO"},
	'p':  {keyCode: 80, key: "p", code: "KeyP"},
	'q':  {keyCode: 81, key: "q", code: "KeyQ"},
	'r':  {keyCode: 82, key: "r", code: "KeyR"},
	's':  {keyCode: 83, key: "s", code: "KeyS"},
	't':  {keyCode: 84, key: "t", code: "KeyT"},
	'u':  {keyCode: 85, key: "u", code: "KeyU"},
	'v':  {keyCode: 86, key: "v", code: "KeyV"},
	'w':  {keyCode: 87, key: "w", code: "KeyW"},
	'x':  {keyCode: 88, key: "x", code: "KeyX"},
	'y':  {keyCode: 89, key: "y", code: "KeyY"},
	'z':  {keyCode: 90, key: "z", code: "KeyZ"},
	'*':  {keyCode: 106, key: "*", code: "NumpadMultiply", location: 3},
	'+':  {keyCode: 107, key: "+", code: "NumpadAdd", location: 3},
	'-':  {keyCode: 109, key: "-", code: "NumpadSubtract", location: 3},
	'/':  {keyCode: 111, key: "/", code: "NumpadDivide", location: 3},
	';':  {keyCode: 186, key: ";", code: "Semicolon"},
	'=':  {keyCode: 187, key: "=", code: "Equal"},
	',':  {keyCode: 188, key: ",", code: "Comma"},
	'.':  {keyCode: 190, key: ".", code: "Period"},
	'`':  {keyCode: 192, key: "`", code: "Backquote"},
	'[':  {keyCode: 219, key: "[", code: "BracketLeft"},
	'\\': {keyCode: 220, key: "\\", code: "Backslash"},
	']':  {keyCode: 221, key: "]", code: "BracketRight"},
	'\'': {keyCode: 222, key: "'", code: "Quote"},
	')':  {keyCode: 48, key: ")", code: "Digit0"},
	'!':  {keyCode: 49, key: "!", code: "Digit1"},
	'@':  {keyCode: 50, key: "@", code: "Digit2"},
	'#':  {keyCode: 51, key: "#", code: "Digit3"},
	'$':  {keyCode: 52, key: "$", code: "Digit4"},
	'%':  {keyCode: 53, key: "%", code: "Digit5"},
	'^':  {keyCode: 54, key: "^", code: "Digit6"},
	'&':  {keyCode: 55, key: "&", code: "Digit7"},
	'(':  {keyCode: 57, key: "(", code: "Digit9"},
	'A':  {keyCode: 65, key: "A", code: "KeyA"},
	'B':  {keyCode: 66, key: "B", code: "KeyB"},
	'C':  {keyCode: 67, key: "C", code: "KeyC"},
	'D':  {keyCode: 68, key: "D", code: "KeyD"},
	'E':  {keyCode: 69, key: "E", code: "KeyE"},
	'F':  {keyCode: 70, key: "F", code: "KeyF"},
	'G':  {keyCode: 71, key: "G", code: "KeyG"},
	'H':  {keyCode: 72, key: "H", code: "KeyH"},
	'I':  {keyCode: 73, key: "I", code: "KeyI"},
	'J':  {keyCode: 74, key: "J", code: "KeyJ"},
	'K':  {keyCode: 75, key: "K", code: "KeyK"},
	'L':  {keyCode: 76, key: "L", code: "KeyL"},
	'M':  {keyCode: 77, key: "M", code: "KeyM"},
	'N':  {keyCode: 78, key: "N", code: "KeyN"},
	'O':  {keyCode: 79, key: "O", code: "KeyO"},
	'P':  {keyCode: 80, key: "P", code: "KeyP"},
	'Q':  {keyCode: 81, key: "Q", code: "KeyQ"},
	'R':  {keyCode: 82, key: "R", code: "KeyR"},
	'S':  {keyCode: 83, key: "S", code: "KeyS"},
	'T':  {keyCode: 84, key: "T", code: "KeyT"},
	'U':  {keyCode: 85, key: "U", code: "KeyU"},
	'V':  {keyCode: 86, key: "V", code: "KeyV"},
	'W':  {keyCode: 87, key: "W", code: "KeyW"},
	'X':  {keyCode: 88, key: "X", code: "KeyX"},
	'Y':  {keyCode: 89, key: "Y", code: "KeyY"},
	'Z':  {keyCode: 90, key: "Z", code: "KeyZ"},
	':':  {keyCode: 186, key: ":", code: "Semicolon"},
	'<':  {keyCode: 188, key: "<", code: "Comma"},
	'_':  {keyCode: 189, key: "_", code: "Minus"},
	'>':  {keyCode: 190, key: ">", code: "Period"},
	'?':  {keyCode: 191, key: "?", code: "Slash"},
	'~':  {keyCode: 192, key: "~", code: "Backquote"},
	'{':  {keyCode: 219, key: "{", code: "BracketLeft"},
	'|':  {keyCode: 220, key: "|", code: "Backslash"},
	'}':  {keyCode: 221, key: "}", code: "BracketRight"},
	'"':  {keyCode: 222, key: "\"", code: "Quote"},
}

type Input struct {
	mx *sync.Mutex
	s  *Session
}

func (i Input) Click(button input.MouseButton, x, y float64) (err error) {
	i.mx.Lock()
	defer i.mx.Unlock()
	if err = i.MouseMove(MouseNone, x, y); err != nil {
		return err
	}
	if err = i.MousePress(button, x, y); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 150)
	if err = i.MouseRelease(button, x, y); err != nil {
		return err
	}
	return
}

func (i Input) MouseMove(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(i.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mouseMoved",
		Button:     button,
		ClickCount: 1,
	})
}

func (i Input) MousePress(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(i.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mousePressed",
		Button:     button,
		ClickCount: 1,
	})
}

func (i Input) MouseRelease(button input.MouseButton, x, y float64) error {
	return input.DispatchMouseEvent(i.s, input.DispatchMouseEventArgs{
		X:          x,
		Y:          y,
		Type:       "mouseReleased",
		Button:     button,
		ClickCount: 1,
	})
}

// Keyboard events
const (
	dispatchKeyEventKeyDown = "keyDown"
	dispatchKeyEventKeyUp   = "keyUp"
)

func (i Input) InsertText(text string) error {
	return input.InsertText(i.s, input.InsertTextArgs{Text: text})
}

func (i Input) PressKey(c rune) error {
	return i.press(keyDefinition{keyCode: int(c), text: string(c)})
}

func (i Input) press(key keyDefinition) error {
	if key.text == "" {
		key.text = key.key
	}
	err := input.DispatchKeyEvent(i.s, input.DispatchKeyEventArgs{
		Type:                  dispatchKeyEventKeyDown,
		Key:                   key.key,
		Code:                  key.code,
		WindowsVirtualKeyCode: key.keyCode,
		Text:                  key.text,
	})
	if err != nil {
		return err
	}
	return input.DispatchKeyEvent(i.s, input.DispatchKeyEventArgs{
		Type: dispatchKeyEventKeyUp,
		Key:  key.key,
		Code: key.code,
		Text: key.text,
	})
}
