package control

import (
	"fmt"
	"math"
	"time"

	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/input"
	"github.com/ecwid/control/protocol/runtime"
)

type Element struct {
	remote *runtime.RemoteObject
	frame  *Frame
}

func (e Element) Description() string {
	return e.remote.Description
}

func (e Element) QuerySelector(selector string) (*Element, error) {
	val, err := e.CallFunction(`function(s){return this.querySelector(s)}`, true, false, NewCallArgument(selector))
	if err != nil {
		return nil, err
	}
	return &Element{frame: e.frame, remote: val}, nil
}

func (e Element) CallFunction(function string, await, returnByValue bool, args ...*runtime.CallArgument) (*runtime.RemoteObject, error) {
	val, err := runtime.CallFunctionOn(e.frame, runtime.CallFunctionOnArgs{
		FunctionDeclaration: function,
		ObjectId:            e.remote.ObjectId,
		AwaitPromise:        await,
		ReturnByValue:       returnByValue,
		Arguments:           args,
	})
	if err != nil {
		return nil, err
	}
	if val.ExceptionDetails != nil {
		return nil, RuntimeError(*val.ExceptionDetails)
	}
	return val.Result, nil
}

func (e Element) DescribeNode() (*dom.Node, error) {
	val, err := dom.DescribeNode(e.frame, dom.DescribeNodeArgs{
		ObjectId: e.remote.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return val.Node, nil
}

func NewCallArgument(v interface{}) *runtime.CallArgument {
	return &runtime.CallArgument{Value: v}
}

func (e Element) dispatchEvents(events ...string) error {
	var args = make([]*runtime.CallArgument, len(events))
	for i, a := range events {
		args[i] = NewCallArgument(a)
	}
	_, err := e.CallFunction(functionDispatchEvents, true, false, args...)
	return err
}

func (e Element) ScrollIntoView() error {
	return dom.ScrollIntoViewIfNeeded(e.frame, dom.ScrollIntoViewIfNeededArgs{ObjectId: e.remote.ObjectId})
}

func (e Element) GetText() (string, error) {
	v, err := e.CallFunction(functionGetText, true, false)
	if err != nil {
		return "null", err
	}
	return fmt.Sprint(v.Value), nil
}

func (e Element) Clear() error {
	_, err := e.CallFunction(functionClearText, true, false)
	return err
}

func (e Element) InsertText(text string) error {
	var err error
	if err = e.ScrollIntoView(); err != nil {
		return err
	}
	if err = e.Clear(); err != nil {
		return err
	}
	if err = e.Focus(); err != nil {
		return err
	}
	if err = e.frame.Session().Keyboard.InsertText(text); err != nil {
		return err
	}
	if err = e.dispatchEvents(
		WebEventKeypress,
		WebEventInput,
		WebEventKeyup,
		WebEventChange,
	); err != nil {
		return err
	}
	return nil
}

// Type ...
func (e *Element) Type(text string, delay time.Duration) error {
	var err error
	if err = e.ScrollIntoView(); err != nil {
		return err
	}
	if err = e.Clear(); err != nil {
		return err
	}
	if err = e.Focus(); err != nil {
		return err
	}
	for _, c := range text {
		if isKey(c) {
			if err = e.frame.Session().Keyboard.press(keyDefinitions[c]); err != nil {
				return err
			}
		} else {
			if err = e.InsertText(string(c)); err != nil {
				return err
			}
		}
		time.Sleep(delay)
	}
	if text == "" {
		return e.dispatchEvents(
			WebEventKeypress,
			WebEventInput,
			WebEventKeyup,
			WebEventChange,
		)
	}
	return nil
}

func (e Element) GetContentQuad(viewportCorrection bool) (Quad, error) {
	val, err := dom.GetContentQuads(e.frame, dom.GetContentQuadsArgs{
		ObjectId: e.remote.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	quads := convertQuads(val.Quads)
	if len(quads) == 0 { // should be at least one
		return nil, ErrElementInvisible
	}
	metric, err := e.frame.session.GetLayoutMetrics()
	if err != nil {
		return nil, err
	}
	for _, quad := range quads {
		/* correction is get sub-quad of element that in viewport
		 _______________  <- Viewport top
		|  1 _______ 2  |
		|   |visible|   | visible part of element
		|__4|visible|3__| <- Viewport bottom
		|   |invisib|   | this invisible part of element omits if viewportCorrection
		|...............|
		*/
		if viewportCorrection {
			for _, point := range quad {
				point.X = math.Min(math.Max(point.X, 0), float64(metric.CssLayoutViewport.ClientWidth))
				point.Y = math.Min(math.Max(point.Y, 0), float64(metric.CssLayoutViewport.ClientHeight))
			}
		}
		if quad.Area() > 1 {
			return quad, nil
		}
	}
	return nil, ErrElementIsOutOfViewport
}

func (e Element) clickablePoint() (x float64, y float64, err error) {
	r, err := e.GetContentQuad(true)
	if err != nil {
		return -1, -1, err
	}
	x, y = r.Middle()
	return x, y, nil
}

func (e Element) Click() error {
	return e.click(MouseLeft)
}

func (e Element) click(button input.MouseButton) error {
	if err := e.ScrollIntoView(); err != nil {
		return err
	}
	x, y, err := e.clickablePoint()
	if err != nil {
		return err
	}
	if _, err = e.CallFunction(functionPreventMissClick, true, false); err != nil {
		return err
	}
	var mouse = e.frame.Session().Mouse
	if err = mouse.Move(MouseNone, x, y); err != nil {
		return err
	}
	if err = mouse.Press(button, x, y); err != nil {
		return err
	}
	if err = mouse.Release(button, x, y); err != nil {
		return err
	}
	clicked, err := e.CallFunction(functionClickDone, true, false)
	if err != nil {
		return nil // context was destroyed by navigate after click
	}
	if val, ok := clicked.Value.(bool); ok && !val {
		return ErrElementMissClick
	}
	return nil
}

func (e Element) Focus() error {
	return dom.Focus(e.frame, dom.FocusArgs{ObjectId: e.remote.ObjectId})
}

func (e Element) Upload(files ...string) error {
	return dom.SetFileInputFiles(e.frame, dom.SetFileInputFilesArgs{
		Files:    files,
		ObjectId: e.remote.ObjectId,
	})
}

func (e Element) Hover() error {
	if err := e.ScrollIntoView(); err != nil {
		return err
	}
	x, y, err := e.clickablePoint()
	if err != nil {
		return err
	}
	return e.frame.Session().Mouse.Move(MouseNone, x, y)
}

func (e Element) SetAttribute(attr string, value string) error {
	_, err := e.CallFunction(functionSetAttr, true, false, NewCallArgument(attr), NewCallArgument(value))
	return err
}

func (e Element) GetAttribute(attr string) (string, error) {
	return e.callFunctionStringValue(functionGetAttr, NewCallArgument(attr))
}

func (e Element) Checkbox(check bool) error {
	if _, err := e.CallFunction(functionCheckbox, true, false, NewCallArgument(check)); err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e *Element) IsChecked() (bool, error) {
	v, err := e.CallFunction(functionIsChecked, true, false)
	if err != nil {
		return false, err
	}
	return remoteObjectPrimitive(*v).Bool()
}

func (e Element) GetRectangle() (*dom.Rect, error) {
	q, err := e.GetContentQuad(false)
	if err != nil {
		return nil, err
	}
	rect := &dom.Rect{
		X:      q[0].X,
		Y:      q[0].Y,
		Width:  q[1].X - q[0].X,
		Height: q[3].Y - q[0].Y,
	}
	return rect, nil
}

func (e Element) GetComputedStyle(style string) (string, error) {
	return e.callFunctionStringValue(functionGetComputedStyle, NewCallArgument(style))
}

func (e Element) SelectValues(values ...string) error {
	if e.remote.ClassName != "HTMLSelectElement" {
		return fmt.Errorf("can't use element as SELECT, not applicable type %s", e.remote.ClassName)
	}
	_, err := e.CallFunction(functionSelect, true, false, NewCallArgument(values))
	if err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e Element) GetSelectedValues() ([]string, error) {
	return e.callFunctionStringArrayValue(functionGetSelectedValues)
}

func (e Element) GetSelectedText() ([]string, error) {
	return e.callFunctionStringArrayValue(functionGetSelectedInnerText)
}

func (e Element) callFunctionStringValue(function string, args ...*runtime.CallArgument) (string, error) {
	v, err := e.CallFunction(function, true, false, args...)
	if err != nil {
		return "", err
	}
	return remoteObjectPrimitive(*v).String()
}

func (e Element) callFunctionStringArrayValue(function string, args ...*runtime.CallArgument) ([]string, error) {
	v, err := e.CallFunction(function, true, false, args...)
	if err != nil {
		return nil, err
	}
	descriptor, err := e.frame.getProperties(v.ObjectId, true, false)
	if err != nil {
		return nil, err
	}
	var options []string
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		options = append(options, d.Value.Value.(string))
	}
	return options, nil
}
