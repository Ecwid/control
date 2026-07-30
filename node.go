package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ecwid/control/key"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

type clkAck struct {
	ID    string `json:"id"`
	Error string `json:"error,omitempty"`
}

func parseAck(payload string) (clkAck, error) {
	var ack clkAck
	if payload == "" {
		return ack, errors.New("empty payload")
	}
	if err := json.Unmarshal([]byte(payload), &ack); err != nil {
		return ack, err
	}
	if ack.ID == "" {
		return ack, errors.New("missing id in payload")
	}
	return ack, nil
}

type (
	NodeNotClickableError string
	NodeNotFocusableError string
	NodeNotVisibleError   string
	NoSuchSelectorError   string
)

func (n NodeNotClickableError) Error() string {
	return fmt.Sprintf("selector %s is not clickable", string(n))
}

func (n NodeNotVisibleError) Error() string {
	return fmt.Sprintf("selector %s is not visible", string(n))
}

func (n NodeNotFocusableError) Error() string {
	return fmt.Sprintf("selector %s is not focusable", string(n))
}

func (s NoSuchSelectorError) Error() string {
	return fmt.Sprintf("no such selector: %s", string(s))
}

type Node struct {
	object            runtime.RemoteObjectId
	requestedSelector string
	frame             *Frame
}

type NodeList []*Node

func (nl NodeList) Foreach(predicate func(*Node) error) error {
	for _, node := range nl {
		if err := predicate(node); err != nil {
			return err
		}
	}
	return nil
}

func (e Node) OwnerFrame() *Frame {
	return e.frame
}

func (e Node) Call(method string, send, recv any) error {
	return e.frame.Call(method, send, recv)
}

func (e Node) IsConnected() bool {
	isConnected, err := e.eval(`function(){return this.isConnected}`)
	if err != nil {
		return false
	}
	if value, ok := isConnected.(bool); ok {
		return value
	}
	return false
}

func (e Node) callFunctionOn(function string, awaitPromise bool, args ...any) (any, error) {
	arguments := make([]*runtime.CallArgument, 0, len(args))
	for _, arg := range args {
		carg := &runtime.CallArgument{}
		if roId, ok := arg.(runtime.RemoteObjectId); ok {
			carg.ObjectId = roId
		} else {
			carg.Value = arg
		}
		arguments = append(arguments, carg)
	}
	if len(arguments) == 0 {
		arguments = nil
	}
	value, err := runtime.CallFunctionOn(e.frame, runtime.CallFunctionOnArgs{
		FunctionDeclaration: function,
		ObjectId:            e.object,
		AwaitPromise:        awaitPromise,
		Arguments:           arguments,
		SerializationOptions: &runtime.SerializationOptions{
			Serialization: "deep",
		},
	})
	if err != nil {
		return nil, err
	}
	if hasException(value.ExceptionDetails) {
		return nil, RuntimeExceptionError{value: value.ExceptionDetails}
	}
	return e.frame.unserialize(value.Result)
}

func (e Node) eval(function string, args ...any) (any, error) {
	return e.callFunctionOn(function, true, args...)
}

func (e Node) dispatchEvents(events ...any) error {
	_, err := e.eval(`function(l){for(const e of l)this.dispatchEvent(new Event(e,{'bubbles':!0}))}`, events)
	return err
}

func (e Node) Log(msg string, args ...any) {
	args = append(args, "self", e.requestedSelector)
	e.frame.Log(msg, args...)
}

func (e Node) HasClass(class string) Optional[bool] {
	return optional[bool](e.eval(`function(c){return this.classList.contains(c)}`, class))
}

func (e Node) CallFunctionOn(function string, args ...any) Optional[any] {
	return optional[any](e.eval(function, args...))
}

func (e Node) Query(cssSelector string) Optional[*Node] {
	value, err := e.eval(`function(s){return this.querySelector(s)}`, cssSelector)
	if err != nil {
		return Optional[*Node]{err: err}
	}
	if value == nil {
		return Optional[*Node]{err: NoSuchSelectorError(cssSelector)}
	}
	if node, ok := value.(*Node); ok {
		node.requestedSelector = cssSelector
		return Optional[*Node]{value: node}
	}
	return Optional[*Node]{err: fmt.Errorf("interface conversion failed: got %T, want *Node", value)}
}

func (e Node) QueryAll(cssSelector string) Optional[NodeList] {
	value, err := e.eval(`function(s){return this.querySelectorAll(s)}`, cssSelector)
	if err != nil {
		return Optional[NodeList]{err: err}
	}
	if value == nil {
		return Optional[NodeList]{err: NoSuchSelectorError(cssSelector)}
	}
	if nodeList, ok := value.(NodeList); ok {
		return Optional[NodeList]{value: nodeList}
	}
	return Optional[NodeList]{err: fmt.Errorf("interface conversion failed: got %T, want NodeList", value)}
}

func (e Node) ContentFrame() Optional[*Frame] {
	value, err := e.frame.describeNode(e)
	if err != nil {
		return Optional[*Frame]{err: err}
	}
	return Optional[*Frame]{value: &Frame{
		id:      value.FrameId,
		session: e.frame.session,
		parent:  e.frame,
		node:    &e,
	}}
}

func (e Node) scrollIntoView() error {
	return dom.ScrollIntoViewIfNeeded(e, dom.ScrollIntoViewIfNeededArgs{ObjectId: e.object})
}

func (e Node) ScrollIntoView() error {
	return e.scrollIntoView()
}

func (e Node) GetText() Optional[string] {
	return optional[string](e.eval(`function(){return ('INPUT'===this.nodeName||'TEXTAREA'===this.nodeName)?this.value:this.innerText}`))
}

func (e Node) Focus() error {
	err := dom.Focus(e, dom.FocusArgs{ObjectId: e.object})
	if err != nil && err.Error() == `Element is not focusable` {
		return NodeNotFocusableError(e.requestedSelector)
	}
	return err
}

func (e Node) Blur() error {
	_, err := e.eval(`function(){this.blur()}`)
	return err
}

func (e Node) clearInput() error {
	_, err := e.eval(`function(){('INPUT'===this.nodeName||'TEXTAREA'===this.nodeName)?this.select():this.innerText=''}`)
	if err != nil {
		return err
	}
	return NewKeyboard(e.frame.session).Press(key.Keys[key.Backspace], time.Millisecond*85)
}

func (e Node) InsertText(value string) error {
	return e.setText(value, false)
}

func (e Node) SetText(value string) error {
	return e.setText(value, true)
}

func (e Node) setText(value string, clearBefore bool) (err error) {
	if err = e.Focus(); err != nil {
		return err
	}
	if clearBefore {
		if err = e.clearInput(); err != nil {
			return err
		}
	}
	if err = NewKeyboard(e.frame.session).Insert(value); err != nil {
		return err
	}
	return nil
}

func (e Node) CheckVisibility() Optional[bool] {
	return optional[bool](e.eval(`function(){return this.checkVisibility({opacityProperty: false, visibilityProperty: true})}`))
}

func (e Node) Upload(files ...string) error {
	return dom.SetFileInputFiles(e, dom.SetFileInputFilesArgs{
		ObjectId: e.object,
		Files:    files,
	})
}

func (e Node) pointerAction(eventName string, preventDefault bool, dispatch func(Point) error) (err error) {
	if err = e.scrollIntoView(); err != nil {
		return err
	}
	point, err := e.middle()
	if err != nil {
		return err
	}
	actionID := fmt.Sprintf("%d", time.Now().UnixNano())

	futureBindingCalled := subscribeToMethod(e.frame.session, "Runtime.bindingCalled", func(value runtime.BindingCalled) (bool, error) {
		if value.Name == hitCheckFunc {
			ack, err := parseAck(value.Payload)
			return ack.ID == actionID, err
		}
		return false, nil
	})
	defer futureBindingCalled.Cancel()

	/*
		const dist = `function(func, actionId, eventName, preventDefault) {
			const notify = window[func]
			const notifyResult = (error) => {
				notify(JSON.stringify({ id: actionId, error: error ?? null }))
			}

			const isSelfOrDescendant = (target) => {
				for (let node = target; node; node = node.parentNode) {
					if (node === this) {
						return true
					}
				}
				return false
			}

			const onPointerEvent = (event) => {
				if (event.isTrusted && isSelfOrDescendant(event.target)) {
					notifyResult(null)
					return
				}
				if (preventDefault) {
					event.preventDefault()
				}
				event.stopImmediatePropagation()
				notifyResult('target overlapped')
			}

			this.ownerDocument.addEventListener(eventName, onPointerEvent, { capture: true, once: true })
			window.addEventListener("beforeunload", () => notifyResult(null), { once: true })
		}`
	*/

	const minified = `function(e,t,n,r){const o=window[e],i=e=>{o(JSON.stringify({id:t,error:e??null}))},a=e=>{for(let t=e;t;t=t.parentNode)if(t===this)return!0;return!1};this.ownerDocument.addEventListener(n,e=>{e.isTrusted&&a(e.target)?i(null):(r&&e.preventDefault(),e.stopImmediatePropagation(),i("target overlapped"))},{capture:!0,once:!0}),window.addEventListener("beforeunload",()=>i(null),{once:!0})}`
	_, err = e.eval(minified, hitCheckFunc, actionID, eventName, preventDefault)
	if err != nil {
		return err
	}

	if err = dispatch(point); err != nil {
		return err
	}

	call, err := GetWithTimeout(e.frame.session, futureBindingCalled)
	if err != nil {
		return err
	}
	ack, err := parseAck(call.Payload)
	if err != nil {
		return err
	}
	if ack.Error != "" {
		return errors.New(ack.Error)
	}
	return nil
}

func (e Node) Click() (err error) {
	return e.pointerAction("click", true, e.frame.session.Click)
}

func (e Node) Down() (err error) {
	return e.pointerAction("mousedown", false, e.frame.session.MouseDown)
}

func (e Node) GetClickablePoint() Optional[Point] {
	return optional[Point](e.middle())
}

func (e Node) middle() (middle Point, err error) {
	value, err := e.CheckVisibility().Unwrap()
	if err != nil {
		return middle, err
	}
	if !value {
		return middle, NodeNotVisibleError(e.requestedSelector)
	}
	var r0 Quad
	r0, err = e.getContentQuad()
	if err != nil {
		return middle, err
	}
	return r0.Middle(), nil
}

func (e Node) GetBoundingClientRect() Optional[Rectangle] {
	return optional[Rectangle](e.getBoundingClientRect())
}

func (e Node) getBoundingClientRect() (Rectangle, error) {
	value, err := e.eval(`function() {
		const e = this.getBoundingClientRect()
		const t = this.ownerDocument.documentElement.getBoundingClientRect()
		return [e.left - t.left, e.top - t.top, e.width, e.height]
	}`)
	if err != nil {
		return Rectangle{}, err
	}
	if arr, ok := value.([]any); ok {
		return Rectangle{
			X:      arr[0].(float64),
			Y:      arr[1].(float64),
			Width:  arr[2].(float64),
			Height: arr[3].(float64),
		}, nil
	}
	return Rectangle{}, errors.New("getBoundingClientRect: eval result is not an array")
}

func (e Node) getContentQuad() (Quad, error) {
	val, err := dom.GetContentQuads(e, dom.GetContentQuadsArgs{
		ObjectId: e.object,
	})
	if err != nil {
		return nil, err
	}
	quads := convertQuads(val.Quads)
	if len(quads) == 0 {
		return nil, errors.New("node has no visible bounds")
	}
	for _, quad := range quads {
		if quad.Area() > 1 {
			return quad, nil
		}
	}
	return nil, errors.New("node bounds have no size")
}

func (e Node) Hover() error {
	if err := e.scrollIntoView(); err != nil {
		return err
	}
	p, err := e.middle()
	if err != nil {
		return err
	}
	return e.frame.session.Hover(p)
}

func (e Node) GetComputedStyle(style string, pseudo string) Optional[string] {
	var pseudoVar any = nil
	if pseudo != "" {
		pseudoVar = pseudo
	}
	return optional[string](e.eval(`function(p,s){return getComputedStyle(this, p)[s]}`, pseudoVar, style))
}

func (e Node) SetAttribute(attr, value string) error {
	_, err := e.eval(`function(a,v){this.setAttribute(a,v)}`, attr, value)
	return err
}

func (e Node) GetAttribute(attr string) Optional[string] {
	return optional[string](e.eval(`function(a){return this.getAttribute(a)}`, attr))
}

func (e Node) GetRectangle() Optional[Rectangle] {
	return optional[Rectangle](e.getViewportRectangle())
}

func (e Node) getViewportRectangle() (Rectangle, error) {
	q, err := e.getContentQuad()
	if err != nil {
		return Rectangle{}, err
	}
	rect := Rectangle{
		X:      q[0].X,
		Y:      q[0].Y,
		Width:  q[1].X - q[0].X,
		Height: q[3].Y - q[0].Y,
	}
	return rect, nil
}

func (e Node) SelectByValues(values ...string) error {
	_, err := e.eval(`function(a){const b=Array.from(this.options);this.value=void 0;for(const c of b)if(c.selected=a.includes(c.value),c.selected&&!this.multiple)break}`, values)
	if err != nil {
		return err
	}
	return e.dispatchEvents("click", "input", "change")
}

func (e Node) GetSelected(textContent bool) Optional[[]string] {
	return optional[[]string](e.getSelected(textContent))
}

func (e Node) getSelected(textContent bool) ([]string, error) {
	values, err := e.eval(`function(text){return Array.from(this.options).filter(a=>a.selected).map(a=>text?a.textContent.trim():a.value)}`, textContent)
	if err != nil {
		return nil, err
	}
	stringsValues := make([]string, len(values.([]any)))
	for n, val := range values.([]any) {
		stringsValues[n] = val.(string)
	}
	return stringsValues, nil
}

func (e Node) SetCheckbox(check bool) error {
	_, err := e.eval(`function(v){this.checked=v}`, check)
	if err != nil {
		return err
	}
	return e.dispatchEvents("click", "input", "change")
}

func (e Node) IsChecked() Optional[bool] {
	return optional[bool](e.eval(`function(){return this.checked}`))
}
