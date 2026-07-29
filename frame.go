package control

import (
	"errors"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/page"
)

type LifecycleEventType string

const (
	LifecycleDOMContentLoaded              LifecycleEventType = "DOMContentLoaded"
	LifecycleIdleNetwork                   LifecycleEventType = "networkIdle"
	LifecycleFirstContentfulPaint          LifecycleEventType = "firstContentfulPaint"
	LifecycleFirstMeaningfulPaint          LifecycleEventType = "firstMeaningfulPaint"
	LifecycleFirstMeaningfulPaintCandidate LifecycleEventType = "firstMeaningfulPaintCandidate"
	LifecycleFirstPaint                    LifecycleEventType = "firstPaint"
	LifecycleFirstTextPaint                LifecycleEventType = "firstTextPaint"
	LifecycleInit                          LifecycleEventType = "init"
	LifecycleLoad                          LifecycleEventType = "load"
	LifecycleNetworkAlmostIdle             LifecycleEventType = "networkAlmostIdle"
)

var ErrNavigateNoLoader = errors.New("navigation to same address")

type Frame struct {
	node    *Node
	session *Session
	id      common.FrameId
	parent  *Frame
}

func (f Frame) GetSession() *Session {
	return f.session
}

func (f Frame) GetID() common.FrameId {
	return f.id
}

func (f Frame) executionContextID() string {
	return f.session.getFrameExecutionContextID(f.id)
}

func (f Frame) Call(method string, send, recv any) error {
	return f.session.Call(method, send, recv)
}

func (f *Frame) OwnerFrame() *Frame {
	return f
}

func (f *Frame) Parent() *Frame {
	return f.parent
}

func (f Frame) Log(msg string, args ...any) {
	args = append(args, "frameId", f.id)
	f.session.Log(msg, args...)
}

func (f Frame) Navigate(url string) error {
	nav, err := page.Navigate(f, page.NavigateArgs{
		Url:     url,
		FrameId: f.id,
	})
	if err != nil {
		return err
	}
	if nav.ErrorText != "" {
		return errors.New(nav.ErrorText)
	}
	if nav.LoaderId == "" {
		return ErrNavigateNoLoader
	}
	return nil
}

func (f Frame) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) error {
	return page.Reload(f, page.ReloadArgs{
		IgnoreCache:            ignoreCache,
		ScriptToEvaluateOnLoad: scriptToEvaluateOnLoad,
	})
}

func (f Frame) Evaluate(expression string, awaitPromise bool) Optional[any] {
	return conv[any](f.evaluate(expression, awaitPromise))
}

func (f Frame) Document() Optional[*Node] {
	node, err := f.evaluate("document", true)
	if err != nil {
		return Optional[*Node]{err: err}
	}
	if node != nil {
		if n, ok := node.(*Node); ok {
			n.frame = &f
			n.requestedSelector = "document"
			return Optional[*Node]{value: n}
		}
	}
	return Optional[*Node]{err: NoSuchSelectorError("document")}
}

func (f Frame) Query(cssSelector string) Optional[*Node] {
	doc, err := f.Document().Unwrap()
	if err != nil {
		return Optional[*Node]{err: err}
	}
	return doc.Query(cssSelector)
}

func (f Frame) QueryAll(cssSelector string) Optional[NodeList] {
	doc, err := f.Document().Unwrap()
	if err != nil {
		return Optional[NodeList]{err: err}
	}
	return doc.QueryAll(cssSelector)
}
