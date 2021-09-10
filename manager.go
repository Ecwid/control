package control

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/target"
)

type Manager struct {
	sess *Session
	mx   sync.Mutex
	head *Frame
}

func newManager(s *Session, targetID target.TargetID) *Manager {
	var m = &Manager{
		sess: s,
		mx:   sync.Mutex{},
	}
	m.head = m.newFrame(common.FrameId(targetID), nil)
	return m
}

func (manager *Manager) debug() {
	fmt.Println("******************************")
	i := 1
	manager.head.walk(0, func(f *Frame, level int) bool {
		fmt.Println(strings.Repeat("-", level+1), i, f.id, "-", f.contextID, "=")
		i++
		return true
	})
	fmt.Println("******************************")
}

func (manager *Manager) newFrame(ID common.FrameId, parent *Frame) *Frame {
	return &Frame{
		id:        ID,
		contextID: 0,
		remote:    nil,
		manager:   manager,
		parent:    parent,
		child:     []*Frame{},
	}
}

func (manager *Manager) add(parentID common.FrameId, newID common.FrameId) {
	manager.edit(parentID, func(f *Frame) {
		f.child = append(f.child, manager.newFrame(newID, f))
	})
}

func (manager *Manager) delete(ID common.FrameId) {
	manager.edit(ID, func(f *Frame) {
		if parent := f.parent; parent != nil {
			for n, e := range parent.child {
				if e.id == ID {
					_len := len(parent.child) - 1
					parent.child[n] = parent.child[_len]
					parent.child = parent.child[:_len]
					return
				}
			}
		}
	})
}

func (manager *Manager) edit(ID common.FrameId, fn func(f *Frame)) {
	manager.mx.Lock()
	defer manager.mx.Unlock()
	manager.head.walk(0, func(e *Frame, _ int) bool {
		if e.id == ID {
			fn(e)
			return false
		}
		return true
	})
}
