package control

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
	"github.com/ecwid/control/transport/observe"
)

const (
	blankPage = "about:blank"
)

type Session struct {
	transport  *transport.Client
	id         target.SessionID
	tid        target.TargetID
	runtime    *dict
	eventPool  chan observe.Value
	Ctx        context.Context
	exit       func()
	exitCode   error
	observable *observe.Observable
	guid       uint64 // observers incremental id
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
		return s.transport.Call(s.Ctx, string(s.id), method, send, recv)
	}
}

func New(t *transport.Client) *Session {
	var s = &Session{
		guid:       0,
		id:         "",
		transport:  t,
		eventPool:  make(chan observe.Value, 1000),
		observable: observe.New(),
		Timeout:    time.Second * 60,
	}
	s.Ctx, s.exit = context.WithCancel(t.Ctx)
	s.Input = Input{s: s}
	s.Network = Network{s: s}
	s.Emulation = Emulation{s: s}
	return s
}
func (s Session) GetTransport() *transport.Client {
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

func (s *Session) AttachToTarget(targetID target.TargetID) error {
	if s.id != "" {
		if err := target.DetachFromTarget(s, target.DetachFromTargetArgs{SessionId: s.id}); err != nil {
			return err
		}
	}
	val, err := target.AttachToTarget(s, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return err
	}

	// run session lifecycle
	s.tid = targetID
	s.id = val.SessionId
	s.runtime = newDict()
	go s.lifecycle()
	s.transport.Register(s)

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
	// maxPostDataSize - Longest post body size (in bytes) that would be included in requestWillBeSent notification
	if err = network.Enable(s, network.EnableArgs{MaxPostDataSize: 2 * 1024}); err != nil {
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

func (s Session) Page() *Frame {
	return &Frame{id: common.FrameId(s.tid), session: &s}
}

func (s Session) Frame(id common.FrameId) (*Frame, error) {
	if s.runtime.Load(id) != 0 {
		return &Frame{id: id, session: &s}, nil
	}
	return nil, NoSuchFrameError{id: id}
}

func (s Session) Activate() error {
	return target.ActivateTarget(s, target.ActivateTargetArgs{
		TargetId: s.tid,
	})
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
		s.transport.Unregister(s)
		s.exit()
	}()
	for e := range s.eventPool {
		if err := s.handle(e); err != nil {
			s.exitCode = err
			return
		}
	}
}

func (s *Session) Subscribe(event string, async bool, v func(e observe.Value)) (unsubscribe func()) {
	var (
		uid = atomic.AddUint64(&s.guid, 1)
		val = observe.NewSimpleObserver(fmt.Sprintf("%d", uid), event, v)
	)
	if async {
		s.observable.Register(observe.AsyncSimpleObserver(val))
	} else {
		s.observable.Register(val)
	}
	return func() {
		s.observable.Unregister(val)
	}
}

func (s Session) Close() error {
	err := target.CloseTarget(s, target.CloseTargetArgs{
		TargetId: s.tid,
	})
	/* Target.detachedFromTarget event may come before the response of CloseTarget call */
	if err == context.Canceled {
		return nil
	}
	return err
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

func (s *Session) NewTargetCreatedCondition(createdTargetID *target.TargetID) *Condition {
	return s.NewCondition(func(value observe.Value) (bool, error) {
		if value.Method == "Target.targetCreated" {
			var v = new(target.TargetCreated)
			if err := json.Unmarshal(value.Params, v); err != nil {
				return false, err
			}
			if v.TargetInfo.Type == "page" && v.TargetInfo.OpenerId == s.tid {
				if createdTargetID != nil {
					*createdTargetID = v.TargetInfo.TargetId
				}
				return true, nil
			}
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
	return m.values[key]
}
