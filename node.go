package control

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ecwid/control/key"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/overlay"
	"github.com/ecwid/control/protocol/runtime"
)

type (
	NodeNonClickableError string
	NodeNonFocusableError string
	NodeInvisibleError    string
	NodeUnstableError     string
	NoSuchSelectorError   string
)

func (n NodeNonClickableError) Error() string {
	return fmt.Sprintf("selector `%s` is not clickable", string(n))
}

func (n NodeInvisibleError) Error() string {
	return fmt.Sprintf("selector `%s` is not visible", string(n))
}

func (n NodeUnstableError) Error() string {
	return fmt.Sprintf("selector `%s` is not stable", string(n))
}

func (n NodeNonFocusableError) Error() string {
	return fmt.Sprintf("selector `%s` is not focusable", string(n))
}

func (s NoSuchSelectorError) Error() string {
	return fmt.Sprintf("no such selector found: `%s`", string(s))
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	object            RemoteObject
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

type Point struct {
	X float64
	Y float64
}

type Quad []Point

func (p Point) Equal(a Point) bool {
	return p.X == a.X && p.Y == a.Y
}

func convertQuads(dq []dom.Quad) []Quad {
	var p = make([]Quad, len(dq))
	for n, q := range dq {
		p[n] = Quad{
			Point{q[0], q[1]},
			Point{q[2], q[3]},
			Point{q[4], q[5]},
			Point{q[6], q[7]},
		}
	}
	return p
}

// Middle calc middle of quad
func (q Quad) Middle() Point {
	x := 0.0
	y := 0.0
	for i := 0; i < 4; i++ {
		x += q[i].X
		y += q[i].Y
	}
	return Point{X: x / 4, Y: y / 4}
}

func (q Quad) Area() float64 {
	var area float64
	var x1, x2, y1, y2 float64
	var vertices = len(q)
	for i := 0; i < vertices; i++ {
		x1 = q[i].X
		y1 = q[i].Y
		x2 = q[(i+1)%vertices].X
		y2 = q[(i+1)%vertices].Y
		area += (x1*y2 - x2*y1) / 2
	}
	return math.Abs(area)
}

func (e Node) Highlight() error {
	return overlay.HighlightNode(e.frame, overlay.HighlightNodeArgs{
		HighlightConfig: &overlay.HighlightConfig{
			GridHighlightConfig: &overlay.GridHighlightConfig{
				RowGapColor:      &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				RowHatchColor:    &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
				ColumnGapColor:   &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				ColumnHatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
				RowLineColor:     &dom.RGBA{R: 127, G: 32, B: 210},
				ColumnLineColor:  &dom.RGBA{R: 127, G: 32, B: 210},
				RowLineDash:      true,
				ColumnLineDash:   true,
			},
			FlexContainerHighlightConfig: &overlay.FlexContainerHighlightConfig{
				ContainerBorder: &overlay.LineStyle{
					Color:   &dom.RGBA{R: 127, G: 32, B: 210},
					Pattern: "dashed",
				},
				ItemSeparator: &overlay.LineStyle{
					Color:   &dom.RGBA{R: 127, G: 32, B: 210},
					Pattern: "dashed",
				},
				LineSeparator: &overlay.LineStyle{
					Color:   &dom.RGBA{R: 127, G: 32, B: 210},
					Pattern: "dashed",
				},
				MainDistributedSpace: &overlay.BoxStyle{
					HatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
					FillColor:  &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				},
				CrossDistributedSpace: &overlay.BoxStyle{
					HatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
					FillColor:  &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				},
				RowGapSpace: &overlay.BoxStyle{
					HatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
					FillColor:  &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				},
				ColumnGapSpace: &overlay.BoxStyle{
					HatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
					FillColor:  &dom.RGBA{R: 127, G: 32, B: 210, A: 0.3},
				},
			},
			FlexItemHighlightConfig: &overlay.FlexItemHighlightConfig{
				BaseSizeBox: &overlay.BoxStyle{
					HatchColor: &dom.RGBA{R: 127, G: 32, B: 210, A: 0.8},
				},
				BaseSizeBorder: &overlay.LineStyle{
					Color:   &dom.RGBA{R: 127, G: 32, B: 210},
					Pattern: "dotted",
				},
				FlexibilityArrow: &overlay.LineStyle{
					Color: &dom.RGBA{R: 127, G: 32, B: 210},
				},
			},
			ContrastAlgorithm: overlay.ContrastAlgorithm("aa"),
			ContentColor:      &dom.RGBA{R: 111, G: 168, B: 220, A: 0.66},
			PaddingColor:      &dom.RGBA{R: 147, G: 196, B: 125, A: 0.55},
			BorderColor:       &dom.RGBA{R: 255, G: 229, B: 153, A: 0.66},
			MarginColor:       &dom.RGBA{R: 246, G: 178, B: 107, A: 0.66},
			EventTargetColor:  &dom.RGBA{R: 255, G: 196, B: 196, A: 0.66},
			ShapeColor:        &dom.RGBA{R: 96, G: 82, B: 177, A: 0.8},
			ShapeMarginColor:  &dom.RGBA{R: 96, G: 82, B: 127, A: 0.6},
		},
		ObjectId: e.GetRemoteObjectID(),
	})
}

func (e Node) GetRemoteObjectID() runtime.RemoteObjectId {
	return e.object.GetRemoteObjectID()
}

func (e Node) OwnerFrame() *Frame {
	return e.frame
}

func (e Node) Call(method string, send, recv any) error {
	return e.frame.Call(method, send, recv)
}

func (e Node) IsConnected() bool {
	value, err := e.eval(`function(){return this.isConnected}`)
	if err != nil {
		return false
	}
	return value.(bool)
}

func (e Node) MustReleaseObject() {
	panicIfError(e.ReleaseObject())
}

func (e Node) ReleaseObject() error {
	err := runtime.ReleaseObject(e, runtime.ReleaseObjectArgs{ObjectId: e.GetRemoteObjectID()})
	if err != nil && err.Error() == `Cannot find context with specified id` {
		return nil
	}
	return err
}

func (e Node) eval(function string, args ...any) (any, error) {
	return e.frame.CallFunctionOn(e, function, true, args...)
}

func (e Node) asyncEval(function string, args ...any) (RemoteObject, error) {
	value, err := e.frame.CallFunctionOn(e, function, false, args...)
	if err != nil {
		return nil, err
	}
	if v, ok := value.(RemoteObject); ok {
		return v, nil
	}
	return nil, fmt.Errorf("interface conversion failed, `%+v` not JsObject", value)
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
	return optional[bool](e.eval(`function(c){return this.classList.contains(c)}`))
}

func (e Node) MustHasClass(class string) bool {
	return e.HasClass(class).MustGetValue()
}

func (e Node) CallFunctionOn(function string, args ...any) Optional[any] {
	return optional[any](e.eval(function, args...))
}

func (e Node) MustCallFunctionOn(function string, args ...any) any {
	return e.CallFunctionOn(function, args...).MustGetValue()
}

func (e Node) AsyncCallFunctionOn(function string, args ...any) Optional[RemoteObject] {
	return optional[RemoteObject](e.asyncEval(function, args...))
}

func (e Node) MustAsyncCallFunctionOn(function string, args ...any) RemoteObject {
	return e.AsyncCallFunctionOn(function, args...).MustGetValue()
}

func (e Node) Query(cssSelector string) Optional[*Node] {
	return optional[*Node](e.query(cssSelector))
}

func (e Node) MustQuery(cssSelector string) *Node {
	return e.Query(cssSelector).MustGetValue()
}

func (e Node) query(cssSelector string) (*Node, error) {
	value, err := e.eval(`function(s){return this.querySelector(s)}`, cssSelector)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, NoSuchSelectorError(cssSelector)
	}
	node := value.(*Node)
	if e.frame.session.highlightEnabled {
		_ = node.Highlight()
	}
	node.requestedSelector = cssSelector
	return node, nil
}

func (e Node) QueryAll(cssSelector string) Optional[NodeList] {
	value, err := e.eval(`function(s){return this.querySelectorAll(s)}`, cssSelector)
	if err == nil && value == nil {
		err = NoSuchSelectorError(cssSelector)
	}
	return optional[NodeList](value, err)
}

func (e Node) MustQueryAll(cssSelector string) NodeList {
	return e.QueryAll(cssSelector).MustGetValue()
}

func (e Node) ContentFrame() Optional[*Frame] {
	return optional[*Frame](e.contentFrame())
}

func (e Node) MustContentFrame() *Frame {
	return e.ContentFrame().MustGetValue()
}

func (e *Node) contentFrame() (*Frame, error) {
	value, err := e.frame.describeNode(e)
	if err != nil {
		return nil, err
	}
	return &Frame{
		id:      value.FrameId,
		session: e.frame.session,
		parent:  e.frame,
		node:    e,
	}, nil
}

func (e Node) scrollIntoView() error {
	return dom.ScrollIntoViewIfNeeded(e, dom.ScrollIntoViewIfNeededArgs{ObjectId: e.GetRemoteObjectID()})
}

func (e Node) ScrollIntoView() error {
	return e.scrollIntoView()
}

func (e Node) MustScrollIntoView() {
	panicIfError(e.ScrollIntoView())
}

func (e Node) GetText() Optional[string] {
	return optional[string](e.eval(`function(){return ('INPUT'===this.nodeName||'TEXTAREA'===this.nodeName)?this.value:this.innerText}`))
}

func (e Node) MustGetText() string {
	return e.GetText().MustGetValue()
}

func (e Node) Focus() error {
	err := dom.Focus(e, dom.FocusArgs{ObjectId: e.GetRemoteObjectID()})
	if err != nil && err.Error() == `Element is not focusable` {
		err = NodeNonFocusableError(e.requestedSelector)
	}
	return err
}

func (e Node) MustFocus() {
	panicIfError(e.Focus())
}

func (e Node) Blur() error {
	_, err := e.eval(`function(){this.blur()}`)
	return err
}

func (e Node) MustBlur() {
	panicIfError(e.Blur())
}

func (e Node) clearInput() error {
	_, err := e.eval(`function(){('INPUT'===this.nodeName||'TEXTAREA'===this.nodeName)?this.select():this.innerText=''}`)
	if err != nil {
		return err
	}
	return e.frame.session.kb.Press(key.Keys[key.Backspace], time.Millisecond*85)
}

func (e Node) InsertText(value string) error {
	return e.setText(value, false)
}

func (e Node) MustInsertText(value string) {
	panicIfError(e.InsertText(value))
}

func (e Node) SetText(value string) error {
	return e.setText(value, true)
}

func (e Node) MustSetText(value string) {
	panicIfError(e.SetText(value))
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
	if err = e.frame.session.kb.Insert(value); err != nil {
		return err
	}
	return nil
}

func (e Node) MustCheckVisibility() bool {
	return e.CheckVisibility().MustGetValue()
}

func (e Node) CheckVisibility() Optional[bool] {
	value, err := e.eval(`function(){return this.checkVisibility({opacityProperty: false, visibilityProperty: true})}`)
	return optional[bool](value, err)
}

func (e Node) Upload(files ...string) error {
	return dom.SetFileInputFiles(e, dom.SetFileInputFilesArgs{
		ObjectId: e.GetRemoteObjectID(),
		Files:    files,
	})
}

func (e Node) MustUpload(files ...string) {
	panicIfError(e.Upload(files...))
}

func (e Node) Click() (err error) {
	if err = e.scrollIntoView(); err != nil {
		return err
	}
	point, err := e.clickablePoint()
	if err != nil {
		return err
	}

	future := e.frame.session.funcCalled(hitCheckFunc)
	defer future.Cancel()
	_, err = e.eval(`function(func) {
		let a = window[func],
			d = (b) => {
				for (let d = b; d; d = d.parentNode) {
					if (d === this) {
						return !0
					}
				}
				return !1
			},
			f = (b) => {
				if (b.isTrusted && d(b.target)) {
					a('')
				} else {
					b.stopImmediatePropagation()
					a('target overlapped')
				}
			}
		this.ownerDocument.addEventListener("click", f, { capture: true, once: true })
		window.addEventListener("beforeunload", () => a('document unloaded before click'))
	}`, hitCheckFunc)
	if err != nil {
		return err
	}
	if err = e.frame.session.Click(point); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(e.frame.session.context, e.frame.session.timeout)
	defer cancel()
	call, err := future.Get(ctx)
	if err != nil {
		return err
	}
	if call.Payload != "" {
		return errors.New(call.Payload)
	}
	return nil
}

func (e Node) MustClick() {
	panicIfError(e.Click())
}

func (e Node) Down() (err error) {
	if err = e.scrollIntoView(); err != nil {
		return err
	}
	point, err := e.clickablePoint()
	if err != nil {
		return err
	}
	future := e.frame.session.funcCalled(hitCheckFunc)
	defer future.Cancel()

	_, err = e.eval(`function(func) {
		let a = window[func],
			d = (b) => {
				for (let d = b; d; d = d.parentNode) {
					if (d === this) {
						return !0
					}
				}
				return !1
			},
			f = (b) => {
				if (b.isTrusted && d(b.target)) {
					a('')
				} else {
					b.stopImmediatePropagation()
					a('target overlapped')
				}
			}
		this.ownerDocument.addEventListener("mousedown", f, { capture: true, once: true })
	}`, hitCheckFunc)

	if err = e.frame.session.MouseDown(point); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(e.frame.session.context, e.frame.session.timeout)
	defer cancel()
	call, err := future.Get(ctx)
	if err != nil {
		return err
	}
	if call.Payload != "" {
		return errors.New(call.Payload)
	}
	return nil
}

func (e Node) MustDown() {
	panicIfError(e.Down())
}

func (e Node) GetClickablePoint() Optional[Point] {
	return optional[Point](e.clickablePoint())
}

func (e Node) MustGetClickablePoint() Point {
	return e.GetClickablePoint().MustGetValue()
}

func (e Node) clickablePoint() (middle Point, err error) {
	value, err := e.CheckVisibility().Unwrap()
	if err != nil {
		return middle, err
	}
	if !value {
		return middle, NodeInvisibleError(e.requestedSelector)
	}
	var r0, r1 Quad
	r0, err = e.getContentQuad()
	if err != nil {
		return middle, err
	}
	_, err = e.frame.evaluate(`new Promise(r => setTimeout(r,100))`, true)
	if err != nil {
		return middle, err
	}
	r1, err = e.getContentQuad()
	if err != nil {
		return middle, err
	}
	middle = r0.Middle()
	if middle.Equal(r1.Middle()) {
		return middle, nil
	}
	return middle, NodeUnstableError(e.requestedSelector)
}

func (e Node) GetBoundingClientRect() Optional[dom.Rect] {
	return optional[dom.Rect](e.getBoundingClientRect())
}

func (e Node) MustGetBoundingClientRect() dom.Rect {
	return e.GetBoundingClientRect().MustGetValue()
}

func (e Node) getBoundingClientRect() (dom.Rect, error) {
	value, err := e.eval(`function() {
		const e = this.getBoundingClientRect()
		const t = this.ownerDocument.documentElement.getBoundingClientRect()
		return [e.left - t.left, e.top - t.top, e.width, e.height]
	}`)
	if err != nil {
		return dom.Rect{}, err
	}
	if arr, ok := value.([]any); ok {
		return dom.Rect{
			X:      arr[0].(float64),
			Y:      arr[1].(float64),
			Width:  arr[2].(float64),
			Height: arr[3].(float64),
		}, nil
	}
	return dom.Rect{}, errors.New("getBoundingClientRect: eval result is not array")
}

func (e Node) getContentQuad() (Quad, error) {
	val, err := dom.GetContentQuads(e, dom.GetContentQuadsArgs{
		ObjectId: e.GetRemoteObjectID(),
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
	p, err := e.clickablePoint()
	if err != nil {
		return err
	}
	return e.frame.session.Hover(p)
}

func (e Node) MustHover() {
	panicIfError(e.Hover())
}

func (e Node) GetComputedStyle(style string, pseudo string) Optional[string] {
	var pseudoVar any = nil
	if pseudo != "" {
		pseudoVar = pseudo
	}
	return optional[string](e.eval(`function(p,s){return getComputedStyle(this, p)[s]}`, pseudoVar, style))
}

func (e Node) MustGetComputedStyle(style string, pseudo string) string {
	return e.GetComputedStyle(style, pseudo).MustGetValue()
}

func (e Node) SetAttribute(attr, value string) error {
	_, err := e.eval(`function(a,v){this.setAttribute(a,v)}`, attr, value)
	return err
}

func (e Node) MustSetAttribute(attr, value string) {
	panicIfError(e.SetAttribute(attr, value))
}

func (e Node) GetAttribute(attr string) Optional[string] {
	return optional[string](e.eval(`function(a){return this.getAttribute(a)}`, attr))
}

func (e Node) MustGetAttribute(attr string) string {
	return e.GetAttribute(attr).MustGetValue()
}

func (e Node) GetRectangle() Optional[dom.Rect] {
	return optional[dom.Rect](e.getViewportRectangle())
}

func (e Node) MustGetRectangle() dom.Rect {
	return e.GetRectangle().MustGetValue()
}

func (e Node) getViewportRectangle() (dom.Rect, error) {
	q, err := e.getContentQuad()
	if err != nil {
		return dom.Rect{}, err
	}
	rect := dom.Rect{
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

func (e Node) MustSelectByValues(values ...string) {
	panicIfError(e.SelectByValues(values...))
}

func (e Node) GetSelected(textContent bool) Optional[[]string] {
	return optional[[]string](e.getSelected(textContent))
}

func (e Node) MustGetSelected(textContent bool) []string {
	return e.GetSelected(textContent).MustGetValue()
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

func (e Node) MustSetCheckbox(check bool) {
	panicIfError(e.SetCheckbox(check))
}

func (e Node) IsChecked() Optional[bool] {
	return optional[bool](e.eval(`function(){return this.checked}`))
}

func (e Node) MustIsChecked() bool {
	return e.IsChecked().MustGetValue()
}
