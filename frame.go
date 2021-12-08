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
	id      common.FrameId // readonly
	session *Session
}

func (f Frame) Session() *Session {
	return f.session
}

func (f Frame) ID() common.FrameId {
	return f.id
}

func (f Frame) Call(method string, send, recv interface{}) error {
	return f.Session().Call(method, send, recv)
}

func (f Frame) NewLifecycleEventCondition(event LifecycleEventType) *Promise {
	var initialized = false
	return f.Session().NewEventCondition("Page.lifecycleEvent", func(value observe.Value) (bool, error) {
		var v = page.LifecycleEvent{}
		if err := json.Unmarshal(value.Params, &v); err != nil {
			return false, err
		}
		if v.FrameId == f.id && v.Name == "init" {
			initialized = true
		}
		return initialized && v.FrameId == f.id && v.Name == string(event), nil
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
	promise := f.NewLifecycleEventCondition(waitFor)
	defer promise.Stop()
	if err := navigate(); err != nil {
		return err
	}
	return promise.WaitWithTimeout(f.Session().Timeout)
}

// Reload refresh current page
func (f Frame) Reload(ignoreCache bool, scriptToEvaluateOnLoad string, waitFor LifecycleEventType) error {
	var reload = func() error {
		return page.Reload(f, page.ReloadArgs{
			IgnoreCache:            ignoreCache,
			ScriptToEvaluateOnLoad: scriptToEvaluateOnLoad,
		})
	}
	promise := f.NewLifecycleEventCondition(waitFor)
	defer promise.Stop()
	if err := reload(); err != nil {
		return err
	}
	return promise.WaitWithTimeout(f.Session().Timeout)
}

func safeSelector(v string) string {
	v = strings.TrimSpace(v)
	v = strings.ReplaceAll(v, `"`, `\"`)
	return v
}

func (f Frame) IsExist(selector string) bool {
	selector = safeSelector(selector)
	val, _ := f.evaluate(`document.querySelector("`+selector+`") != null`, true, false)
	if val == nil {
		return false
	}
	b, _ := primitiveRemoteObject(*val).Bool()
	return b
}

func (f Frame) QuerySelector(selector string) (*Element, error) {
	selector = safeSelector(selector)
	var object, err = f.evaluate(`document.querySelector("`+selector+`")`, true, false)
	if err != nil {
		return nil, err
	}
	if object.ObjectId == "" {
		return nil, NoSuchElementError{Selector: selector}
	}
	return f.constructElement(object)
}

func (f Frame) QuerySelectorAll(selector string) ([]*Element, error) {
	selector = safeSelector(selector)
	var array, err = f.evaluate(`document.querySelectorAll("`+selector+`")`, true, false)
	if err != nil {
		return nil, err
	}
	if array == nil || array.Description == "NodeList(0)" {
		return nil, nil
	}
	list := make([]*Element, 0)
	descriptor, err := f.getProperties(array.ObjectId, true, false)
	if err != nil {
		return nil, err
	}
	for _, d := range descriptor {
		if !d.Enumerable {
			continue
		}
		el, err1 := f.constructElement(d.Value)
		if err1 != nil {
			return nil, err1
		}
		list = append(list, el)
	}
	return list, nil
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
	var cid = f.session.runtime.Load(f.id)
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

func (f Frame) RequestDOMIdle(threshold, timeout time.Duration) error {
	script := fmt.Sprintf(functionDOMIdle, threshold.Milliseconds(), timeout.Milliseconds())
	_, err := f.Evaluate(script, true, false)
	switch v := err.(type) {
	case RuntimeError:
		if val, _ := v.Exception.Value.(string); val == "timeout" {
			return WaitTimeoutError{timeout: timeout}
		}
	}
	return err
}
