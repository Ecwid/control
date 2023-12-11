package control

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

const (
	Blank     = "about:blank"
	bindClick = "_on_click"
)

type Session struct {
	browser    BrowserContext
	id         target.SessionID
	tid        target.TargetID
	executions *sync.Map
	eventPool  chan transport.Event
	publisher  *transport.Publisher
	exitCode   error
	context    context.Context
	cancelCtx  func()
	detach     func()

	Network   Network
	Input     Input
	Emulation Emulation
	Animation Animation
}

func (s Session) Call(method string, send, recv interface{}) error {
	select {
	case <-s.context.Done():
		if s.exitCode != nil {
			return s.exitCode
		}
		return s.context.Err()
	default:
		return s.browser.Client.Call(string(s.id), method, send, recv)
	}
}

func (s Session) GetBrowserContext() BrowserContext {
	return s.browser
}

func (s Session) GetTargetID() target.TargetID {
	return s.tid
}

func (s Session) ID() string {
	return string(s.id)
}

func (s Session) Name() string {
	return s.ID()
}

func (s Session) Page() *Frame {
	return &Frame{id: common.FrameId(s.tid), session: &s}
}

func (s Session) Frame(id common.FrameId) (*Frame, error) {
	if _, ok := s.executions.Load(id); ok {
		return &Frame{id: id, session: &s}, nil
	}
	return nil, NoSuchFrameError{id: id}
}

func (s Session) Activate() error {
	return s.browser.ActivateTarget(s.tid)
}

func (s Session) Update(val transport.Event) error {
	select {
	case s.eventPool <- val:
	case <-s.context.Done():
	default:
		return errors.New("eventPool is full")
	}
	return nil
}

func (s *Session) handle(e transport.Event) error {
	switch e.Method {

	case "Runtime.executionContextCreated":
		var v = runtime.ExecutionContextCreated{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		frameID := common.FrameId((v.Context.AuxData.(map[string]interface{}))["frameId"].(string))
		s.executions.Store(frameID, v.Context.UniqueId)

	case "Target.targetCrashed":
		var v = target.TargetCrashed{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		return ErrTargetCrashed(v)

	case "Target.targetDestroyed":
		var v = target.TargetDestroyed{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		if v.TargetId == s.tid {
			return ErrTargetDestroyed
		}

	case "Target.detachedFromTarget":
		var v = target.DetachedFromTarget{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		if v.SessionId == s.id {
			return ErrDetachedFromTarget
		}

	}
	return s.publisher.Notify(e.Method, e)
}

func (s *Session) handleEventPool() {
	defer func() {
		s.detach() // detach from the transport updates
		s.cancelCtx()
	}()
	for e := range s.eventPool {
		if err := s.handle(e); err != nil {
			s.exitCode = err
			return
		}
	}
}

func (s Session) onBindingCalled(name string, function func(string)) (cancel func()) {
	return s.Subscribe("Runtime.bindingCalled", func(value transport.Event) error {
		bindingCalled := runtime.BindingCalled{}
		err := json.Unmarshal(value.Params, &bindingCalled)
		if err != nil {
			return err
		}
		if bindingCalled.Name == name {
			function(bindingCalled.Payload)
		}
		return nil
	})
}

func (s Session) Subscribe(event string, v func(e transport.Event) error) (cancel func()) {
	return s.publisher.Register(transport.NewSimpleObserver(event, v))
}

func (s Session) Close() error {
	return s.browser.CloseTarget(s.tid)
}

func (s Session) IsClosed() bool {
	select {
	case <-s.context.Done():
		return true
	default:
		return false
	}
}
