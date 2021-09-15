package control

import (
	"encoding/json"
	"fmt"
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

type Map map[string]interface{}

const (
	blankPage = "about:blank"
)

type Session struct {
	transport    *transport.Client
	id           target.SessionID
	tid          target.TargetID
	tree         *ctxTree
	eventPool    chan observe.Value
	exited       chan struct{}
	exitCode     error
	observable   *observe.Observable
	obsuid       uint64 // observers incremental id
	Timeout      time.Duration
	PoolingEvery time.Duration
	Network      Network
	Input        Input
	Emulation    Emulation
}

func (s Session) Call(method string, send, recv interface{}) error {
	select {
	case <-s.exited:
		return s.exitCode
	default:
		return s.transport.Call(string(s.id), method, send, recv)
	}
}

func New(t *transport.Client) *Session {
	var hlSess = &Session{
		obsuid:       0,
		id:           "",
		transport:    t,
		eventPool:    make(chan observe.Value, 999),
		observable:   observe.New(),
		exited:       make(chan struct{}, 1),
		Timeout:      time.Second * 60,
		PoolingEvery: time.Millisecond * 500,
	}
	hlSess.Input = Input{s: hlSess}
	hlSess.Network = Network{s: hlSess}
	hlSess.Emulation = Emulation{s: hlSess}
	return hlSess
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
	s.tree = createContextTree(s, targetID)
	go s.lifecycle()
	s.transport.Register(s)

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
	return s.tree.root
}

func (s Session) Frame(id common.FrameId) (*Frame, error) {
	var frame *Frame
	s.tree.find(id, func(f *Frame) {
		frame = f
	})
	if frame == nil {
		return nil, fmt.Errorf("frame with id = %s not found", id)
	}
	return frame, nil
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

	case "Page.frameAttached":
		var v = page.FrameAttached{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		s.tree.appendChild(v.ParentFrameId, v.FrameId)

	case "Page.frameDetached":
		var v = page.FrameDetached{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		s.tree.deleteNode(v.FrameId)

	case "Runtime.executionContextCreated":
		var v = runtime.ExecutionContextCreated{}
		if err := json.Unmarshal(e.Params, &v); err != nil {
			return err
		}
		frameID := common.FrameId((v.Context.AuxData.(map[string]interface{}))["frameId"].(string))
		s.tree.find(frameID, func(f *Frame) {
			atomic.StoreInt32(&f.contextID, int32(v.Context.Id))
		})

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
		close(s.exited)
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
		uid = atomic.AddUint64(&s.obsuid, 1)
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

func (s *Session) NewTargetCreatedCondition(createdTargetID *target.TargetID) *Condition {
	return s.NewCondition(func(value observe.Value) (bool, error) {
		if value.Method == "Target.targetCreated" {
			var v = new(target.TargetCreated)
			if err := json.Unmarshal(value.Params, v); err != nil {
				return false, err
			}
			if v.TargetInfo.Type == "page" && v.TargetInfo.OpenerId == s.tid {
				*createdTargetID = v.TargetInfo.TargetId
				return true, nil
			}
		}
		return false, nil
	})
}
