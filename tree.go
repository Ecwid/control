package control

import (
	"sync"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/target"
)

type ctxTree struct {
	*sync.Mutex
	session *Session
	root    *Frame
}

func createContextTree(s *Session, targetID target.TargetID) *ctxTree {
	var tree = &ctxTree{
		session: s,
		Mutex:   &sync.Mutex{},
	}
	tree.root = tree.initFrame(common.FrameId(targetID), nil)
	return tree
}

func (t *ctxTree) initFrame(id common.FrameId, parent *Frame) *Frame {
	return &Frame{
		id:        id,
		contextID: 0,
		remote:    nil,
		session:   t.session,
		parent:    parent,
		child:     []*Frame{},
	}
}

func (t *ctxTree) appendChild(parentID common.FrameId, newID common.FrameId) {
	t.find(parentID, func(f *Frame) {
		f.child = append(f.child, t.initFrame(newID, f))
	})
}

func (t *ctxTree) deleteNode(ID common.FrameId) {
	t.find(ID, func(f *Frame) {
		if parent := f.parent; parent != nil {
			for n, e := range parent.child {
				if e.id == ID {
					last := len(parent.child) - 1
					parent.child[n] = parent.child[last]
					parent.child = parent.child[:last]
					return
				}
			}
		}
	})
}

func (t *ctxTree) find(ID common.FrameId, fn func(f *Frame)) {
	t.Lock()
	defer t.Unlock()
	t.root.walk(func(e *Frame) bool {
		if e.id == ID {
			fn(e)
			return false
		}
		return true
	})
}
