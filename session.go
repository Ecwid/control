package control

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

const (
	blankPage = "about:blank"
	bindClick = "_on_click"
)

type Session struct {
	browser    *BrowserContext
	id         target.SessionID
	tid        target.TargetID
	executions *sync.Map
	eventPool  chan transport.Event
	context    context.Context
	exit       func()
	exitCode   error
	publisher  *transport.Publisher
	guid       *uint64 // observers incremental id
	Network    Network
	Input      Input
	Emulation  Emulation
}

func (s Session) Call(method string, send, recv interface{}) error {
	select {
	case <-s.context.Done():
		return s.ExitCode()
	default:
		return s.browser.Client.Call(string(s.id), method, send, recv)
	}
}

func (s Session) GetBrowserContext() *BrowserContext {
	return s.browser
}

func (s Session) GetTargetID() target.TargetID {
	return s.tid
}

func (s Session) ID() string {
	return string(s.id)
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

func (s Session) Notify(val transport.Event) {
	select {
	case s.eventPool <- val:
	default:
		s.exitCode = ErrEventPoolFull
	}
}

func (s *Session) handle(e transport.Event) error {
	switch e.Method {

	case "Runtime.executionContextCreated":
		var v = runtime.ExecutionContextCreated{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		frameID := common.FrameId((v.Context.AuxData.(map[string]interface{}))["frameId"].(string))
		s.executions.Store(frameID, v.Context.Id)

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
	s.publisher.Notify(e.Method, e)
	return nil
}

func (s *Session) lifecycle() {
	defer func() {
		s.browser.Client.Unregister(transport.NewSimpleObserver(s.ID(), s.ID(), s.Notify))
		s.exit()
	}()
	for e := range s.eventPool {
		if err := s.handle(e); err != nil {
			s.exitCode = err
			return
		}
	}
}

func (s Session) onBindingCalled(name string, function func(string)) (cancel func()) {
	return s.Subscribe("Runtime.bindingCalled", func(value transport.Event) {
		bindingCalled := runtime.BindingCalled{}
		_ = json.Unmarshal(value.Params, &bindingCalled)
		if bindingCalled.Name == name {
			function(bindingCalled.Payload)
		}
	})
}

func (s Session) Subscribe(event string, v func(e transport.Event)) (cancel func()) {
	var (
		uid = atomic.AddUint64(s.guid, 1)
		val = transport.NewSimpleObserver(fmt.Sprintf("%d", uid), event, v)
	)
	s.publisher.Register(val)
	return func() {
		s.publisher.Unregister(val)
	}
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

func (s Session) ExitCode() error {
	if s.exitCode != nil {
		return s.exitCode
	}
	return s.context.Err()
}
