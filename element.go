package witness

import (
	"strings"
	"time"

	"github.com/ecwid/witness/internal/atom"
	"github.com/ecwid/witness/pkg/devtool"
)

// Element ...
type Element interface {
	Seek(string) (Element, error)
	SeekAll(string) []Element
	Expect(string, bool) Element

	Click() error
	Hover() error
	Type(string, ...rune) error
	Upload(...string) error
	Clear() error
	Select(...string) error
	Checkbox(bool) error
	SetAttr(string, string) error
	Call(string, ...interface{}) (interface{}, error)
	Focus() error
	SwitchToFrame() error

	IsVisible() (bool, error)
	GetText() (string, error)
	GetAttr(attr string) (string, error)
	GetRectangle() (*devtool.Rect, error)
	GetComputedStyle(string) (string, error)
	GetSelected(bool) ([]string, error)
	IsChecked() (bool, error)
	GetEventListeners() ([]string, error)

	Release() error
}

func (e *element) Release() error {
	defer e.detach()
	return e.session.releaseObject(e.ID)
}

func (e *element) detached() bool {
	return e.ID == ""
}

// Element ...
type element struct {
	session     *Session
	ID          string
	description string
	parent      *element
	childs      []*element
}

func newElement(s *Session, parent *element, ID string, description string) *element {
	e := &element{
		ID:          ID,
		session:     s,
		description: description,
		parent:      parent,
		childs:      make([]*element, 0),
	}
	if parent != nil {
		parent.childs = append(parent.childs, e)
	}
	return e
}

func (e *element) detach() {
	e.ID = ""
	for _, c := range e.childs {
		c.detach()
	}
}

func (e *element) renew() error {
	if !e.detached() {
		return nil
	}
	if e.parent == nil {
		// request a new document
		new, err := e.session.evaluate("document", e.session.contextID, false)
		if err != nil {
			return ErrStaleElementReference
		}
		e.ID = new.ObjectID
		return nil
	}
	e.session.client.Logging.Printf(LevelFatal, "todo: renew element is not implemented yet")
	return nil
}

func (e *element) Seek(selector string) (Element, error) {
	return e.findElement(selector)
}

// Expect searching selector (visible) with implicity wait timeout
func (e *element) Expect(selector string, visible bool) Element {
	el, err := e.session.Ticker(func() (interface{}, error) {
		new, err := e.Seek(selector)
		if err != nil {
			return nil, err
		}
		if visible {
			v, err := new.IsVisible()
			if err != nil {
				return nil, err
			}
			if !v {
				return nil, ErrElementInvisible
			}
		}
		return new, nil
	})
	if err != nil {
		panic(err)
	}
	return el.(Element)
}

func (e *element) SeekAll(selector string) []Element {
	v, err := e.findElements(selector)
	if err != nil {
		return []Element{}
	}
	return v
}

func (e *element) findElements(selector string) ([]Element, error) {
	if err := e.renew(); err != nil {
		return nil, err
	}
	selector = strings.ReplaceAll(selector, `"`, `\"`)
	ro, err := e.call(atom.QueryAll, selector)
	if err != nil {
		return nil, err
	}
	if ro == nil || ro.Description == "NodeList(0)" {
		e.session.releaseObject(ro.ObjectID)
		return nil, ErrNoSuchElement
	}
	els := make([]Element, 0)
	descriptor, err := e.session.getProperties(ro.ObjectID)
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		els = append(els, newElement(e.session, e, d.Value.ObjectID, d.Value.Description))
	}
	return els, nil
}

func (e *element) findElement(selector string) (Element, error) {
	if err := e.renew(); err != nil {
		return nil, err
	}
	selector = strings.ReplaceAll(selector, `"`, `\"`)
	element, err := e.call(atom.Query, selector)
	if err != nil {
		return nil, err
	}
	if element.Subtype == "null" {
		return nil, ErrNoSuchElement
	}
	return newElement(e.session, e, element.ObjectID, element.Description), nil
}

func (e *element) getNode() (*devtool.Node, error) {
	msg, err := e.session.blockingSend("DOM.describeNode", Map{
		"objectId": e.ID,
		"depth":    1,
	})
	if err != nil {
		return nil, err
	}
	describeNode := new(devtool.DescribeNode)
	if err = msg.Unmarshal(describeNode); err != nil {
		return nil, err
	}
	return describeNode.Node, nil
}

// Focus focus element
func (e *element) Focus() error {
	_, err := e.session.blockingSend("DOM.focus", Map{"objectId": e.ID})
	return err
}

func (e *element) call(functionDeclaration string, arg ...interface{}) (*devtool.RemoteObject, error) {
	if err := e.renew(); err != nil {
		return nil, err
	}
	return e.session.callFunctionOn(e.ID, functionDeclaration, arg...)
}

// Call ...
func (e *element) Call(functionDeclaration string, arg ...interface{}) (interface{}, error) {
	v, err := e.session.callFunctionOn(e.ID, functionDeclaration, arg...)
	if err != nil {
		return nil, err
	}
	return v.Value, nil
}

// Upload upload files
func (e *element) Upload(files ...string) error {
	_, err := e.session.blockingSend("DOM.setFileInputFiles", Map{
		"files":    files,
		"objectId": e.ID,
	})
	return err
}

func (e *element) clickablePoint() (x float64, y float64, err error) {
	r, err := e.session.getContentQuads(0, e.ID, true)
	if err != nil {
		return -1, -1, err
	}
	x, y = r.Middle()
	/*
		При определении элемента, по которому произойдет клик по координатам x, y, необходимо учесть случай,
		когда клик происходит во фрейме. Во фрейме метод elementFromPoint(x, y) работает относительно координат самого фрейма,
		а не абсолютных координат viewport браузера. Здесь мы вычисляем координаты фрейма и вычитаем их
		из значений x, y для теста клика.
	*/
	cX, cY := x, y
	if e.session.frameID != e.session.targetID {
		frameElement, err := e.session.getFrameOwner(e.session.frameID)
		if err != nil {
			return x, y, err
		}
		size, err := e.session.getContentQuads(frameElement, "", false)
		if err != nil {
			return x, y, err
		}
		cX -= size[0].X
		cY -= size[0].Y
	}
	// Выполняет тест клика по координатам cX, cY (координаты относительно текущего фрейма)
	// Если клик принимает другой элемент, либо элемент не родитель ожидаемого, то выбросим ошибк
	clickable, err := e.call(atom.IsClickableAt, cX, cY)
	if err != nil || !clickable.Bool() {
		return x, y, ErrElementOverlapped
	}
	return x, y, nil
}

// Click ...
func (e *element) Click() error {
	if _, err := e.call(atom.ScrollIntoView); err != nil {
		return err
	}
	if _, err := e.call(atom.AddEventFired, "click"); err != nil {
		return err
	}
	// calculate click point
	x, y, err := e.clickablePoint()
	if err != nil {
		return err
	}
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventMoved, "none")
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventPressed, "left")
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventReleased, "left")
	triggered, err := e.call(atom.IsEventFired)
	if (err == nil && triggered.Bool()) || (err == ErrCannotFindContext) {
		return nil
	}
	return ErrClickNotTriggered
}

// SwitchToFrame switch context to frame
func (e *element) SwitchToFrame() error {
	node, err := e.getNode()
	if err != nil {
		return err
	}
	if "IFRAME" != node.NodeName {
		return ErrInvalidElementFrame
	}
	return e.session.createIsolatedWorld(node.FrameID)
}

// IsVisible is element visible (element has area that clickable in viewport)
func (e *element) IsVisible() (bool, error) {
	if _, _, err := e.clickablePoint(); err != nil {
		if err == ErrElementInvisible || err == ErrElementOverlapped {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Hover hover mouse on element
func (e *element) Hover() error {
	if _, err := e.call(atom.ScrollIntoView); err != nil {
		return err
	}
	x, y, err := e.clickablePoint()
	if err != nil {
		return err
	}
	// add onmouseover event listener on element
	if _, err = e.call(atom.AddEventFired, "mouseover"); err != nil {
		return err
	}
	if err = e.session.MouseMove(x, y); err != nil {
		return err
	}
	// check to mouseover happens
	if triggered, err := e.call(atom.IsEventFired); err == nil && triggered.Bool() {
		return nil
	}
	return ErrHoverNotTriggered
}

// Clear ...
func (e *element) Clear() error {
	var err error
	if err = e.Focus(); err != nil {
		return err
	}
	_, err = e.call(atom.ClearInput)
	return err
}

// Type ...
func (e *element) Type(text string, key ...rune) error {
	var err error
	if enable, err := e.call(atom.IsFocusable); err != nil || !enable.Bool() {
		return ErrElementNotFocusable
	}
	if err = e.Clear(); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 200)
	if _, err := e.call(atom.DispatchEvents, []string{"keydown"}); err != nil {
		return err
	}
	// insert text, not typing
	err = e.session.insertText(text)
	if err != nil {
		return err
	}
	if _, err := e.call(atom.DispatchEvents, []string{"keypress", "input", "keyup", "change"}); err != nil {
		return err
	}
	// send keyboard key after some pause
	if key != nil {
		time.Sleep(time.Millisecond * 200)
		return e.session.SendKeys(key...)
	}
	return nil
}

func (e *element) string(functionDeclaration string, arg ...interface{}) (string, error) {
	res, err := e.call(functionDeclaration, arg...)
	if err != nil {
		return "", err
	}
	if res.Type != "string" {
		return "", ErrInvalidString
	}
	return res.Value.(string), nil
}

// GetText ...
func (e *element) GetText() (string, error) {
	return e.string(atom.GetInnerText)
}

// SetAttr ...
func (e *element) SetAttr(attr string, value string) error {
	_, err := e.call(atom.SetAttr, attr, value)
	return err
}

// GetAttr ...
func (e *element) GetAttr(attr string) (string, error) {
	return e.string(atom.GetAttr, attr)
}

// GetRectangle ...
func (e *element) GetRectangle() (*devtool.Rect, error) {
	q, err := e.session.getContentQuads(0, e.ID, false)
	if err != nil {
		return nil, err
	}
	rect := &devtool.Rect{
		X:      q[0].X,
		Y:      q[0].Y,
		Width:  q[1].X - q[0].X,
		Height: q[3].Y - q[0].Y,
	}
	return rect, nil
}

// GetComputedStyle ...
func (e *element) GetComputedStyle(style string) (string, error) {
	return e.string(atom.GetComputedStyle, style)
}

// GetSelected ...
func (e *element) GetSelected(selectedText bool) ([]string, error) {
	a := atom.GetSelected
	if selectedText {
		a = atom.GetSelectedText
	}
	ro, err := e.call(a)
	if err != nil {
		return nil, err
	}
	descriptor, err := e.session.getProperties(ro.ObjectID)
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

// Select ...
func (e *element) Select(values ...string) error {
	node, err := e.getNode()
	if err != nil {
		return err
	}
	if "SELECT" != node.NodeName {
		return ErrInvalidElementSelect
	}
	has, err := e.call(atom.SelectHasOptions, values)
	if !has.Bool() {
		return ErrInvalidElementOption
	}
	_, err = e.call(atom.Select, values)
	return err
}

// Checkbox Checkbox
func (e *element) Checkbox(check bool) error {
	if _, err := e.call(atom.CheckBox, check); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250) // todo
	if _, err := e.call(atom.DispatchEvents, []string{"click", "change"}); err != nil {
		return err
	}
	return nil
}

// IsChecked ...
func (e *element) IsChecked() (bool, error) {
	checked, err := e.call(atom.IsChecked)
	return checked.Bool(), err
}

// GetEventListeners returns event listeners of the given object.
func (e *element) GetEventListeners() ([]string, error) {
	msg, err := e.session.blockingSend("DOMDebugger.getEventListeners", Map{
		"objectId": e.ID,
		"depth":    1,
		"pierce":   true,
	})
	if err != nil {
		return nil, err
	}
	events := new(devtool.EventListeners)
	if err = msg.Unmarshal(events); err != nil {
		return nil, err
	}
	types := make([]string, len(events.Listeners))
	for n, e := range events.Listeners {
		types[n] = e.Type
	}
	return types, nil
}
