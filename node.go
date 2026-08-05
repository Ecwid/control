package control

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/ecwid/control/key"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/domdebugger"
	"github.com/ecwid/control/protocol/runtime"
)

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
		context: e.frame.session.getOrCreateFrameContext(value.FrameId),
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

func (e Node) hasClickListener() (bool, error) {
	val, err := domdebugger.GetEventListeners(e.frame.session, domdebugger.GetEventListenersArgs{ObjectId: e.object})
	if err != nil {
		return false, err
	}
	for _, listener := range val.Listeners {
		if listener != nil && listener.Type == "click" {
			return true, nil
		}
	}
	return false, nil
}

func (e Node) HasClickListener() Optional[bool] {
	return optional[bool](e.hasClickListener())
}

func (e Node) isStableAfterAnimationFrame() (bool, error) {
	value, err := e.eval(`function() {
		const isSame = (a, b) => a.x === b.x && a.y === b.y && a.width === b.width && a.height === b.height
		const readRectIfConnected = () => this.isConnected ? this.getBoundingClientRect() : null

		const initial = readRectIfConnected()
		if (!initial) {
			return false
		}

		return new Promise(resolve => {
			requestAnimationFrame(() => {
				const first = readRectIfConnected()
				if (!first || !isSame(first, initial)) {
					resolve(false)
					return
				}

				requestAnimationFrame(() => {
					const second = readRectIfConnected()
					resolve(!!second && isSame(second, first))
				})
			})
		})
	}`)
	if err != nil {
		return false, err
	}
	stable, ok := value.(bool)
	if !ok {
		return false, errors.New("requestAnimationFrame stability check result is not a bool")
	}
	return stable, nil
}

func (e Node) receivesEventsAt(point Point) (bool, error) {
	value, err := e.eval(`function(x, y) {
		const target = this.ownerDocument.elementFromPoint(x, y)
		if (!target) {
			return false
		}
		for (let node = target; node; node = node.parentNode) {
			if (node === this) {
				return true
			}
		}
		return false
	}`, point.X, point.Y)
	if err != nil {
		return false, err
	}
	receivesEvents, ok := value.(bool)
	if !ok {
		return false, errors.New("receive events check result is not a bool")
	}
	return receivesEvents, nil
}

func (e Node) contentOffsetInParentViewport() (Point, error) {
	/*
		function() {
			const rect = this.getBoundingClientRect()
			const style = this.ownerDocument.defaultView.getComputedStyle(this)
			const left = rect.left + parseFloat(style.borderLeftWidth || "0") + parseFloat(style.paddingLeft || "0")
			const top = rect.top + parseFloat(style.borderTopWidth || "0") + parseFloat(style.paddingTop || "0")
			return [left, top]
		}
	*/
	const script = `function(){const t=this.getBoundingClientRect(),e=this.ownerDocument.defaultView.getComputedStyle(this);return[t.left+parseFloat(e.borderLeftWidth||"0")+parseFloat(e.paddingLeft||"0"),t.top+parseFloat(e.borderTopWidth||"0")+parseFloat(e.paddingTop||"0")]}`
	value, err := e.eval(script)
	if err != nil {
		return Point{}, err
	}
	arr, ok := value.([]any)
	if !ok || len(arr) < 2 {
		return Point{}, errors.New("frame content offset result is not a 2-element array")
	}
	left, okX := arr[0].(float64)
	top, okY := arr[1].(float64)
	if !okX || !okY {
		return Point{}, errors.New("frame content offset values are not numbers")
	}
	return Point{X: left, Y: top}, nil
}

func (e Node) toOwnerDocumentPoint(pagePoint Point) (Point, error) {
	localPoint := pagePoint
	for frame := e.frame; frame != nil && frame.parent != nil; frame = frame.parent {
		if frame.node == nil {
			return Point{}, errors.New("frame parent chain is broken: missing owner iframe node")
		}
		offset, err := frame.node.contentOffsetInParentViewport()
		if err != nil {
			return Point{}, err
		}
		localPoint.X -= offset.X
		localPoint.Y -= offset.Y
	}
	return localPoint, nil
}

func (e Node) setHitTargetInterceptor(eventName string) (runtime.RemoteObjectId, error) {
	const script = `function(eventName) {
		let resolved = false
		let observer = null

		return new Promise(done => {
			const finish = (error) => {
				if (resolved) {
					return
				}
				resolved = true
				if (observer) {
					observer.disconnect()
				}
				window.removeEventListener("beforeunload", onBeforeUnload)
				this.ownerDocument.removeEventListener(eventName, listener, true)
				done(error ?? null)
			}

			const onBeforeUnload = () => finish(null)

			const isSelfOrDescendant = (target) => {
				for (let node = target; node; node = node.parentNode) {
					if (node === this) {
						return true
					}
				}
				return false
			}

			const listener = (e) => {
				if (resolved) {
					return
				}
				if (e.isTrusted && isSelfOrDescendant(e.target)) {
					finish(null)
					return
				}
				e.preventDefault()
				e.stopImmediatePropagation()
				finish("target overlapped")
			}

			observer = new MutationObserver(() => {
				if (!this.isConnected) {
					finish("target disconnected")
				}
			})

			observer.observe(this.ownerDocument, { childList: true, subtree: true })
			window.addEventListener("beforeunload", onBeforeUnload, { once: true })
			this.ownerDocument.addEventListener(eventName, listener, { capture: true, once: true })
		})
	}`

	result, err := e.callFunctionOn(script, false, eventName)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", errors.New("hit-target interceptor promise object is missing")
	}
	if objectId, ok := result.(runtime.RemoteObjectId); ok {
		return objectId, nil
	}
	return "", errors.New("unexpected hit-target interceptor result type")
}

func isFrameContextGone(err error) bool {
	return err != nil && slices.Contains([]string{errCannotFindContext, errCannotFindObject}, err.Error())
}

func (e Node) dispatchPointerEvent(event string, dispatchFunc func(Point) error) (err error) {
	if err = e.scrollIntoView(); err != nil {
		return err
	}

	point, err := e.middle()
	if err != nil {
		return err
	}

	stable, err := e.isStableAfterAnimationFrame()
	if err != nil {
		return err
	}
	if !stable {
		return errors.New("element changed after requestAnimationFrame")
	}

	pointInOwnerDocument, err := e.toOwnerDocumentPoint(point)
	if err != nil {
		return err
	}

	receivesEvents, err := e.receivesEventsAt(pointInOwnerDocument)
	if err != nil {
		return err
	}
	if !receivesEvents {
		return errors.New("element does not receive events at its center point")
	}

	promise, err := e.setHitTargetInterceptor(event)
	if err != nil {
		return err
	}
	contextRevisionBeforeDispatch := e.frame.contextRevision()

	if err = dispatchFunc(point); err != nil {
		return err
	}

	ackValue, err := e.frame.AwaitPromise(promise)
	if err != nil {
		if isFrameContextGone(err) && e.frame.contextChangedSince(contextRevisionBeforeDispatch) {
			return nil
		}
		return err
	}
	if ackValue == nil {
		return nil
	}
	if ackError, ok := ackValue.(string); ok {
		if ackError != "" {
			return errors.New(ackError)
		}
		return nil
	}
	return errors.New("unexpected hit-target interceptor result type")
}

func (e Node) Click() (err error) {
	return e.dispatchPointerEvent("click", e.frame.session.Click)
}

func (e Node) Down() (err error) {
	return e.dispatchPointerEvent("mousedown", e.frame.session.MouseDown)
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
		const nr = this.getBoundingClientRect()
		const dr = this.ownerDocument.documentElement.getBoundingClientRect()
		return [nr.left - dr.left, nr.top - dr.top, nr.width, nr.height]
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
