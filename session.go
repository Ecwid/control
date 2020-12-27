package witness

import (
	"container/list"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/ecwid/witness/internal/atom"
	"github.com/ecwid/witness/pkg/devtool"
)

// CDPSession CDP session
type CDPSession struct {
	client    *CDP
	rw        sync.Mutex
	contexts  sync.Map
	id        string
	targetID  string
	frameID   string
	incoming  chan *Event
	callbacks map[string]*list.List
	closed    chan bool
}

// TickerFunc ...
type TickerFunc func() (interface{}, error)

func (session *CDPSession) panic(p interface{}) {
	session.client.Logging.Print(LevelFatal, p)
	panic(p)
}

// NewSession ...
func (c *CDP) newSession(targetID string) (*Session, error) {
	session := &CDPSession{
		client:    c,
		incoming:  make(chan *Event, 1),
		callbacks: make(map[string]*list.List),
		closed:    make(chan bool, 1),
		targetID:  targetID,
		frameID:   targetID,
	}
	go session.listener()
	sess := &Session{
		Network:   session,
		Input:     session,
		Runtime:   session,
		Page:      session,
		Message:   session,
		Tabs:      session,
		Emulation: session,
	}
	return sess, session.switchTarget()
}

func (session *CDPSession) switchTarget() error {
	s, err := session.blockingSend("Target.attachToTarget", Map{"targetId": session.targetID, "flatten": true})
	if err != nil {
		return err
	}
	session.id = s.json().String("sessionId")
	session.client.addSession(session)
	enables := map[string]Map{
		"Page.enable":    nil,
		"Runtime.enable": nil,
		"Network.enable": Map{"maxPostDataSize": 1024}, // maxPostDataSize - Longest post body size (in bytes) that would be included in requestWillBeSent notification
	}
	for k, v := range enables {
		if _, err := session.blockingSend(k, v); err != nil {
			return err
		}
	}
	// context is may not be created yet
	_ = session.setFrame(session.targetID)
	return nil
}

func (session *CDPSession) setFrame(frameID string) error {
	v, ok := session.contexts.Load(frameID)
	if !ok {
		return ErrNoSuchFrame
	}
	session.client.Logging.Printf(LevelInfo, "session switch -> %s-%d", frameID, v.(int64))
	session.frameID = frameID
	return nil
}

func (session *CDPSession) getContextID() (int64, error) {
	id := session.frameID
	if v, ok := session.contexts.Load(id); ok {
		return v.(int64), nil
	}
	// todo remove
	// for main frame we can use default context with ID = 0
	if id == session.targetID {
		session.client.Logging.Printf(LevelInfo, "context for '%s' not found, trying to use default 0 context as it is main frame", id)
		return 0, nil
	}
	session.client.Logging.Printf(LevelInfo, "context for '%s' not found, try again later", id)
	return -1, ErrFrameDetached
}

func (session *CDPSession) listener() {
	for e := range session.incoming {

		session.rw.Lock()
		if list, has := session.callbacks[e.Method]; has {
			for p := list.Front(); p != nil; p = p.Next() {
				p.Value.(func(*Event))(e)
			}
		}
		session.rw.Unlock()

		switch e.Method {
		case "Runtime.executionContextsCleared":
			session.contexts.Range(func(k interface{}, v interface{}) bool {
				session.contexts.Delete(k)
				return true
			})

		case "Runtime.executionContextCreated":
			c := new(devtool.ExecutionContextCreated)
			if err := json.Unmarshal(e.Params, c); err != nil {
				session.panic(err)
			}
			session.contexts.Store(c.Context.AuxData["frameId"].(string), c.Context.ID)

		case "Runtime.executionContextDestroyed":
			c := new(devtool.ExecutionContextDestroyed)
			if err := json.Unmarshal(e.Params, c); err != nil {
				session.panic(err)
			}
			session.contexts.Range(func(k interface{}, v interface{}) bool {
				if v.(int64) == c.ExecutionContextID {
					session.contexts.Delete(k)
					return false
				}
				return true
			})

		case "Target.targetCrashed":
			targetCrashed := new(devtool.TargetCrashed)
			if err := json.Unmarshal(e.Params, targetCrashed); err != nil {
				session.panic(err)
			}
			session.panic(string(e.Params))

		case "Target.targetDestroyed":
			targetDestroyed := new(devtool.TargetDestroyed)
			if err := json.Unmarshal(e.Params, targetDestroyed); err != nil {
				session.panic(err)
			}
			if targetDestroyed.TargetID == session.targetID {
				close(session.closed)
				session.client.deleteSession(session.id)
				return
			}
		}
	}
}

// BlockingSend send message over CDP protocol
func (session *CDPSession) BlockingSend(method string, send interface{}) ([]byte, error) {
	return session.blockingSend(method, send)
}

func (session *CDPSession) blockingSend(method string, send interface{}) (bytes, error) {
	recv := session.client.sendOverProtocol(session.id, method, send)
	select {
	case message := <-recv:
		if message.isError() {
			return nil, message.Error.known()
		}
		return message.Result, nil
	case <-session.closed:
		return nil, ErrSessionClosed
	case <-time.After(session.client.Timeouts.WSTimeout):
		return nil, fmt.Errorf("websocket response reached timeout %s for %s -> %+v", session.client.Timeouts.WSTimeout.String(), method, send)
	}
}

// Ticker ...
func (session *CDPSession) Ticker(call TickerFunc) (interface{}, error) {
	var err error
	var v interface{}
	// first time without ticker
	if v, err = call(); err == nil {
		return v, nil
	}
	tick := time.NewTicker(session.client.Timeouts.Poll)
	implicitly := time.NewTimer(session.client.Timeouts.Implicitly)
	defer tick.Stop()
	defer implicitly.Stop()
	for {
		select {
		case <-implicitly.C:
			return nil, err
		case <-tick.C:
			if v, err = call(); err == nil {
				return v, nil
			}
		}
	}
}

func (session *CDPSession) getNavigationHistory() (*devtool.NavigationHistory, error) {
	msg, err := session.blockingSend("Page.getNavigationHistory", Map{})
	if err != nil {
		return nil, err
	}
	history := new(devtool.NavigationHistory)
	if err = msg.Unmarshal(history); err != nil {
		return nil, err
	}
	return history, nil
}

func (session *CDPSession) navigateToHistoryEntry(entryID int64) error {
	_, err := session.blockingSend("Page.navigateToHistoryEntry", Map{
		"entryId": entryID,
	})
	return err
}

// Subscribe subscribe to CDP event
func (session *CDPSession) subscribe(method string, callback func(event *Event)) (unsubscribe func()) {
	session.rw.Lock()
	defer session.rw.Unlock()
	if _, has := session.callbacks[method]; !has {
		session.callbacks[method] = list.New()
	}
	p := session.callbacks[method].PushBack(callback)
	return func() {
		session.rw.Lock()
		defer session.rw.Unlock()
		session.callbacks[method].Remove(p)
	}
}

func (session *CDPSession) getFrameOwner(frameID string) (int64, error) {
	msg, err := session.blockingSend("DOM.getFrameOwner", Map{"frameId": frameID})
	if err != nil {
		return 0, err
	}
	return msg.json().Int("backendNodeId"), nil
}

func (session *CDPSession) getContentQuads(backendNodeID int64, objectID string, viewportCorrection bool) (devtool.Quad, error) {
	p := Map{
		"backendNodeId": backendNodeID,
		"objectId":      objectID,
	}
	p.omitempty()
	msg, err := session.blockingSend("DOM.getContentQuads", p)
	if err != nil {
		return nil, err
	}
	cq := new(devtool.ContentQuads)
	if err = msg.Unmarshal(cq); err != nil {
		return nil, err
	}
	calc := cq.Calc()
	// should be at least one
	if len(calc) == 0 {
		return nil, ErrElementInvisible
	}
	metric, err := session.getLayoutMetrics()
	if err != nil {
		return nil, err
	}
	for _, quad := range calc {
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
				point.X = math.Min(math.Max(point.X, 0), float64(metric.LayoutViewport.ClientWidth))
				point.Y = math.Min(math.Max(point.Y, 0), float64(metric.LayoutViewport.ClientHeight))
			}
		}
		if quad.Area() > 1 {
			return quad, nil
		}
	}
	return nil, ErrElementIsOutOfViewport
}

// GetLayoutMetrics ...
func (session *CDPSession) GetLayoutMetrics() (*devtool.LayoutMetrics, error) {
	return session.getLayoutMetrics()
}

func (session *CDPSession) getLayoutMetrics() (*devtool.LayoutMetrics, error) {
	msg, err := session.blockingSend("Page.getLayoutMetrics", Map{})
	if err != nil {
		return nil, err
	}
	l := new(devtool.LayoutMetrics)
	if err = msg.Unmarshal(l); err != nil {
		return nil, err
	}
	return l, nil
}

func (session *CDPSession) getFrameTree() (*devtool.FrameTree, error) {
	msg, err := session.blockingSend("Page.getFrameTree", Map{})
	if err != nil {
		return nil, err
	}
	tree := new(devtool.FrameTreeResult)
	if err = msg.Unmarshal(tree); err != nil {
		return nil, err
	}
	return tree.FrameTree, nil
}

func (session *CDPSession) queryAll(parent *element, selector string) ([]Element, error) {
	selector = strings.ReplaceAll(selector, `"`, `\"`)
	var array *devtool.RemoteObject
	var err error
	if parent == nil {
		c, cerr := session.getContextID()
		if cerr != nil {
			return nil, cerr
		}
		array, err = session.evaluate(`document.querySelectorAll("`+selector+`")`, c, false, false)
	} else {
		array, err = parent.call(atom.QueryAll, selector)
	}
	if err != nil {
		return nil, err
	}
	if array == nil || array.Description == "NodeList(0)" {
		session.releaseObject(array.ObjectID)
		return nil, ErrNoSuchElement
	}
	els := make([]Element, 0)
	descriptor, err := session.getProperties(array.ObjectID)
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		els = append(els, newElement(session, parent, d.Value.ObjectID, d.Value.Description))
	}
	return els, nil
}

func (session *CDPSession) query(parent *element, selector string) (*devtool.RemoteObject, error) {
	selector = strings.ReplaceAll(selector, `"`, `\"`)
	var (
		queried   *devtool.RemoteObject
		contextID int64
		err       error
	)
	if parent == nil {
		if contextID, err = session.getContextID(); err != nil {
			return nil, err
		}
		queried, err = session.evaluate(`document.querySelector("`+selector+`")`, contextID, false, false)
	} else {
		queried, err = parent.call(atom.Query, selector)
	}
	if err != nil {
		return nil, err
	}
	if queried.ObjectID == "" {
		return nil, ErrNoSuchElement
	}
	return queried, nil
}
