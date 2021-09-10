package control

import (
	"fmt"
	"math"

	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/input"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
)

type Element struct {
	remote *runtime.RemoteObject
	frame  *Frame
}

func (e Element) Description() string {
	return e.remote.Description
}

func (e Element) callFunction(function string, await, returnByValue bool, args []*runtime.CallArgument) (*runtime.RemoteObject, error) {
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

func (e Element) ToFrame() (*Frame, error) {
	if e.remote.ClassName != "HTMLIFrameElement" &&
		e.remote.ClassName != "HTMLFrameElement" {
		return nil, fmt.Errorf("can't use element as frame, not applicable type %s", e.remote.ClassName)
	}
	val, err := dom.DescribeNode(e.frame, dom.DescribeNodeArgs{
		ObjectId: e.remote.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	var frame *Frame
	e.frame.manager.edit(val.Node.FrameId, func(f *Frame) {
		f.remote = e.remote
		frame = f
	})
	if frame == nil {
		return nil, fmt.Errorf("frame with id = %s not found", val.Node.FrameId)
	}
	return frame, nil
}

func stringArguments(val ...string) []*runtime.CallArgument {
	var args = make([]*runtime.CallArgument, len(val))
	for i, a := range val {
		args[i] = &runtime.CallArgument{Value: a}
	}
	return args
}

func (e Element) dispatchEvents(events ...string) error {
	_, err := e.callFunction(functionDispatchEvents, true, false, stringArguments(events...))
	return err
}

func (e Element) ScrollIntoView() error {
	return dom.ScrollIntoViewIfNeeded(e.frame, dom.ScrollIntoViewIfNeededArgs{ObjectId: e.remote.ObjectId})
}

func (e Element) GetText() (string, error) {
	v, err := e.callFunction(functionGetText, true, false, nil)
	if err != nil {
		return "null", err
	}
	return fmt.Sprint(v.Value), nil
}

func (e Element) Clear() error {
	_, err := e.callFunction(functionClearText, true, false, nil)
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
	if err = input.InsertText(e.frame, input.InsertTextArgs{Text: text}); err != nil {
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
	metric, err := page.GetLayoutMetrics(e.frame)
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
	if _, err = e.callFunction(functionPreventMissClick, true, false, nil); err != nil {
		return err
	}
	var mouse = e.frame.Session().Mouse()
	if err = mouse.Move(MouseNone, x, y); err != nil {
		return err
	}
	if err = mouse.Press(button, x, y); err != nil {
		return err
	}
	if err = mouse.Release(button, x, y); err != nil {
		return err
	}
	clicked, err := e.callFunction(functionClickDone, true, false, nil)
	if err != nil {
		if err == ErrStaleElementReference || err == ErrSessionClosed {
			return nil
		}
		return err
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
	return e.frame.Session().Mouse().Move(MouseNone, x, y)
}

func (e Element) SetAttribute(attr string, value string) error {
	_, err := e.callFunction(functionSetAttr, true, false, stringArguments(attr, value))
	return err
}

func (e Element) GetAttribute(attr string) (string, error) {
	return e.callFunctionStringValue(functionGetAttr, []*runtime.CallArgument{{Value: attr}})
}

func (e Element) Checkbox(check bool) error {
	if _, err := e.callFunction(functionCheckbox, true, false, []*runtime.CallArgument{{Value: check}}); err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e *Element) IsChecked() (bool, error) {
	v, err := e.callFunction(functionIsChecked, true, false, nil)
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
	return e.callFunctionStringValue(functionGetComputedStyle, []*runtime.CallArgument{{Value: style}})
}

func (e Element) SelectValues(values ...string) error {
	if e.remote.ClassName != "HTMLSelectElement" {
		return fmt.Errorf("can't use element as SELECT, not applicable type %s", e.remote.ClassName)
	}
	_, err := e.callFunction(functionSelect, true, false, []*runtime.CallArgument{{Value: values}})
	if err != nil {
		return err
	}
	return e.dispatchEvents(WebEventClick, WebEventInput, WebEventChange)
}

func (e Element) GetSelectedValues() ([]string, error) {
	return e.callFunctionStringArrayValue(functionGetSelectedValues, nil)
}

//func (e Element) GetSelectedText() ([]string, error) {
//	return e.callFunctionStringArrayValue(functionGetSelectedValues, nil)
//}

func (e Element) callFunctionStringValue(function string, args []*runtime.CallArgument) (string, error) {
	v, err := e.callFunction(function, true, false, args)
	if err != nil {
		return "", err
	}
	return remoteObjectPrimitive(*v).String()
}

func (e Element) callFunctionStringArrayValue(function string, args []*runtime.CallArgument) ([]string, error) {
	v, err := e.callFunction(function, true, false, args)
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
