package control

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ecwid/control/cdp"
	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/overlay"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
)

// The Longest post body size (in bytes) that would be included in requestWillBeSent notification
var (
	MaxPostDataSize = 20 * 1024 // 20KB
)

const Blank = "about:blank"
const hitCheckFunc = `__control_clk_backend_hit`

var (
	ErrTargetDestroyed           error = errors.New("target destroyed")
	ErrTargetDetached            error = errors.New("session detached from target")
	ErrNetworkIdleReachedTimeout error = errors.New("session network idle reached timeout")
)

type TargetCrashedError []byte

func (t TargetCrashedError) Error() string {
	return string(t)
}

func mustUnmarshal[T any](u cdp.Message) T {
	var value T
	err := json.Unmarshal(u.Params, &value)
	if err != nil {
		panic(err)
	}
	return value
}

type Session struct {
	timeout          time.Duration
	context          context.Context
	transport        *cdp.Transport
	targetID         target.TargetID
	sessionID        string
	frames           *sync.Map
	Frame            *Frame
	highlightEnabled bool
	mouse            Mouse
	kb               Keyboard
	touch            Touch
}

func (s *Session) Transport() *cdp.Transport {
	return s.transport
}

func (s *Session) Context() context.Context {
	return s.context
}

func (s *Session) Log(msg string, args ...any) {
	level := slog.LevelInfo
	args = append(args, "sessionId", s.sessionID)
	for n := range args {
		switch a := args[n].(type) {
		case error:
			if a != nil {
				args[n] = a.Error()
				level = slog.LevelWarn
			}
		}
	}
	s.transport.Log(level, msg, args...)
}

func (s *Session) GetID() string {
	return s.sessionID
}

func (s *Session) IsDone() bool {
	select {
	case <-s.context.Done():
		return true
	default:
		return false
	}
}

func (s *Session) Call(method string, send, recv any) error {
	select {
	case <-s.context.Done():
		return context.Cause(s.context)
	default:
	}
	future := s.transport.Send(&cdp.Request{
		SessionID: string(s.sessionID),
		Method:    method,
		Params:    send,
	})
	defer future.Cancel()

	ctxTo, cancel := context.WithTimeout(s.context, s.timeout)
	defer cancel()
	value, err := future.Get(ctxTo)
	if err != nil {
		return err
	}

	if recv != nil {
		return json.Unmarshal(value.Result, recv)
	}
	return nil
}

func (s *Session) Subscribe(desc string) (channel chan cdp.Message, cancel func()) {
	return s.transport.Subscribe(s.sessionID, desc)
}

func NewSession(transport *cdp.Transport, targetID target.TargetID) (*Session, error) {
	var session = &Session{
		transport: transport,
		targetID:  targetID,
		timeout:   60 * time.Second,
		frames:    &sync.Map{},
	}
	session.mouse = NewMouse(session)
	session.kb = NewKeyboard(session)
	session.touch = NewTouch(session)
	session.Frame = &Frame{
		session: session,
		id:      common.FrameId(session.targetID),
	}
	var cancel func(error)
	session.context, cancel = context.WithCancelCause(transport.Context())
	val, err := target.AttachToTarget(session, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return nil, err
	}
	session.sessionID = string(val.SessionId)
	channel, unsubscribe := session.Subscribe("session-core-handler")
	go func() {
		if err := session.handle(channel); err != nil {
			unsubscribe()
			cancel(err)
		}
	}()
	if err = page.Enable(session); err != nil {
		return nil, err
	}
	if err = page.SetLifecycleEventsEnabled(session, page.SetLifecycleEventsEnabledArgs{Enabled: true}); err != nil {
		return nil, err
	}
	if err = runtime.Enable(session); err != nil {
		return nil, err
	}
	if err = dom.Enable(session, dom.EnableArgs{IncludeWhitespace: "none"}); err != nil {
		return nil, err
	}
	if err = target.SetDiscoverTargets(session, target.SetDiscoverTargetsArgs{Discover: true}); err != nil {
		return nil, err
	}
	if err = network.Enable(session, network.EnableArgs{MaxPostDataSize: MaxPostDataSize}); err != nil {
		return nil, err
	}
	if err = runtime.AddBinding(session, runtime.AddBindingArgs{Name: hitCheckFunc}); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *Session) EnableHighlight() error {
	if err := overlay.Enable(s); err != nil {
		return err
	}
	s.highlightEnabled = true
	return nil
}

func (s *Session) handle(channel chan cdp.Message) error {
	for message := range channel {
		switch message.Method {

		case "Runtime.executionContextCreated":
			executionContextCreated := mustUnmarshal[runtime.ExecutionContextCreated](message)
			aux := executionContextCreated.Context.AuxData.(map[string]any)
			frameID := aux["frameId"].(string)
			s.frames.Store(common.FrameId(frameID), executionContextCreated.Context.UniqueId)

		case "Page.frameDetached":
			frameDetached := mustUnmarshal[page.FrameDetached](message)
			s.frames.Delete(frameDetached.FrameId)

		case "Target.detachedFromTarget":
			detachedFromTarget := mustUnmarshal[target.DetachedFromTarget](message)
			if s.sessionID == string(detachedFromTarget.SessionId) {
				return ErrTargetDetached
			}

		case "Target.targetDestroyed":
			targetDestroyed := mustUnmarshal[target.TargetDestroyed](message)
			if s.targetID == targetDestroyed.TargetId {
				return ErrTargetDestroyed
			}

		case "Target.targetCrashed":
			targetCrashed := mustUnmarshal[target.TargetCrashed](message)
			if s.targetID == targetCrashed.TargetId {
				return TargetCrashedError(message.Params)
			}
		}
	}
	return nil
}

func (s *Session) funcCalled(fn string) cdp.Future[runtime.BindingCalled] {
	var channel, cancel = s.Subscribe(fmt.Sprintf("func-%s-called-listener", fn))
	callback := func(resolve func(runtime.BindingCalled), reject func(error)) {
		for value := range channel {
			if value.Method == "Runtime.bindingCalled" {
				var result runtime.BindingCalled
				if err := json.Unmarshal(value.Params, &result); err != nil {
					reject(err)
					return
				}
				if result.Name == fn {
					resolve(result)
					return
				}
			}
		}
	}
	return cdp.NewPromise(callback, cancel)
}

func (s *Session) CaptureScreenshot(format string, quality int, clip *page.Viewport, fromSurface, captureBeyondViewport, optimizeForSpeed bool) ([]byte, error) {
	val, err := page.CaptureScreenshot(s, page.CaptureScreenshotArgs{
		Format:                format,
		Quality:               quality,
		Clip:                  clip,
		FromSurface:           fromSurface,
		CaptureBeyondViewport: captureBeyondViewport,
		OptimizeForSpeed:      optimizeForSpeed,
	})
	if err != nil {
		return nil, err
	}
	return val.Data, nil
}

func (s *Session) SetDownloadBehavior(behavior string, downloadPath string, eventsEnabled bool) error {
	return browser.SetDownloadBehavior(s, browser.SetDownloadBehaviorArgs{
		Behavior:      behavior,
		DownloadPath:  downloadPath,
		EventsEnabled: eventsEnabled, // default false
	})
}

func (s *Session) MustSetDownloadBehavior(behavior string, downloadPath string, eventsEnabled bool) {
	if err := s.SetDownloadBehavior(behavior, downloadPath, eventsEnabled); err != nil {
		panic(err)
	}
}

func (s *Session) GetTargetCreated() cdp.Future[target.TargetCreated] {
	return Subscribe(s, "Target.targetCreated", func(t target.TargetCreated) bool {
		return t.TargetInfo.Type == "page" && t.TargetInfo.OpenerId == s.targetID
	})
}

func (s *Session) AttachToTarget(id target.TargetID) (*Session, error) {
	return NewSession(s.transport, id)
}

func (s *Session) CreatePageTargetTab(url string) (*Session, error) {
	if url == "" {
		url = Blank // headless chrome crash when url is empty
	}
	r, err := target.CreateTarget(s, target.CreateTargetArgs{Url: url})
	if err != nil {
		return nil, err
	}
	return s.AttachToTarget(r.TargetId)
}

func (s *Session) Activate() error {
	return target.ActivateTarget(s, target.ActivateTargetArgs{TargetId: s.targetID})
}

func (s *Session) Close() error {
	return s.CloseTarget(s.targetID)
}

func (s *Session) CloseTarget(id target.TargetID) (err error) {
	err = target.CloseTarget(s, target.CloseTargetArgs{TargetId: id})
	/* Target.detachedFromTarget event may come before the response of CloseTarget call */
	if err == ErrTargetDetached {
		return nil
	}
	return err
}

func (s *Session) Click(point Point) error {
	return s.mouse.Click(MouseLeft, point, time.Millisecond*85)
}

func (s *Session) MouseDown(point Point) error {
	return s.mouse.Down(MouseLeft, point)
}

func (s *Session) MustClick(point Point) {
	if err := s.Click(point); err != nil {
		panic(err)
	}
}

func (s *Session) Swipe(from, to Point) error {
	return s.touch.Swipe(from, to)
}

func (s *Session) MustSwipe(from, to Point) {
	if err := s.Swipe(from, to); err != nil {
		panic(err)
	}
}

func (s *Session) Hover(point Point) error {
	return s.mouse.Move(MouseNone, point)
}

func (s *Session) MustHover(point Point) {
	if err := s.Hover(point); err != nil {
		panic(err)
	}
}

func (s *Session) GetLayout() Optional[page.GetLayoutMetricsVal] {
	view, err := page.GetLayoutMetrics(s)
	if err != nil {
		return Optional[page.GetLayoutMetricsVal]{err: err}
	}
	return Optional[page.GetLayoutMetricsVal]{value: *view}
}

func (s *Session) GetNavigationEntry() Optional[page.NavigationEntry] {
	val, err := page.GetNavigationHistory(s)
	if err != nil {
		return Optional[page.NavigationEntry]{err: err}
	}
	if val.CurrentIndex == -1 {
		return Optional[page.NavigationEntry]{value: page.NavigationEntry{Url: Blank}}
	}
	return Optional[page.NavigationEntry]{value: *val.Entries[val.CurrentIndex]}
}

func (s *Session) GetCurrentURL() Optional[string] {
	return optional[string](s.getCurrentURL())
}

func (s *Session) getCurrentURL() (string, error) {
	e, err := s.GetNavigationEntry().Unwrap()
	if err != nil {
		return "", err
	}
	return e.Url, nil
}

func (s *Session) NavigateHistory(delta int) error {
	val, err := page.GetNavigationHistory(s)
	if err != nil {
		return err
	}
	move := val.CurrentIndex + delta
	if move >= 0 && move < len(val.Entries) {
		return page.NavigateToHistoryEntry(s, page.NavigateToHistoryEntryArgs{
			EntryId: val.Entries[move].Id,
		})
	}
	return nil
}
