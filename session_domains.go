package control

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ecwid/control/future"
	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

func (s *Session) CaptureScreenshot(format string, quality int, clip *page.Viewport, fromSurface, captureBeyondViewport, optimizeForSpeed bool) ([]byte, error) {
	val, err := page.CaptureScreenshot(s, page.CaptureScreenshotArgs{
		Format:                format,
		Quality:               quality,
		Clip:                  clip,
		FromSurface:           fromSurface,
		CaptureBeyondViewport: captureBeyondViewport,
		OptimizeForSpeed:      optimizeForSpeed,
	})
	if err != nil {
		return nil, err
	}
	return val.Data, nil
}

func (s *Session) SetDownloadBehavior(behavior string, downloadPath string, eventsEnabled bool) error {
	return browser.SetDownloadBehavior(s, browser.SetDownloadBehaviorArgs{
		Behavior:      behavior,
		DownloadPath:  downloadPath,
		EventsEnabled: eventsEnabled, // default false
	})
}

func (s *Session) GetTargetCreated() future.Future[target.TargetCreated] {
	return subscribeToMethod(s, "Target.targetCreated", func(value target.TargetCreated) (bool, error) {
		return value.TargetInfo.Type == "page" && value.TargetInfo.OpenerId == s.targetID, nil
	})
}

func (s *Session) CreatePageTargetTab(url string) (*Session, error) {
	return s.browser.NewTab(url)
}

func (s *Session) Activate() error {
	return target.ActivateTarget(s, target.ActivateTargetArgs{TargetId: s.targetID})
}

func (s *Session) Close() error {
	return s.closeTarget(s.targetID)
}

func (s *Session) closeTarget(id target.TargetID) (err error) {
	err = target.CloseTarget(s, target.CloseTargetArgs{TargetId: id})
	/* Target.detachedFromTarget event may come before the response of CloseTarget call */
	if err == ErrTargetDetached {
		return nil
	}
	return err
}

func (s *Session) GetLayout() Optional[page.GetLayoutMetricsVal] {
	view, err := page.GetLayoutMetrics(s)
	if err != nil {
		return Optional[page.GetLayoutMetricsVal]{err: err}
	}
	return Optional[page.GetLayoutMetricsVal]{value: *view}
}

func (s *Session) GetNavigationEntry() Optional[page.NavigationEntry] {
	val, err := page.GetNavigationHistory(s)
	if err != nil {
		return Optional[page.NavigationEntry]{err: err}
	}
	if val.CurrentIndex == -1 {
		return Optional[page.NavigationEntry]{value: page.NavigationEntry{Url: Blank}}
	}
	return Optional[page.NavigationEntry]{value: *val.Entries[val.CurrentIndex]}
}

func (s *Session) GetCurrentURL() Optional[string] {
	return optional[string](s.getCurrentURL())
}

func (s *Session) getCurrentURL() (string, error) {
	e, err := s.GetNavigationEntry().Unwrap()
	if err != nil {
		return "", err
	}
	return e.Url, nil
}

func (s *Session) NavigateHistory(delta int) error {
	val, err := page.GetNavigationHistory(s)
	if err != nil {
		return err
	}
	move := val.CurrentIndex + delta
	if move >= 0 && move < len(val.Entries) {
		return page.NavigateToHistoryEntry(s, page.NavigateToHistoryEntryArgs{
			EntryId: val.Entries[move].Id,
		})
	}
	return nil
}

func (s *Session) GetBindingCalled(fn string) future.Future[runtime.BindingCalled] {
	return subscribeToMethod(s, "Runtime.bindingCalled", func(value runtime.BindingCalled) (bool, error) {
		return value.Name == fn, nil
	})
}

func (s *Session) Click(point Point) error {
	return NewMouse(s).Click(MouseLeft, point, time.Millisecond*85)
}

func (s *Session) MouseDown(point Point) error {
	return NewMouse(s).Down(MouseLeft, point)
}

func (s *Session) Swipe(from, to Point) error {
	return NewTouch(s).Swipe(from, to)
}

func (s *Session) Hover(point Point) error {
	return NewMouse(s).Move(MouseNone, point)
}

func shouldTrackNetworkRequest(willBeSent network.RequestWillBeSent) bool {
	switch willBeSent.Type {
	case network.ResourceType("WebSocket"), network.ResourceType("EventSource"):
		return false
	}
	if willBeSent.Request != nil && strings.HasPrefix(willBeSent.Request.Url, "blob:") {
		return false
	}
	return true
}

func (s *Session) NetworkRequest(timeout time.Duration, matcher func(network.RequestWillBeSent) bool, init func()) (network.ResponseReceived, error) {
	channel, unsubscribe := s.Subscribe()
	defer unsubscribe()

	if init != nil {
		init()
	}

	ctxTo, cancel := context.WithTimeout(s.Context(), timeout)
	defer cancel()

	var requestId network.RequestId
	var response *network.ResponseReceived
	for {
		select {

		case <-ctxTo.Done():
			return network.ResponseReceived{}, context.Cause(ctxTo)

		case value, ok := <-channel:
			if !ok {
				if err := context.Cause(s.Context()); err != nil {
					return network.ResponseReceived{}, err
				}
				return network.ResponseReceived{}, ErrSubscriptionClosed
			}

			switch value.Method {

			case "Network.requestWillBeSent":
				willBeSent, err := transport.Unmarshal[network.RequestWillBeSent](value)
				if err != nil {
					return network.ResponseReceived{}, err
				}
				if requestId != "" {
					continue
				}
				if !shouldTrackNetworkRequest(willBeSent) {
					continue
				}
				if matcher != nil && !matcher(willBeSent) {
					continue
				}
				requestId = willBeSent.RequestId

			case "Network.responseReceived":
				if requestId == "" {
					continue
				}
				responseReceived, err := transport.Unmarshal[network.ResponseReceived](value)
				if err != nil {
					return network.ResponseReceived{}, err
				}
				if responseReceived.RequestId == requestId {
					response = &responseReceived
					if responseReceived.Response != nil && responseReceived.Response.Status >= 400 {
						return network.ResponseReceived{}, errors.New(http.StatusText(responseReceived.Response.Status))
					}
				}

			case "Network.loadingFailed":
				if requestId == "" {
					continue
				}
				loadingFailed, err := transport.Unmarshal[network.LoadingFailed](value)
				if err != nil {
					return network.ResponseReceived{}, err
				}
				if loadingFailed.RequestId == requestId {
					return network.ResponseReceived{}, errors.New(loadingFailed.ErrorText)
				}

			case "Network.loadingFinished":
				if requestId == "" {
					continue
				}
				loadingFinished, err := transport.Unmarshal[network.LoadingFinished](value)
				if err != nil {
					return network.ResponseReceived{}, err
				}
				if loadingFinished.RequestId == requestId {
					if response == nil {
						return network.ResponseReceived{}, errors.New("network response missing")
					}
					return *response, nil
				}
			}
		}
	}
}

type InflightDeadlineExceededError struct {
	Inflight map[network.RequestId]network.RequestWillBeSent
}

func (e InflightDeadlineExceededError) Error() string {
	var sb strings.Builder
	sb.WriteString("network idle timeout, inflight requests:\n")
	for _, req := range e.Inflight {
		sb.WriteString(req.Request.Url)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (s *Session) NetworkIdle(timeout, threshold time.Duration, init func()) error {
	channel, unsubscribe := s.Subscribe()
	defer unsubscribe()

	if init != nil {
		init()
	}

	ctxTo, cancel := context.WithTimeout(s.Context(), timeout)
	defer cancel()

	start := time.Now()
	totalRequests := 0
	inflight := map[network.RequestId]network.RequestWillBeSent{}
	idleTimer := time.NewTimer(threshold)
	stopIdleTimer := func() {
		if !idleTimer.Stop() {
			select {
			case <-idleTimer.C:
			default:
			}
		}
	}
	resetIdleTimer := func() {
		stopIdleTimer()
		idleTimer.Reset(threshold)
	}

	defer func() {
		stopIdleTimer()
		s.Log("NetworkIdle", "requests", totalRequests, "inflight", inflight, "duration", time.Since(start).String())
	}()

	for {
		select {

		case <-ctxTo.Done():
			ctxErr := context.Cause(ctxTo)
			if ctxErr == context.DeadlineExceeded {
				return InflightDeadlineExceededError{Inflight: inflight}
			}
			return ctxErr

		case value, ok := <-channel:
			if !ok {
				if err := context.Cause(s.Context()); err != nil {
					return err
				}
				return ErrSubscriptionClosed
			}

			switch value.Method {

			case "Network.requestWillBeSent":
				willBeSent, err := transport.Unmarshal[network.RequestWillBeSent](value)
				if err != nil {
					return err
				}
				if !shouldTrackNetworkRequest(willBeSent) {
					continue
				}
				if _, exists := inflight[willBeSent.RequestId]; !exists {
					totalRequests++
				}
				inflight[willBeSent.RequestId] = willBeSent
				resetIdleTimer()

			case "Network.loadingFailed":
				loadingFailed, err := transport.Unmarshal[network.LoadingFailed](value)
				if err != nil {
					return err
				}
				if _, exists := inflight[loadingFailed.RequestId]; exists {
					delete(inflight, loadingFailed.RequestId)
					resetIdleTimer()
				}

			case "Network.loadingFinished":
				loadingFinished, err := transport.Unmarshal[network.LoadingFinished](value)
				if err != nil {
					return err
				}
				if _, exists := inflight[loadingFinished.RequestId]; exists {
					delete(inflight, loadingFinished.RequestId)
					resetIdleTimer()
				}

			case "Page.frameDetached":
				frameDetached, err := transport.Unmarshal[page.FrameDetached](value)
				if err != nil {
					return err
				}
				for requestID, willBeSent := range inflight {
					if willBeSent.FrameId == frameDetached.FrameId {
						delete(inflight, requestID)
					}
				}
			}

		case <-idleTimer.C:
			if len(inflight) == 0 {
				return nil
			}
			resetIdleTimer()
		}
	}
}
