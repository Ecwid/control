package control

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/ecwid/control/future"
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

// The Longest post body size (in bytes) that would be included in requestWillBeSent notification
var (
	MaxPostDataSize = 20 * 1024 // 20KB
)

const Blank = "about:blank"
const hitCheckFunc = `__control_clk_backend_hit`

var (
	ErrTargetDestroyed    error = errors.New("target destroyed")
	ErrTargetDetached     error = errors.New("session detached from target")
	ErrSubscriptionClosed error = errors.New("event subscription channel closed")
)

type TargetCrashedError []byte

func (t TargetCrashedError) Error() string {
	return string(t)
}

type Session struct {
	browser     *Browser
	caller      CdpCaller
	close       sync.Once
	unsubscribe func()
	targetID    target.TargetID
	framesMu    sync.RWMutex
	frames      map[common.FrameId]string
	Frame       *Frame
}

func (s *Session) GetBrowser() *Browser {
	return s.browser
}

func (s *Session) SetTimeout(timeout time.Duration) {
	s.caller.timeout = timeout
}

func (s *Session) GetCaller() CdpCaller {
	return s.caller
}

func (s *Session) Call(method string, send, recv any) error {
	return s.caller.Call(method, send, recv)
}

func (s *Session) Context() context.Context {
	return s.GetCaller().Context()
}

func (s *Session) Timeout() time.Duration {
	return s.GetCaller().timeout
}

func GetWithTimeout[T any](s *Session, future future.Future[T]) (T, error) {
	ctxTo, cancel := context.WithTimeout(s.Context(), s.Timeout())
	defer cancel()
	return future.Get(ctxTo)
}

func (s *Session) Log(msg string, args ...any) {
	level := slog.LevelInfo
	args = append(args, "sessionId", s.GetCaller().sessionID)
	for n := range args {
		switch a := args[n].(type) {
		case error:
			if a != nil {
				args[n] = a.Error()
				level = slog.LevelWarn
			}
		}
	}
	s.GetCaller().transport.Log(level, msg, args...)
}

func (s *Session) GetID() string {
	return s.GetCaller().sessionID
}

func (s *Session) IsDone() bool {
	return s.GetCaller().IsDone()
}

func (s *Session) Subscribe() (channel <-chan transport.Message, cancel func()) {
	return s.GetCaller().Subscribe()
}

func (s *Session) fatal(err error) {
	s.close.Do(func() {
		if err != nil {
			s.Log("session closed", "targetId", s.targetID, "error", err)
		}
		if s.unsubscribe != nil {
			s.unsubscribe()
		}
		s.caller.Cancel(err)
	})
}

func (s *Session) setFrameExecutionContextID(id common.FrameId, executionContextID string) {
	s.framesMu.Lock()
	s.frames[id] = executionContextID
	s.framesMu.Unlock()
}

func (s *Session) deleteFrameExecutionContextID(id common.FrameId) {
	s.framesMu.Lock()
	delete(s.frames, id)
	s.framesMu.Unlock()
}

func (s *Session) getFrameExecutionContextID(id common.FrameId) string {
	s.framesMu.RLock()
	defer s.framesMu.RUnlock()
	return s.frames[id]
}

func (b *Browser) NewSession(targetID target.TargetID) (*Session, error) {
	sessionCtx, sessionCancel := context.WithCancelCause(b.caller.ctx)
	cdpCaller := CdpCaller{
		ctx:       sessionCtx,
		cancel:    sessionCancel,
		timeout:   b.caller.timeout,
		transport: b.caller.transport,
	}
	var session = &Session{
		browser:  b,
		caller:   cdpCaller,
		targetID: targetID,
		frames:   make(map[common.FrameId]string),
	}
	session.Frame = &Frame{session: session, id: common.FrameId(session.targetID)}

	sessionID, err := b.attachToTarget(targetID)
	if err != nil {
		session.fatal(err)
		return nil, err
	}
	session.caller.sessionID = string(sessionID)
	session.startHandleLoop()
	if err = session.enableDefaults(); err != nil {
		// non necessary to detach from target, because session will be closed on error
		session.fatal(err)
		return nil, err
	}
	return session, nil
}

func (s *Session) enableDefaults() error {
	if err := page.Enable(s, page.EnableArgs{EnableFileChooserOpenedEvent: true}); err != nil {
		return err
	}
	if err := page.SetLifecycleEventsEnabled(s, page.SetLifecycleEventsEnabledArgs{Enabled: true}); err != nil {
		return err
	}
	if err := runtime.Enable(s); err != nil {
		return err
	}
	if err := dom.Enable(s, dom.EnableArgs{IncludeWhitespace: "none"}); err != nil {
		return err
	}
	if err := target.SetDiscoverTargets(s, target.SetDiscoverTargetsArgs{Discover: true}); err != nil {
		return err
	}
	if err := network.Enable(s, network.EnableArgs{MaxPostDataSize: MaxPostDataSize}); err != nil {
		return err
	}
	if err := runtime.AddBinding(s, runtime.AddBindingArgs{Name: hitCheckFunc}); err != nil {
		return err
	}
	return nil
}

func (s *Session) startHandleLoop() {
	channel, unsubscribe := s.Subscribe()
	s.unsubscribe = unsubscribe
	go func() {
		if err := s.handle(channel); err != nil {
			s.fatal(err)
		}
	}()
}

func handleExecutionContextCreated(s *Session, message transport.Message) error {
	executionContextCreated, err := transport.Unmarshal[runtime.ExecutionContextCreated](message)
	if err != nil {
		return err
	}
	aux, ok := executionContextCreated.Context.AuxData.(map[string]any)
	if !ok {
		return fmt.Errorf("runtime.executionContextCreated: invalid auxData type %T", executionContextCreated.Context.AuxData)
	}

	frameIDRaw, ok := aux["frameId"]
	if !ok {
		return errors.New("runtime.executionContextCreated: frameId is missing")
	}

	frameID, ok := frameIDRaw.(string)
	if !ok {
		return fmt.Errorf("runtime.executionContextCreated: invalid frameId type %T", frameIDRaw)
	}
	if frameID == "" {
		return errors.New("runtime.executionContextCreated: frameId is empty")
	}

	s.setFrameExecutionContextID(common.FrameId(frameID), executionContextCreated.Context.UniqueId)
	return nil
}

func handleFrameDetached(s *Session, message transport.Message) error {
	frameDetached, err := transport.Unmarshal[page.FrameDetached](message)
	if err != nil {
		return err
	}
	s.deleteFrameExecutionContextID(frameDetached.FrameId)
	return nil
}

func handleDetachedFromTarget(s *Session, message transport.Message) error {
	detachedFromTarget, err := transport.Unmarshal[target.DetachedFromTarget](message)
	if err != nil {
		return err
	}
	if s.caller.sessionID == string(detachedFromTarget.SessionId) {
		return ErrTargetDetached
	}
	return nil
}

func handleTargetDestroyed(s *Session, message transport.Message) error {
	targetDestroyed, err := transport.Unmarshal[target.TargetDestroyed](message)
	if err != nil {
		return err
	}
	if s.targetID == targetDestroyed.TargetId {
		return ErrTargetDestroyed
	}
	return nil
}

func handleTargetCrashed(s *Session, message transport.Message) error {
	targetCrashed, err := transport.Unmarshal[target.TargetCrashed](message)
	if err != nil {
		return err
	}
	if s.targetID == targetCrashed.TargetId {
		return TargetCrashedError(message.Params)
	}
	return nil
}

func (s *Session) handle(channel <-chan transport.Message) error {
	for message := range channel {
		switch message.Method {
		case "Runtime.executionContextCreated":
			if err := handleExecutionContextCreated(s, message); err != nil {
				return err
			}
		case "Page.frameDetached":
			if err := handleFrameDetached(s, message); err != nil {
				return err
			}
		case "Target.detachedFromTarget":
			if err := handleDetachedFromTarget(s, message); err != nil {
				return err
			}
		case "Target.targetDestroyed":
			if err := handleTargetDestroyed(s, message); err != nil {
				return err
			}
		case "Target.targetCrashed":
			if err := handleTargetCrashed(s, message); err != nil {
				return err
			}
		}
	}
	return ErrSubscriptionClosed
}

func subscribeMessage[T any](s *Session, finder func(transport.Message) (T, bool, error)) future.Future[T] {
	return future.Execute(func(resolve func(T), reject func(error), canceled <-chan struct{}) {
		channel, unsubscribe := s.Subscribe()
		defer unsubscribe()
		for {
			select {
			case <-canceled:
				return

			case <-s.Context().Done():
				reject(context.Cause(s.Context()))
				return

			case value, ok := <-channel:
				if !ok {
					reject(ErrSubscriptionClosed)
					return
				}
				result, ok, err := finder(value)
				if err != nil {
					reject(err)
					return
				}
				if ok {
					resolve(result)
					return
				}
			}
		}
	})
}

func subscribeMethod[T any](s *Session, method string, match func(T) (bool, error)) future.Future[T] {
	return subscribeMessage(s, func(value transport.Message) (T, bool, error) {
		var zero T
		if value.Method != method {
			return zero, false, nil
		}
		result, err := transport.Unmarshal[T](value)
		if err != nil {
			return zero, false, err
		}
		isMatched, err := match(result)
		if err != nil {
			return zero, false, err
		}
		if !isMatched {
			return zero, false, nil
		}

		return result, true, nil
	})
}
