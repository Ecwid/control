package control

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/transport/observe"
)

type NoSuchElementError struct {
	Selector string
}

func (n NoSuchElementError) Error() string {
	return fmt.Sprintf("no such element `%s`", n.Selector)
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
	contextID int32
	remote    *runtime.RemoteObject // todo
	session   *Session
	parent    *Frame
	child     []*Frame
}

func (f *Frame) Session() *Session {
	return f.session
}

func (f *Frame) ID() common.FrameId {
	return f.id
}

func (f *Frame) walk(val func(*Frame) bool) {
	if !val(f) {
		return
	}
	for _, e := range f.child {
		e.walk(val)
	}
}

func (f Frame) Call(method string, send, recv interface{}) error {
	return f.Session().Call(method, send, recv)
}

func (f Frame) NewLifecycleEventCondition(event LifecycleEventType) Condition {
	var isInit = false
	return NewCondition(f.Session(), f.Session().Timeout, func(value observe.Value) (bool, error) {
		if value.Method == "Page.lifecycleEvent" {
			var v = new(page.LifecycleEvent)
			if err := json.Unmarshal(value.Params, v); err != nil {
				return false, err
			}
			if v.FrameId == f.id && v.Name == "init" {
				isInit = true
			}
			return isInit && v.FrameId == f.id && v.Name == string(event), nil
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
		return nil, NoSuchElementError{Selector: selector}
	}
	return createElement(selector, object, f), nil
}

func (f *Frame) QuerySelectorAll(selector string) ([]*Element, error) {
	selector = safeSelector(selector)
	var array, err = f.evaluate(`document.querySelectorAll("`+selector+`")`, true, false)
	if err != nil {
		return nil, err
	}
	if array == nil || array.Description == "NodeList(0)" {
		return nil, nil
	}
	all := make([]*Element, 0)
	descriptor, err := f.getProperties(array.ObjectId, true, false)
	if err != nil {
		return nil, err
	}
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		all = append(all, createElement(selector, d.Value, f))
	}
	return all, nil
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
	var cid = runtime.ExecutionContextId(atomic.LoadInt32(&f.contextID))
	val, err := runtime.Evaluate(f, runtime.EvaluateArgs{
		Expression:            expression,
		IncludeCommandLineAPI: true,
		ContextId:             cid,
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

// NavigateHistory -1 = Back, +1 = Forward
func (f Frame) NavigateHistory(delta int) error {
	val, err := page.GetNavigationHistory(f)
	if err != nil {
		return err
	}
	move := val.CurrentIndex + delta
	if move >= 0 && move < len(val.Entries) {
		return page.NavigateToHistoryEntry(f, page.NavigateToHistoryEntryArgs{
			EntryId: val.Entries[move].Id,
		})
	}
	return nil
}

func (f *Frame) WithTimeout(retry func() error) error {
	var err error
	for start := time.Now(); time.Since(start) < f.Session().Timeout; {
		if err = retry(); err == nil {
			return nil
		}
		time.Sleep(f.Session().PoolingEvery)
	}
	return err
}
