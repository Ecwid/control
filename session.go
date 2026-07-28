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
	ErrSubscriptionClosed error = errors.New("event subscription closed")
)

type TargetCrashedError []byte

func (t TargetCrashedError) Error() string {
	return string(t)
}

type Session struct {
	// timeout defines default per-operation timeout for protocol calls.
	timeout time.Duration
	// context is canceled when the session closes or fails.
	context context.Context
	// cancel terminates the session context with a cause.
	cancel context.CancelCauseFunc
	// teardown guarantees fail/close sequence runs exactly once.
	teardown sync.Once
	// unsubscribeHandle detaches the main session event subscription.
	unsubscribeHandle func()
	// transport is the shared CDP transport used by the session.
	transport *transport.Transport
	// targetID is the browser target this session is attached to.
	targetID target.TargetID
	// sessionID is the CDP session identifier returned by attachToTarget.
	sessionID string
	// framesMu protects frames map concurrent access.
	framesMu sync.RWMutex
	// frames maps frame ID to runtime execution context unique ID.
	frames map[common.FrameId]string
	// Frame is the root frame wrapper for this session target.
	Frame *Frame
	// mouse provides mouse input actions scoped to this session.
	mouse Mouse
	// kb provides keyboard input actions scoped to this session.
	kb Keyboard
	// touch provides touch input actions scoped to this session.
	touch Touch
}

func newSession(tp *transport.Transport, targetID target.TargetID, timeout time.Duration) *Session {
	var session = &Session{
		transport: tp,
		targetID:  targetID,
		timeout:   timeout,
		frames:    make(map[common.FrameId]string),
	}
	session.mouse = NewMouse(session)
	session.kb = NewKeyboard(session)
	session.touch = NewTouch(session)
	session.Frame = &Frame{session: session, id: common.FrameId(session.targetID)}
	return session
}

func (s *Session) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *Session) Transport() *transport.Transport {
	return s.transport
}

func (s *Session) Context() context.Context {
	return s.context
}

func (s *Session) startContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(s.context, s.timeout)
}

func (s *Session) withTimeout(run func(context.Context) error) error {
	ctxTo, cancel := s.startContext()
	defer cancel()
	return run(ctxTo)
}

func GetWithTimeout[T any](s *Session, future future.Future[T]) (T, error) {
	ctxTo, cancel := s.startContext()
	defer cancel()
	return future.Get(ctxTo)
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
	return s.withTimeout(func(ctx context.Context) error {
		return s.transport.Call(ctx, s.sessionID, method, send, recv)
	})
}

func (s *Session) Subscribe() (channel <-chan transport.Message, cancel func()) {
	return s.transport.Subscribe(s.sessionID, transport.DefaultEventBuffer)
}

func (s *Session) fail(err error) {
	s.teardown.Do(func() {
		if err != nil {
			s.Log("session failed", "targetId", s.targetID, "error", err)
		} else {
			s.Log("session closed", "targetId", s.targetID)
		}
		if s.unsubscribeHandle != nil {
			s.unsubscribeHandle()
		}
		if s.cancel != nil {
			s.cancel(err)
		}
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

func NewSession(transport *transport.Transport, targetID target.TargetID, timeout time.Duration) (*Session, error) {
	session := newSession(transport, targetID, timeout)
	session.context, session.cancel = context.WithCancelCause(transport.Context())
	sessionID, err := session.attachToTarget(targetID)
	if err != nil {
		session.fail(err)
		return nil, err
	}
	session.sessionID = string(sessionID)
	session.startHandleLoop()
	if err = session.enableDefaults(); err != nil {
		// non necessary to detach from target, because session will be closed on error
		session.fail(err)
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
	s.unsubscribeHandle = unsubscribe
	go func() {
		if err := s.handle(channel); err != nil {
			s.fail(err)
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
	if s.sessionID == string(detachedFromTarget.SessionId) {
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

func awaitMessage[T any](s *Session, watcher func(transport.Message) (T, bool, error)) future.Future[T] {
	return future.Execute(func(resolve func(T), reject func(error), canceled <-chan struct{}) {
		channel, unsubscribe := s.Subscribe()
		defer unsubscribe()

		for {
			select {
			case <-canceled:
				return

			case <-s.context.Done():
				reject(context.Cause(s.context))
				return

			case value, ok := <-channel:
				if !ok {
					reject(ErrSubscriptionClosed)
					return
				}
				result, isMatched, err := watcher(value)
				if err != nil {
					reject(err)
					return
				}
				if isMatched {
					resolve(result)
					return
				}
			}
		}
	})
}

func awaitMethod[T any](s *Session, method string, match func(T) bool) future.Future[T] {
	return awaitMessage(s, func(value transport.Message) (T, bool, error) {
		var zero T
		if value.Method != method {
			return zero, false, nil
		}

		result, err := transport.Unmarshal[T](value)
		if err != nil {
			return zero, false, err
		}
		if !match(result) {
			return zero, false, nil
		}

		return result, true, nil
	})
}
