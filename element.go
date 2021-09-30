package control

import (
	"fmt"
	"math"
	"time"

	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/input"
	"github.com/ecwid/control/protocol/runtime"
)

func (f Frame) constructElement(object *runtime.RemoteObject) (*Element, error) {
	val, err := dom.DescribeNode(f, dom.DescribeNodeArgs{
		ObjectId: object.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &Element{node: val.Node, runtime: object, frame: &f}, nil
}

type Element struct {
	runtime *runtime.RemoteObject
	node    *dom.Node
	frame   *Frame
}

func (e Element) Description() string {
	return e.runtime.Description
}

func (e Element) Node() *dom.Node {
	return e.node
}

func (e Element) QuerySelector(selector string) (*Element, error) {
	val, err := e.CallFunction(`function(s){return this.querySelector(s)}`, true, false, NewSingleCallArgument(selector))
	if err != nil {
		return nil, err
	}
	return e.frame.constructElement(val)
}

func (e Element) CallFunction(function string, await, returnByValue bool, args []*runtime.CallArgument) (*runtime.RemoteObject, error) {
	val, err := runtime.CallFunctionOn(e.frame, runtime.CallFunctionOnArgs{
		FunctionDeclaration: function,
		ObjectId:            e.runtime.ObjectId,
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

func NewSingleCallArgument(arg interface{}) []*runtime.CallArgument {
	return []*runtime.CallArgument{{Value: arg}}
}

func (e Element) dispatchEvents(events ...string) error {
	_, err := e.CallFunction(functionDispatchEvents, true, false, NewSingleCallArgument(events))
	return err
}

func (e Element) ScrollIntoView() error {
	return dom.ScrollIntoViewIfNeeded(e.frame, dom.ScrollIntoViewIfNeededArgs{
		BackendNodeId: e.node.BackendNodeId,
	})
}

func (e Element) GetText() (string, error) {
	v, err := e.CallFunction(functionGetText, true, false, nil)
	if err != nil {
		return "null", err
	}
	return fmt.Sprint(v.Value), nil
}

func (e Element) Clear() error {
	_, err := e.CallFunction(functionClearText, true, false, nil)
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
	if err = e.frame.Session().Input.InsertText(text); err != nil {
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
			if err = e.frame.Session().Input.press(keyDefinitions[c]); err != nil {
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
		BackendNodeId: e.node.BackendNodeId,
	})
	if err != nil {
		return nil, err
	}
	quads := convertQuads(val.Quads)
	if len(quads) == 0 { // should be at least one
		return nil, ErrElementInvisible
	}
	metric, err := e.frame.Session().GetLayoutMetrics()
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
	if _, err = e.CallFunction(functionPreventMissClick, true, false, nil); err != nil {
		return err
	}
	var clickValue = make(chan string, 1)
	defer close(clickValue)
	cancel := e.frame.session.onBindingCalled(bindClick, func(payload string) {
		clickValue <- payload
	})
	defer cancel()
	if err = e.frame.Session().Input.Click(button, x, y); err != nil {
		return err
	}
	var timeout = time.NewTimer(e.frame.session.Timeout)
	defer timeout.Stop()
	select {
	case v := <-clickValue:
		if v != "1" {
			return ErrElementMissClick
		}
	case <-timeout.C:
		return WaitTimeoutError{timeout: e.frame.session.Timeout}
	}
	return nil
}

func (e Element) Focus() error {
	return dom.Focus(e.frame, dom.FocusArgs{BackendNodeId: e.node.BackendNodeId})
}

func (e Element) Upload(files ...string) error {
	return dom.SetFileInputFiles(e.frame, dom.SetFileInputFilesArgs{
		Files:         files,
		BackendNodeId: e.node.BackendNodeId,
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
	return e.frame.Session().Input.MouseMove(MouseNone, x, y)
}

func (e Element) SetAttribute(attr string, value string) error {
	_, err := e.CallFunction(functionSetAttr, true, false, []*runtime.CallArgument{
		{Value: attr},
		{Value: value},
	})
	return err
}

func (e Element) GetAttribute(attr string) (string, error) {
	v, err := e.CallFunction(functionGetAttr, true, false, NewSingleCallArgument(attr))
	if err != nil {
		return "", err
	}
	return primitiveRemoteObject(*v).String()
}

func (e Element) Checkbox(check bool) error {
	if _, err := e.CallFunction(functionCheckbox, true, false, NewSingleCallArgument(check)); err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e *Element) IsChecked() (bool, error) {
	v, err := e.CallFunction(functionIsChecked, true, false, nil)
	if err != nil {
		return false, err
	}
	return primitiveRemoteObject(*v).Bool()
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
	v, err := e.CallFunction(functionGetComputedStyle, true, false, NewSingleCallArgument(style))
	if err != nil {
		return "", err
	}
	return primitiveRemoteObject(*v).String()
}

func (e Element) SelectValues(values ...string) error {
	if "SELECT" != e.node.NodeName {
		return fmt.Errorf("can't use element as SELECT, not applicable type %s", e.node.NodeName)
	}
	_, err := e.CallFunction(functionSelect, true, false, NewSingleCallArgument(values))
	if err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e Element) GetSelectedValues() ([]string, error) {
	v, err := e.CallFunction(functionGetSelectedValues, true, false, nil)
	if err != nil {
		return nil, err
	}
	return e.stringArray(v)
}

func (e Element) GetSelectedText() ([]string, error) {
	v, err := e.CallFunction(functionGetSelectedInnerText, true, false, nil)
	if err != nil {
		return nil, err
	}
	return e.stringArray(v)
}

func (e Element) stringArray(v *runtime.RemoteObject) ([]string, error) {
	descriptor, err := e.frame.getProperties(v.ObjectId, true, false)
	if err != nil {
		return nil, err
	}
	var options []string
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		val, err1 := primitiveRemoteObject(*d.Value).String()
		if err1 != nil {
			return nil, err1
		}
		options = append(options, val)
	}
	return options, nil
}
