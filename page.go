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

var ErrNavigateNoLoader = errors.New("navigation to the same address")

type Queryable interface {
	Query(string) Optional[*Node]
	MustQuery(string) *Node
	QueryAll(string) Optional[NodeList]
	MustQueryAll(string) NodeList
	OwnerFrame() *Frame
}

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
	if value, ok := f.session.frames.Load(f.id); ok {
		return value.(string)
	}
	return ""
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

func (f Frame) MustNavigate(url string) {
	if err := f.Navigate(url); err != nil {
		panic(err)
	}
}

func (f Frame) Reload(ignoreCache bool, scriptToEvaluateOnLoad string) error {
	return page.Reload(f, page.ReloadArgs{
		IgnoreCache:            ignoreCache,
		ScriptToEvaluateOnLoad: scriptToEvaluateOnLoad,
	})
}

func (f Frame) MustReload(ignoreCache bool, scriptToEvaluateOnLoad string) {
	if err := f.Reload(ignoreCache, scriptToEvaluateOnLoad); err != nil {
		panic(err)
	}
}

func (f Frame) Evaluate(expression string, awaitPromise bool) Optional[any] {
	return optional[any](f.evaluate(expression, awaitPromise))
}

func (f Frame) Document() Optional[*Node] {
	opt := optional[*Node](f.evaluate("document", true))
	if opt.err == nil && opt.value == nil {
		opt.err = NoSuchSelectorError("document")
	}
	if opt.value != nil {
		opt.value.requestedSelector = "document"
	}
	return opt
}

func (f Frame) MustQuery(cssSelector string) *Node {
	return f.Query(cssSelector).MustGetValue()
}

func (f Frame) Query(cssSelector string) Optional[*Node] {
	doc, err := f.Document().Unwrap()
	if err != nil {
		return Optional[*Node]{err: err}
	}
	return doc.Query(cssSelector)
}

func (f Frame) MustQueryAll(cssSelector string) NodeList {
	return f.QueryAll(cssSelector).MustGetValue()
}

func (f Frame) QueryAll(cssSelector string) Optional[NodeList] {
	doc, err := f.Document().Unwrap()
	if err != nil {
		return Optional[NodeList]{err: err}
	}
	return doc.QueryAll(cssSelector)
}
