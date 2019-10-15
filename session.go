package witness

import (
	"container/list"
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/ecwid/witness/pkg/devtool"
)

// Session CDP session
type Session struct {
	client        *CDP
	id            string
	targetID      string
	contextID     int64
	frameID       string
	incomingEvent chan *rpcEvent
	callbacks     map[string]*list.List
	closed        chan bool
	context       context.Context
	document      *element
	mxContext     sync.Mutex
	mxCall        sync.Mutex
}

// TickerFunc ...
type TickerFunc func() (interface{}, error)

func (session *Session) panic(p interface{}) {
	session.client.Logging.Print(LevelFatal, p)
	panic(p)
}

// NewSession ...
func (c *CDP) newSession(targetID string) (*Session, error) {
	session := &Session{
		client:        c,
		incomingEvent: make(chan *rpcEvent, 1),
		callbacks:     make(map[string]*list.List),
		closed:        make(chan bool, 1),
		targetID:      targetID,
		frameID:       targetID,
	}
	session.document = newElement(session, nil, "", "document")
	go session.listener()
	return session, session.switchTarget()
}

func (session *Session) switchTarget() error {
	s, err := session.blockingSend("Target.attachToTarget", Map{"targetId": session.targetID, "flatten": true})
	if err != nil {
		return err
	}
	session.id = s.json().String("sessionId")
	session.client.addSession(session)
	enables := map[string]Map{
		"Page.enable":    nil,
		"Runtime.enable": nil,
		"Network.enable": Map{"maxPostDataSize": 1024},
	}
	for k, v := range enables {
		if _, err := session.blockingSend(k, v); err != nil {
			return err
		}
	}
	return session.createIsolatedWorld(session.targetID)
}

func (session *Session) lockContext(fn func()) {
	session.mxContext.Lock()
	defer session.mxContext.Unlock()
	fn()
}

func (session *Session) getContextID() int64 {
	session.mxContext.Lock()
	defer session.mxContext.Unlock()
	return session.contextID
}

func (session *Session) listener() {
	for e := range session.incomingEvent {

		session.mxCall.Lock()
		lst, has := session.callbacks[e.Method]
		if has {
			for p := lst.Front(); p != nil; p = p.Next() {
				go p.Value.(func([]byte))(e.Params)
			}
		}
		session.mxCall.Unlock()

		switch e.Method {
		case "Runtime.executionContextsCleared":
			session.lockContext(func() {
				session.contextID = 0
			})

		case "Runtime.executionContextCreated":
			ecc := new(devtool.ExecutionContextCreated)
			if err := e.Params.Unmarshal(ecc); err != nil {
				session.panic(err)
			}
			session.lockContext(func() {
				frameID := ecc.Context.AuxData["frameId"].(string)
				if session.frameID == frameID {
					session.contextID = ecc.Context.ID
				}
			})

		case "Runtime.executionContextDestroyed":
			ecd := new(devtool.ExecutionContextDestroyed)
			if err := e.Params.Unmarshal(ecd); err != nil {
				session.panic(err)
			}
			session.lockContext(func() {
				if session.contextID == ecd.ExecutionContextID {
					session.contextID = 0
				}
			})

		case "Target.targetCrashed":
			targetCrashed := new(devtool.TargetCrashed)
			if err := e.Params.Unmarshal(targetCrashed); err != nil {
				session.panic(err)
			}
			session.panic(targetCrashed)

		case "Target.targetDestroyed":
			targetDestroyed := new(devtool.TargetDestroyed)
			if err := e.Params.Unmarshal(targetDestroyed); err != nil {
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

func (session *Session) blockingSend(method string, send interface{}) (bytes, error) {
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

func (session *Session) createIsolatedWorld(frameID string) error {
	var err error
	msg, err := session.blockingSend("Page.createIsolatedWorld", Map{
		"frameId":             frameID,
		"name":                "__utilityWorld__",
		"grantUniveralAccess": true,
	})
	if err != nil {
		return err
	}
	session.lockContext(func() {
		session.frameID = frameID
		session.contextID = msg.json().Int("executionContextId")
	})
	return nil
}

// Ticker ...
func (session *Session) Ticker(call TickerFunc) (interface{}, error) {
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

func (session *Session) getNavigationHistory() (*devtool.NavigationHistory, error) {
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

// Subscribe subscribe to CDP event
func (session *Session) subscribe(method string, callback func(params []byte)) (unsubscribe func()) {
	session.mxCall.Lock()
	defer session.mxCall.Unlock()
	if _, has := session.callbacks[method]; !has {
		session.callbacks[method] = list.New()
	}
	p := session.callbacks[method].PushBack(callback)
	return func() {
		session.mxCall.Lock()
		defer session.mxCall.Unlock()
		session.callbacks[method].Remove(p)
	}
}

func (session *Session) getFrameOwner(frameID string) (int64, error) {
	msg, err := session.blockingSend("DOM.getFrameOwner", Map{"frameId": frameID})
	if err != nil {
		return 0, err
	}
	return msg.json().Int("backendNodeId"), nil
}

func (session *Session) getContentQuads(backendNodeID int64, objectID string, viewportCorrection bool) (devtool.Quad, error) {
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

func (session *Session) getLayoutMetrics() (*devtool.LayoutMetrics, error) {
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
