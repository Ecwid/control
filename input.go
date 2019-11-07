package witness

import "time"

// Input events
const (
	dispatchKeyEventChar       = "char"
	dispatchKeyEventKeyDown    = "keyDown"
	dispatchKeyEventKeyUp      = "keyUp"
	dispatchMouseEventMoved    = "mouseMoved"
	dispatchMouseEventPressed  = "mousePressed"
	dispatchMouseEventReleased = "mouseReleased"
)

// MouseMove ...
func (session *CDPSession) MouseMove(x, y float64) error {
	return session.dispatchMouseEvent(x, y, dispatchMouseEventMoved, "none")
}

func (session *CDPSession) sendRune(c rune) error {
	_, err := session.blockingSend("Input.dispatchKeyEvent", Map{
		"type":                  dispatchKeyEventKeyDown,
		"windowsVirtualKeyCode": int(c),
		"nativeVirtualKeyCode":  int(c),
		"unmodifiedText":        string(c),
		"text":                  string(c),
	})
	if err != nil {
		return err
	}
	_, err = session.blockingSend("Input.dispatchKeyEvent", Map{
		"type":                  dispatchKeyEventKeyUp,
		"windowsVirtualKeyCode": int(c),
		"nativeVirtualKeyCode":  int(c),
		"unmodifiedText":        string(c),
		"text":                  string(c),
	})
	return err
}

func (session *CDPSession) dispatchKeyEvent(text string) error {
	for _, c := range text {
		time.Sleep(time.Millisecond * 10)
		if _, err := session.blockingSend("Input.dispatchKeyEvent", Map{
			"type":                  dispatchKeyEventChar,
			"windowsVirtualKeyCode": int(c),
			"nativeVirtualKeyCode":  int(c),
			"unmodifiedText":        string(c),
			"text":                  string(c),
		}); err != nil {
			return err
		}
	}
	return nil
}

// InsertText method emulates inserting text that doesn't come from a key press, for example an emoji keyboard or an IME
func (session *CDPSession) InsertText(text string) error {
	_, err := session.blockingSend("Input.insertText", Map{"text": text})
	return err
}
func (session *CDPSession) dispatchMouseEvent(x float64, y float64, eventType string, button string) error {
	_, err := session.blockingSend("Input.dispatchMouseEvent", Map{
		"type":       eventType,
		"button":     button,
		"x":          x,
		"y":          y,
		"clickCount": 1,
	})
	return err
}

// SendKeys send keyboard keys to focused element
func (session *CDPSession) SendKeys(key ...rune) error {
	var err error
	for _, k := range key {
		err = session.sendRune(k)
		if err != nil {
			return err
		}
	}
	return nil
}
