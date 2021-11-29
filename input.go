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

type KeyDefinition struct {
	KeyCode      int
	ShiftKeyCode int
	Key          string
	ShiftKey     string
	Code         string
	Text         string
	ShiftText    string
	Location     int
}

func isKey(r rune) bool {
	_, ok := keyDefinitions[r]
	return ok
}

var keyDefinitions = map[rune]KeyDefinition{
	'0':  {KeyCode: 48, Key: "0", Code: "Digit0"},
	'1':  {KeyCode: 49, Key: "1", Code: "Digit1"},
	'2':  {KeyCode: 50, Key: "2", Code: "Digit2"},
	'3':  {KeyCode: 51, Key: "3", Code: "Digit3"},
	'4':  {KeyCode: 52, Key: "4", Code: "Digit4"},
	'5':  {KeyCode: 53, Key: "5", Code: "Digit5"},
	'6':  {KeyCode: 54, Key: "6", Code: "Digit6"},
	'7':  {KeyCode: 55, Key: "7", Code: "Digit7"},
	'8':  {KeyCode: 56, Key: "8", Code: "Digit8"},
	'9':  {KeyCode: 57, Key: "9", Code: "Digit9"},
	'\r': {KeyCode: 13, Code: "Enter", Key: "Enter", Text: "\r"},
	'\n': {KeyCode: 13, Code: "Enter", Key: "Enter", Text: "\r"},
	' ':  {KeyCode: 32, Key: " ", Code: "Space"},
	'a':  {KeyCode: 65, Key: "a", Code: "KeyA"},
	'b':  {KeyCode: 66, Key: "b", Code: "KeyB"},
	'c':  {KeyCode: 67, Key: "c", Code: "KeyC"},
	'd':  {KeyCode: 68, Key: "d", Code: "KeyD"},
	'e':  {KeyCode: 69, Key: "e", Code: "KeyE"},
	'f':  {KeyCode: 70, Key: "f", Code: "KeyF"},
	'g':  {KeyCode: 71, Key: "g", Code: "KeyG"},
	'h':  {KeyCode: 72, Key: "h", Code: "KeyH"},
	'i':  {KeyCode: 73, Key: "i", Code: "KeyI"},
	'j':  {KeyCode: 74, Key: "j", Code: "KeyJ"},
	'k':  {KeyCode: 75, Key: "k", Code: "KeyK"},
	'l':  {KeyCode: 76, Key: "l", Code: "KeyL"},
	'm':  {KeyCode: 77, Key: "m", Code: "KeyM"},
	'n':  {KeyCode: 78, Key: "n", Code: "KeyN"},
	'o':  {KeyCode: 79, Key: "o", Code: "KeyO"},
	'p':  {KeyCode: 80, Key: "p", Code: "KeyP"},
	'q':  {KeyCode: 81, Key: "q", Code: "KeyQ"},
	'r':  {KeyCode: 82, Key: "r", Code: "KeyR"},
	's':  {KeyCode: 83, Key: "s", Code: "KeyS"},
	't':  {KeyCode: 84, Key: "t", Code: "KeyT"},
	'u':  {KeyCode: 85, Key: "u", Code: "KeyU"},
	'v':  {KeyCode: 86, Key: "v", Code: "KeyV"},
	'w':  {KeyCode: 87, Key: "w", Code: "KeyW"},
	'x':  {KeyCode: 88, Key: "x", Code: "KeyX"},
	'y':  {KeyCode: 89, Key: "y", Code: "KeyY"},
	'z':  {KeyCode: 90, Key: "z", Code: "KeyZ"},
	'*':  {KeyCode: 106, Key: "*", Code: "NumpadMultiply", Location: 3},
	'+':  {KeyCode: 107, Key: "+", Code: "NumpadAdd", Location: 3},
	'-':  {KeyCode: 109, Key: "-", Code: "NumpadSubtract", Location: 3},
	'/':  {KeyCode: 111, Key: "/", Code: "NumpadDivide", Location: 3},
	';':  {KeyCode: 186, Key: ";", Code: "Semicolon"},
	'=':  {KeyCode: 187, Key: "=", Code: "Equal"},
	',':  {KeyCode: 188, Key: ",", Code: "Comma"},
	'.':  {KeyCode: 190, Key: ".", Code: "Period"},
	'`':  {KeyCode: 192, Key: "`", Code: "Backquote"},
	'[':  {KeyCode: 219, Key: "[", Code: "BracketLeft"},
	'\\': {KeyCode: 220, Key: "\\", Code: "Backslash"},
	']':  {KeyCode: 221, Key: "]", Code: "BracketRight"},
	'\'': {KeyCode: 222, Key: "'", Code: "Quote"},
	')':  {KeyCode: 48, Key: ")", Code: "Digit0"},
	'!':  {KeyCode: 49, Key: "!", Code: "Digit1"},
	'@':  {KeyCode: 50, Key: "@", Code: "Digit2"},
	'#':  {KeyCode: 51, Key: "#", Code: "Digit3"},
	'$':  {KeyCode: 52, Key: "$", Code: "Digit4"},
	'%':  {KeyCode: 53, Key: "%", Code: "Digit5"},
	'^':  {KeyCode: 54, Key: "^", Code: "Digit6"},
	'&':  {KeyCode: 55, Key: "&", Code: "Digit7"},
	'(':  {KeyCode: 57, Key: "(", Code: "Digit9"},
	'A':  {KeyCode: 65, Key: "A", Code: "KeyA"},
	'B':  {KeyCode: 66, Key: "B", Code: "KeyB"},
	'C':  {KeyCode: 67, Key: "C", Code: "KeyC"},
	'D':  {KeyCode: 68, Key: "D", Code: "KeyD"},
	'E':  {KeyCode: 69, Key: "E", Code: "KeyE"},
	'F':  {KeyCode: 70, Key: "F", Code: "KeyF"},
	'G':  {KeyCode: 71, Key: "G", Code: "KeyG"},
	'H':  {KeyCode: 72, Key: "H", Code: "KeyH"},
	'I':  {KeyCode: 73, Key: "I", Code: "KeyI"},
	'J':  {KeyCode: 74, Key: "J", Code: "KeyJ"},
	'K':  {KeyCode: 75, Key: "K", Code: "KeyK"},
	'L':  {KeyCode: 76, Key: "L", Code: "KeyL"},
	'M':  {KeyCode: 77, Key: "M", Code: "KeyM"},
	'N':  {KeyCode: 78, Key: "N", Code: "KeyN"},
	'O':  {KeyCode: 79, Key: "O", Code: "KeyO"},
	'P':  {KeyCode: 80, Key: "P", Code: "KeyP"},
	'Q':  {KeyCode: 81, Key: "Q", Code: "KeyQ"},
	'R':  {KeyCode: 82, Key: "R", Code: "KeyR"},
	'S':  {KeyCode: 83, Key: "S", Code: "KeyS"},
	'T':  {KeyCode: 84, Key: "T", Code: "KeyT"},
	'U':  {KeyCode: 85, Key: "U", Code: "KeyU"},
	'V':  {KeyCode: 86, Key: "V", Code: "KeyV"},
	'W':  {KeyCode: 87, Key: "W", Code: "KeyW"},
	'X':  {KeyCode: 88, Key: "X", Code: "KeyX"},
	'Y':  {KeyCode: 89, Key: "Y", Code: "KeyY"},
	'Z':  {KeyCode: 90, Key: "Z", Code: "KeyZ"},
	':':  {KeyCode: 186, Key: ":", Code: "Semicolon"},
	'<':  {KeyCode: 188, Key: "<", Code: "Comma"},
	'_':  {KeyCode: 189, Key: "_", Code: "Minus"},
	'>':  {KeyCode: 190, Key: ">", Code: "Period"},
	'?':  {KeyCode: 191, Key: "?", Code: "Slash"},
	'~':  {KeyCode: 192, Key: "~", Code: "Backquote"},
	'{':  {KeyCode: 219, Key: "{", Code: "BracketLeft"},
	'|':  {KeyCode: 220, Key: "|", Code: "Backslash"},
	'}':  {KeyCode: 221, Key: "}", Code: "BracketRight"},
	'"':  {KeyCode: 222, Key: "\"", Code: "Quote"},
}

type Input struct {
	mx *sync.Mutex
	s  *Session
}

func (i Input) Click(button input.MouseButton, x, y float64, delay time.Duration) (err error) {
	i.mx.Lock()
	defer i.mx.Unlock()
	if err = i.MouseMove(MouseNone, x, y); err != nil {
		return err
	}
	if err = i.MousePress(button, x, y); err != nil {
		return err
	}
	time.Sleep(delay)
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
	return i.Press(KeyDefinition{KeyCode: int(c), Text: string(c)})
}

func (i Input) Press(key KeyDefinition) error {
	if key.Text == "" {
		key.Text = key.Key
	}
	err := input.DispatchKeyEvent(i.s, input.DispatchKeyEventArgs{
		Type:                  dispatchKeyEventKeyDown,
		Key:                   key.Key,
		Code:                  key.Code,
		WindowsVirtualKeyCode: key.KeyCode,
		Text:                  key.Text,
	})
	if err != nil {
		return err
	}
	return input.DispatchKeyEvent(i.s, input.DispatchKeyEventArgs{
		Type: dispatchKeyEventKeyUp,
		Key:  key.Key,
		Code: key.Code,
		Text: key.Text,
	})
}
