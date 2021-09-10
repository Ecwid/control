package control

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
	"github.com/ecwid/control/transport/observe"
)

type Map map[string]interface{}

const (
	blankPage = "about:blank"
)

type Session struct {
	transport    transport.T
	id           target.SessionID
	eventPool    chan observe.Value
	manager      *Manager
	exited       chan struct{}
	exitCode     error
	observable   *observe.Observable
	obsuid       uint64 // observers incremental id
	Timeout      time.Duration
	PoolingEvery time.Duration
}

func (s Session) Call(method string, send, recv interface{}) error {
	select {
	case <-s.exited:
		return s.exitCode
	default:
		return s.transport.Call(string(s.id), method, send, recv)
	}
}

func New(t transport.T) *Session {
	var headlessSession = &Session{
		obsuid:       0,
		id:           "",
		transport:    t,
		eventPool:    make(chan observe.Value, 999),
		observable:   observe.New(),
		exited:       make(chan struct{}, 1),
		Timeout:      time.Second * 15,
		PoolingEvery: time.Millisecond * 500,
	}
	return headlessSession
}
func (s Session) GetTransport() transport.T {
	return s.transport
}

func (s *Session) CreateTarget(url string) error {
	if url == "" {
		url = blankPage // headless chrome crash when url is empty
	}
	r, err := target.CreateTarget(s, target.CreateTargetArgs{Url: url})
	if err != nil {
		return err
	}
	return s.AttachToTarget(r.TargetId)
}

func (s *Session) detachFromTarget() error {
	if s.ID() != "" {
		return target.DetachFromTarget(s, target.DetachFromTargetArgs{SessionId: s.id})
	}
	return nil
}

func (s *Session) AttachToTarget(targetID target.TargetID) error {
	val, err := target.AttachToTarget(s, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return err
	}
	if err = s.detachFromTarget(); err != nil {
		return err
	}

	// run session lifecycle
	// s.targetID = targetID
	s.id = val.SessionId
	s.manager = newManager(s, targetID)
	go s.lifecycle()
	s.transport.Add(s)

	// default settings
	if err = page.Enable(s); err != nil {
		return err
	}
	if err = runtime.Enable(s); err != nil {
		return err
	}
	if err = page.SetLifecycleEventsEnabled(s, page.SetLifecycleEventsEnabledArgs{Enabled: true}); err != nil {
		return err
	}
	if err = target.SetDiscoverTargets(s, target.SetDiscoverTargetsArgs{Discover: true}); err != nil {
		return err
	}
	return nil
}

func (s Session) ID() string {
	return string(s.id)
}

func (s Session) Event() string {
	return s.ID()
}

func (s Session) Target() *Frame {
	return s.manager.head
}

func (s Session) TargetID() target.TargetID {
	return target.TargetID(s.Target().id)
}

func (s Session) Activate() error {
	return target.ActivateTarget(s, target.ActivateTargetArgs{
		TargetId: s.TargetID(),
	})
}

func (s Session) abort(err error) {
	s.exitCode = err
	close(s.exited)
}

func (s Session) Notify(val observe.Value) {
	s.eventPool <- val
}

func (s *Session) lifecycle() {
	defer func() {
		s.transport.Remove(s)
	}()
	for e := range s.eventPool {
		switch e.Method {

		case "Page.frameAttached":
			var v = new(page.FrameAttached)
			if err := json.Unmarshal(e.Params, v); err != nil {
				s.abort(err)
				return
			}
			s.manager.add(v.ParentFrameId, v.FrameId)

		case "Page.frameDetached":
			var v = new(page.FrameDetached)
			if err := json.Unmarshal(e.Params, v); err != nil {
				s.abort(err)
				return
			}
			s.manager.delete(v.FrameId)

		case "Runtime.executionContextCreated":
			var v = new(runtime.ExecutionContextCreated)
			if err := json.Unmarshal(e.Params, v); err != nil {
				s.abort(err)
				return
			}
			frameID := common.FrameId((v.Context.AuxData.(map[string]interface{}))["frameId"].(string))
			s.manager.edit(frameID, func(f *Frame) {
				f.contextID = v.Context.Id
			})

		case "Target.targetCrashed":
			s.abort(errors.New(string(e.Params)))
			return

		case "Target.targetDestroyed":
			var v = new(target.TargetDestroyed)
			if err := json.Unmarshal(e.Params, v); err != nil {
				s.abort(err)
				return
			}
			if v.TargetId == s.TargetID() {
				s.abort(ErrSessionClosed)
				return
			}

		case "Target.detachedFromTarget":
			var v = new(target.DetachedFromTarget)
			if err := json.Unmarshal(e.Params, v); err != nil {
				s.abort(err)
				return
			}
			if v.SessionId == s.id {
				return
			}

		}
		s.observable.Notify(e.Method, e)
	}
}

func (s *Session) Subscribe(event string, async bool, v func(e observe.Value)) (unsubscribe func()) {
	var (
		uid = atomic.AddUint64(&s.obsuid, 1)
		val = observe.NewSimpleObserver(fmt.Sprintf("%d", uid), event, v)
	)
	if async {
		s.observable.Add(observe.AsyncSimpleObserver(val))
	} else {
		s.observable.Add(val)
	}
	return func() {
		s.observable.Remove(val)
	}
}

func (s *Session) Mouse() Mouse {
	return Mouse{s: s}
}

func (s *Session) Network() Network {
	return Network{s: s}
}

func (s *Session) Emulation() Emulation {
	return Emulation{s: s}
}

// CaptureScreenshot get screen of current page
func (s Session) CaptureScreenshot(format string, quality int, clip *page.Viewport, fromSurface, captureBeyondViewport bool) ([]byte, error) {
	val, err := page.CaptureScreenshot(s, page.CaptureScreenshotArgs{
		Format:                format,
		Quality:               quality,
		Clip:                  clip,
		FromSurface:           fromSurface,
		CaptureBeyondViewport: captureBeyondViewport,
	})
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(string(val.Data))
}

// AddScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addScriptToEvaluateOnNewDocument
func (s Session) AddScriptToEvaluateOnNewDocument(source string) (page.ScriptIdentifier, error) {
	val, err := page.AddScriptToEvaluateOnNewDocument(s, page.AddScriptToEvaluateOnNewDocumentArgs{
		Source: source,
	})
	if err != nil {
		return "", err
	}
	return val.Identifier, nil
}

// RemoveScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-removeScriptToEvaluateOnNewDocument
func (s Session) RemoveScriptToEvaluateOnNewDocument(identifier page.ScriptIdentifier) error {
	return page.RemoveScriptToEvaluateOnNewDocument(s, page.RemoveScriptToEvaluateOnNewDocumentArgs{
		Identifier: identifier,
	})
}

// SetDownloadBehavior https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setDownloadBehavior
func (s Session) SetDownloadBehavior(behavior string, downloadPath string, eventsEnabled bool) error {
	return browser.SetDownloadBehavior(s, browser.SetDownloadBehaviorArgs{
		Behavior:      behavior,
		DownloadPath:  downloadPath,
		EventsEnabled: eventsEnabled, // default false
	})
}

// HandleJavaScriptDialog ...
func (s Session) HandleJavaScriptDialog(accept bool, promptText string) error {
	return page.HandleJavaScriptDialog(s, page.HandleJavaScriptDialogArgs{
		Accept:     accept,
		PromptText: promptText,
	})
}

func (s Session) Close() error {
	err := target.CloseTarget(s, target.CloseTargetArgs{
		TargetId: s.TargetID(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s Session) IsClosed() bool {
	select {
	case <-s.exited:
		return true
	default:
		return false
	}
}

func (s Session) ExitCode() error {
	return s.exitCode
}

func (s *Session) NewTargetCreatedCondition(createdTargetID *target.TargetID) Condition {
	return NewCondition(s, s.Timeout, func(value observe.Value) (bool, error) {
		if value.Method == "Target.targetCreated" {
			var v = new(target.TargetCreated)
			if err := json.Unmarshal(value.Params, v); err != nil {
				return false, err
			}
			if v.TargetInfo.Type == "page" && v.TargetInfo.OpenerId == s.TargetID() {
				*createdTargetID = v.TargetInfo.TargetId
				return true, nil
			}
		}
		return false, nil
	})
}
