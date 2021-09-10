package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/transport/observe"
)

type NoSuchElementError struct {
	selector string
}

func (n NoSuchElementError) Error() string {
	return fmt.Sprintf("no such element `%s`", n.selector)
}

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

type Frame struct {
	id        common.FrameId
	contextID runtime.ExecutionContextId
	remote    *runtime.RemoteObject // todo
	manager   *Manager
	parent    *Frame
	child     []*Frame
}

func (f *Frame) Session() *Session {
	return f.manager.sess
}

func (f *Frame) walk(level int, val func(*Frame, int) bool) {
	if !val(f, level) {
		return
	}
	for _, e := range f.child {
		e.walk(level+1, val)
	}
}

func (f Frame) Call(method string, send, recv interface{}) error {
	return f.Session().Call(method, send, recv)
}

func (f Frame) NewLifecycleEventCondition(event LifecycleEventType) Condition {
	return NewCondition(f.Session(), f.Session().Timeout, func(value observe.Value) (bool, error) {
		if value.Method == "Page.lifecycleEvent" {
			var v = new(page.LifecycleEvent)
			if err := json.Unmarshal(value.Params, v); err != nil {
				return false, err
			}
			return v.FrameId == f.id && v.Name == string(event), nil
		}
		return false, nil
	})
}

func (f Frame) Navigate(url string, waitFor LifecycleEventType) error {
	var navigate = func() error {
		nav, err := page.Navigate(f, page.NavigateArgs{
			Url:            url,
			TransitionType: "typed",
			FrameId:        f.id,
		})
		if err != nil {
			return err
		}
		if nav.ErrorText != "" {
			return errors.New(nav.ErrorText)
		}
		if nav.LoaderId == "" {
			return ErrAlreadyNavigated
		}
		return nil
	}
	return f.NewLifecycleEventCondition(waitFor).Do(navigate)
}

// Reload refresh current page
func (f Frame) Reload(ignoreCache bool, scriptToEvaluateOnLoad string, waitFor LifecycleEventType) error {
	var reload = func() error {
		return page.Reload(f, page.ReloadArgs{
			IgnoreCache:            ignoreCache,
			ScriptToEvaluateOnLoad: scriptToEvaluateOnLoad,
		})
	}
	return f.NewLifecycleEventCondition(waitFor).Do(reload)
}

func (f *Frame) MustElement(selector string) Element {
	var (
		v   *Element
		err error
	)
	for start := time.Now(); time.Since(start) < f.Session().Timeout; {
		v, err = f.QuerySelector(selector)
		if err == nil {
			return *v
		}
		time.Sleep(f.Session().PoolingEvery)
	}
	panic(err)
}

func safeSelector(v string) string {
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, `"`, `\"`)
	return v
}

func (f *Frame) QuerySelector(selector string) (*Element, error) {
	selector = safeSelector(selector)
	var object, err = f.evaluate(`document.querySelector("`+selector+`")`, true, false)
	if err != nil {
		return nil, err
	}
	if object.ObjectId == "" {
		return nil, NoSuchElementError{selector: selector}
	}
	return &Element{frame: f, remote: object}, nil
}

type RuntimeError runtime.ExceptionDetails

func (r RuntimeError) Error() string {
	b, _ := json.Marshal(r)
	return fmt.Sprintf("%s", b)
}

func (f Frame) Evaluate(expression string, await, returnByValue bool) (interface{}, error) {
	val, err := f.evaluate(expression, await, returnByValue)
	if err != nil {
		return "", err
	}
	return val.Value, nil
}

func (f Frame) evaluate(expression string, await, returnByValue bool) (*runtime.RemoteObject, error) {
	val, err := runtime.Evaluate(f, runtime.EvaluateArgs{
		Expression:            expression,
		IncludeCommandLineAPI: true,
		ContextId:             f.contextID,
		AwaitPromise:          await,
		ReturnByValue:         returnByValue,
	})
	if err != nil {
		return nil, err
	}
	if val.ExceptionDetails != nil {
		return nil, RuntimeError(*val.ExceptionDetails)
	}
	return val.Result, nil
}

// GetNavigationEntry get current tab info
func (f Frame) GetNavigationEntry() (*page.NavigationEntry, error) {
	val, err := page.GetNavigationHistory(f)
	if err != nil {
		return nil, err
	}
	if val.CurrentIndex == -1 {
		return &page.NavigationEntry{Url: blankPage}, nil
	}
	return val.Entries[val.CurrentIndex], nil
}
