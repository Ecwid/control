package witness

import (
	"time"

	"github.com/ecwid/witness/internal/atom"
	"github.com/ecwid/witness/pkg/devtool"
)

func (e *element) Release() error {
	return e.session.releaseObject(e.ID)
}

// Element ...
type element struct {
	session     *CDPSession
	ID          string
	description string
	parent      *element
	childs      []*element
	context     int64
}

func newElement(s *CDPSession, parent *element, ID string, description string) *element {
	c, _ := s.getContextID()
	e := &element{
		ID:          ID,
		session:     s,
		description: description,
		parent:      parent,
		childs:      make([]*element, 0),
		context:     c,
	}
	if parent != nil {
		parent.childs = append(parent.childs, e)
	}
	return e
}

func (e *element) isDetached() bool {
	c, err := e.session.getContextID()
	if err != nil {
		return true
	}
	return e.context != c
}

func (e *element) renew() error {
	if !e.isDetached() {
		return nil
	}
	sessContext, err := e.session.getContextID() // session's context actual at this moment
	if err != nil {
		return err
	}
	// this ErrStaleElementReference but we can renew element
	// request new element in parent's context by description
	new, err := e.session.query(e.parent, e.description)
	if err != nil {
		return err
	}
	e.ID = new.ObjectID
	e.description = new.Description
	e.context = sessContext
	return nil
}

func (e *element) Query(selector string) (Element, error) {
	element, err := e.session.query(e, selector)
	if err != nil {
		return nil, err
	}
	return newElement(e.session, e, element.ObjectID, element.Description), nil
}

func (e *element) QueryAll(selector string) []Element {
	v, err := e.session.queryAll(e, selector)
	if err != nil {
		return []Element{}
	}
	return v
}

// Expect searching selector (visible) with implicity wait timeout
func (e *element) C(selector string, visible bool) Element {
	el, err := e.session.Ticker(func() (interface{}, error) {
		new, err := e.Query(selector)
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

// Call evaluate javascript for element, for example `function() {return this.innerHTML}`
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
	return x, y, nil
}

// Click ...
func (e *element) Click() error {
	if _, err := e.call(atom.ScrollIntoView); err != nil {
		return err
	}
	x, y, err := e.clickablePoint()
	if err != nil {
		return err
	}
	if _, err := e.call(atom.PreventMissClick); err != nil {
		return err
	}
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventMoved, "none")
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventPressed, "left")
	e.session.dispatchMouseEvent(x, y, dispatchMouseEventReleased, "left")
	hit, err := e.call(atom.IsClickHit)
	// in case when click is initiate navigation which destroyed context of element (ErrNoSuchContext)
	// or click may closes a popup (ErrSessionClosed)
	if (err == nil && hit.Bool()) || err == ErrNoSuchContext || err == ErrSessionClosed {
		return nil
	}
	return ErrElementMissClick
}

// GetFrameID get if for IFRAME element
func (e *element) GetFrameID() (string, error) {
	node, err := e.getNode()
	if err != nil {
		return "", err
	}
	if "IFRAME" != node.NodeName {
		return "", ErrInvalidElementFrame
	}
	return node.FrameID, nil
}

// IsVisible is element visible (element has area that clickable in viewport)
func (e *element) IsVisible() (bool, error) {
	if _, _, err := e.clickablePoint(); err != nil {
		if err == ErrElementInvisible {
			return false, nil
		}
		return false, err
	}
	if vis, err := e.call(atom.IsVisible); err != nil || !vis.Bool() {
		return false, nil
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
	if err = e.session.MouseMove(x, y); err != nil {
		return err
	}
	return nil
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
	if enable, err := e.call(atom.IsVisible); err != nil || !enable.Bool() {
		return err
	}
	if err = e.Clear(); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250)
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
		time.Sleep(time.Millisecond * 250)
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

// ObserveMutation create MutationObserver Promise for element, returns type of first mutation
func (e *element) ObserveMutation(attributes, childList, subtree bool) (chan string, chan error) {
	chanerr := make(chan error, 1)
	mutation := make(chan string, 1)
	go func() {
		val, err := e.call(atom.MutationObserver, attributes, childList, subtree)
		if err != nil {
			chanerr <- err
			return
		}
		if val.Type != "string" {
			chanerr <- ErrInvalidString
		}
		mutation <- val.Value.(string)
	}()
	return mutation, chanerr
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
	if _, err := e.call(atom.ScrollIntoView); err != nil {
		return err
	}
	if _, err = e.call(atom.Select, values); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250)
	if _, err := e.call(atom.DispatchEvents, []string{"input", "change"}); err != nil {
		return err
	}
	return nil
}

// Checkbox Checkbox
func (e *element) Checkbox(check bool) error {
	if _, err := e.call(atom.CheckBox, check); err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 250)
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
