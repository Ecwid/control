package control

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport/observe"
)

const (
	blankPage = "about:blank"
	bindClick = "_on_click"
)

type Session struct {
	browser    *BrowserContext
	id         target.SessionID
	tid        target.TargetID
	runtime    *dict
	eventPool  chan observe.Value
	Ctx        context.Context
	exit       func()
	exitCode   error
	observable *observe.Observable
	guid       *uint64 // observers incremental id
	Timeout    time.Duration
	Network    Network
	Input      Input
	Emulation  Emulation
}

func (s Session) Call(method string, send, recv interface{}) error {
	select {
	case <-s.Ctx.Done():
		return s.ExitCode()
	default:
		return s.browser.Client.Call(s.Ctx, string(s.id), method, send, recv)
	}
}

func (s Session) GetBrowserContext() *BrowserContext {
	return s.browser
}

func (s Session) ID() string {
	return string(s.id)
}

func (s Session) Event() string {
	return s.ID()
}

func (s Session) Page() *Frame {
	return &Frame{id: common.FrameId(s.tid), session: &s}
}

func (s Session) Frame(id common.FrameId) (*Frame, error) {
	if s.runtime.Load(id) != -1 {
		return &Frame{id: id, session: &s}, nil
	}
	return nil, NoSuchFrameError{id: id}
}

func (s Session) Activate() error {
	return s.browser.ActivateTarget(s.tid)
}

func (s Session) Notify(val observe.Value) {
	s.eventPool <- val
}

func (s *Session) handle(e observe.Value) error {
	switch e.Method {

	case "Runtime.executionContextCreated":
		var v = runtime.ExecutionContextCreated{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		frameID := common.FrameId((v.Context.AuxData.(map[string]interface{}))["frameId"].(string))
		s.runtime.Store(frameID, v.Context.Id)

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
	s.observable.Notify(e.Method, e)
	return nil
}

func (s *Session) lifecycle() {
	defer func() {
		s.browser.Client.Unregister(s)
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
	return s.Subscribe("Runtime.bindingCalled", false, func(value observe.Value) {
		bindingCalled := runtime.BindingCalled{}
		_ = json.Unmarshal(value.Params, &bindingCalled)
		if bindingCalled.Name == name {
			function(bindingCalled.Payload)
		}
	})
}

func (s Session) Subscribe(event string, inSeparateThread bool, v func(e observe.Value)) (cancel func()) {
	var (
		uid = atomic.AddUint64(s.guid, 1)
		val = observe.NewSimpleObserver(fmt.Sprintf("%d", uid), event, v)
	)
	if inSeparateThread {
		s.observable.Register(observe.AsyncSimpleObserver(val))
	} else {
		s.observable.Register(val)
	}
	return func() {
		s.observable.Unregister(val)
	}
}

func (s Session) Close() error {
	return s.browser.CloseTarget(s.tid)
}

func (s Session) IsClosed() bool {
	select {
	case <-s.Ctx.Done():
		return true
	default:
		return false
	}
}

func (s Session) ExitCode() error {
	if s.exitCode != nil {
		return s.exitCode
	}
	return s.Ctx.Err()
}

func (s Session) NewTargetCreatedCondition(createdTargetID *target.TargetID) *Promise {
	return s.NewEventCondition("Target.targetCreated", func(value observe.Value) (bool, error) {
		var v = target.TargetCreated{}
		if err := json.Unmarshal(value.Params, &v); err != nil {
			return false, err
		}
		if v.TargetInfo.Type == "page" && v.TargetInfo.OpenerId == s.tid {
			if createdTargetID != nil {
				*createdTargetID = v.TargetInfo.TargetId
			}
			return true, nil
		}
		return false, nil
	})
}

type dict struct {
	sync.Mutex
	values map[common.FrameId]runtime.ExecutionContextId
}

func newDict() *dict {
	return &dict{values: map[common.FrameId]runtime.ExecutionContextId{}}
}

func (m *dict) Store(key common.FrameId, val runtime.ExecutionContextId) {
	m.Lock()
	defer m.Unlock()
	m.values[key] = val
}

func (m *dict) Load(key common.FrameId) runtime.ExecutionContextId {
	m.Lock()
	defer m.Unlock()
	if val, ok := m.values[key]; ok {
		return val
	}
	return -1
}
