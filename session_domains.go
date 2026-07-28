package control

import (
	"time"

	"github.com/ecwid/control/future"
	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/page"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/target"
)

func (s *Session) attachToTarget(targetID target.TargetID) (target.SessionID, error) {
	val, err := target.AttachToTarget(s, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return "", err
	}
	return val.SessionId, nil
}

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
	return awaitMethod(s, "Target.targetCreated", func(value target.TargetCreated) (bool, error) {
		return value.TargetInfo.Type == "page" && value.TargetInfo.OpenerId == s.targetID, nil
	})
}

func (s *Session) AttachToTarget(id target.TargetID) (*Session, error) {
	return NewSession(s.transport, id, s.timeout)
}

func (s *Session) CreatePageTargetTab(url string) (*Session, error) {
	if url == "" {
		url = Blank // headless chrome crash when url is empty
	}
	r, err := target.CreateTarget(s, target.CreateTargetArgs{Url: url})
	if err != nil {
		return nil, err
	}
	return s.AttachToTarget(r.TargetId)
}

func (s *Session) Activate() error {
	return target.ActivateTarget(s, target.ActivateTargetArgs{TargetId: s.targetID})
}

func (s *Session) Close() error {
	return s.CloseTarget(s.targetID)
}

func (s *Session) CloseTarget(id target.TargetID) (err error) {
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
	return conv[string](s.getCurrentURL())
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
	return awaitMethod(s, "Runtime.bindingCalled", func(value runtime.BindingCalled) (bool, error) {
		return value.Name == fn, nil
	})
}

func (s *Session) Click(point Point) error {
	return s.mouse.Click(MouseLeft, point, time.Millisecond*85)
}

func (s *Session) MouseDown(point Point) error {
	return s.mouse.Down(MouseLeft, point)
}

func (s *Session) Swipe(from, to Point) error {
	return s.touch.Swipe(from, to)
}

func (s *Session) Hover(point Point) error {
	return s.mouse.Move(MouseNone, point)
}
