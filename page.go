package witness

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/ecwid/witness/pkg/devtool"
)

var blankPage = "about:blank"

// ID session's ID
func (session *CDPSession) ID() string {
	return session.id
}

// Query query element on page by css selector
func (session *CDPSession) Query(selector string) (Element, error) {
	element, err := session.query(nil, selector)
	if err != nil {
		return nil, err
	}
	return newElement(session, nil, element.ObjectID, element.Description), nil
}

// QueryAll queryAll elements on page by css selector
func (session *CDPSession) QueryAll(selector string) []Element {
	v, err := session.queryAll(nil, selector)
	if err != nil {
		return []Element{}
	}
	return v
}

// C searching selector (visible) with implicity wait timeout
func (session *CDPSession) C(selector string, visible bool) Element {
	el, err := session.Ticker(func() (interface{}, error) {
		new, err := session.Query(selector)
		if err != nil {
			return nil, err
		}
		if visible {
			v, err := new.IsVisible()
			if err != nil {
				return nil, err
			}
			if !v {
				return nil, ErrElementInvisible
			}
		}
		return new, nil
	})
	if err != nil {
		panic(err)
	}
	return el.(Element)
}

// Close close this sessions
func (session *CDPSession) Close() error {
	_, err := session.blockingSend("Target.closeTarget", Map{"targetId": session.targetID})
	// event 'Target.targetDestroyed' can be received early than message response
	if err != nil && err != ErrSessionClosed {
		return err
	}
	return nil
}

func (session *CDPSession) navigateHistory(delta int64) error {
	history, err := session.getNavigationHistory()
	if err != nil {
		return err
	}
	move := history.CurrentIndex + delta
	if move >= 0 && move < int64(len(history.Entries)) {
		return session.navigateToHistoryEntry(history.Entries[move].ID)
	}
	return nil
}

// Back navigate back
func (session *CDPSession) Back() error {
	return session.navigateHistory(-1)
}

// Forward navigate forward
func (session *CDPSession) Forward() error {
	return session.navigateHistory(+1)
}

// Navigate navigate to url
func (session *CDPSession) Navigate(urlStr string) error {
	eventFired := make(chan bool, 1)
	unsubscribe := session.subscribe("Page.loadEventFired", func(*Event) {
		select {
		case eventFired <- true:
		default:
		}
	})
	defer close(eventFired)
	defer unsubscribe()
	// do navigate strict for main frameID (same as targetID)
	// in case of frame's navigate the 'Page.loadEventFired' will never fires
	// to implement it you should probably wait for 'Page.frameNavigated'
	msg, err := session.blockingSend("Page.navigate", Map{
		"url":            urlStr,
		"transitionType": "typed",
		"frameId":        session.targetID,
	})
	if err != nil {
		return err
	}
	nav := new(devtool.NavigationResult)
	if err = msg.Unmarshal(nav); err != nil {
		return err
	}
	if nav.ErrorText != "" {
		return fmt.Errorf(nav.ErrorText)
	}
	if nav.LoaderID == "" {
		// no navigate need
		return nil
	}
	select {
	case <-eventFired:
		return session.setFrame(nav.FrameID)
	case <-time.After(session.client.Timeouts.Navigation):
		return ErrNavigateTimeout
	}
}

// Reload refresh current page ignores cache
func (session *CDPSession) Reload() error {
	eventFired := make(chan bool, 1)
	unsubscribe := session.subscribe("Page.loadEventFired", func(*Event) {
		select {
		case eventFired <- true:
		default:
		}
	})
	defer close(eventFired)
	defer unsubscribe()
	if _, err := session.blockingSend("Page.reload", Map{"ignoreCache": true}); err != nil {
		return err
	}
	select {
	case <-eventFired:
		// reload destroys all frames so we should switch to main frame
		session.MainFrame()
		return nil
	case <-time.After(session.client.Timeouts.Navigation):
		return ErrNavigateTimeout
	}
}

// Evaluate evaluate javascript code at context of web page
func (session *CDPSession) Evaluate(code string, async bool) (interface{}, error) {
	result, err := session.evaluate(code, 0, async, false)
	if err != nil {
		return "", err
	}
	return result.Value, nil
}

// GetNavigationEntry get current tab info
func (session *CDPSession) GetNavigationEntry() (*devtool.NavigationEntry, error) {
	history, err := session.getNavigationHistory()
	if err != nil {
		return nil, err
	}
	if history.CurrentIndex == -1 {
		return &devtool.NavigationEntry{URL: blankPage}, nil
	}
	return history.Entries[history.CurrentIndex], nil
}

// FitToWindow ...
func (session *CDPSession) fitToWindow() error {
	view, err := session.getLayoutMetrics()
	if err != nil {
		return err
	}
	return session.setDeviceMetricsOverride(&devtool.DeviceMetrics{
		Width:             view.LayoutViewport.ClientWidth,
		Height:            int64(math.Ceil(view.ContentSize.Height)),
		DeviceScaleFactor: 1,
		Mobile:            false,
	})
}

// CaptureScreenshot get screen of current page
func (session *CDPSession) CaptureScreenshot(format string, quality int8, fullPage bool, cb func() error) ([]byte, error) {
	if fullPage {
		if err := session.fitToWindow(); err != nil {
			return nil, err
		}
		defer session.clearDeviceMetricsOverride()
	}
	if cb != nil {
		if err := cb(); err != nil {
			return nil, err
		}
	}
	if err := session.setScrollbarsHidden(true); err != nil {
		return nil, err
	}
	msg, err := session.blockingSend("Page.captureScreenshot", Map{
		"format":      format,
		"quality":     quality,
		"fromSurface": true,
	})
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(msg.json().String("data"))
}

// NewTab ...
func (session *CDPSession) NewTab(url string) (string, error) {
	if url == "" {
		url = blankPage // headless chrome crash when url is empty
	}
	msg, err := session.blockingSend("Target.createTarget", Map{
		"url": url,
	})
	if err != nil {
		return "", err
	}
	return msg.json().String("targetId"), nil
}

// SwitchToTab switch to another tab (new independent session will be created)
func (session *CDPSession) SwitchToTab(id string) (*Session, error) {
	return session.client.newSession(id)
}

// GetTabs list of opened tabs in browser (targetID)
func (session *CDPSession) GetTabs() ([]string, error) {
	ts, err := session.client.getTargets()
	if err != nil {
		return nil, err
	}
	handles := make([]string, 0)
	for _, t := range ts {
		if t.Type == "page" {
			handles = append(handles, t.TargetID)
		}
	}
	return handles, nil
}

// IsClosed check is session (tab) closed
func (session *CDPSession) IsClosed() bool {
	select {
	case <-session.closed:
		return true
	default:
		return false
	}
}

// MainFrame switch context to main frame of page
func (session *CDPSession) MainFrame() error {
	return session.setFrame(session.targetID)
}

// SwitchToFrame switch context to frame
func (session *CDPSession) SwitchToFrame(frameID string) error {
	return session.setFrame(frameID)
}

// Activate activate current Target
func (session *CDPSession) Activate() error {
	_, err := session.blockingSend("Target.activateTarget", Map{"targetId": session.targetID})
	return err
}

// AddScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addScriptToEvaluateOnNewDocument
func (session *CDPSession) AddScriptToEvaluateOnNewDocument(source string) (string, error) {
	msg, err := session.blockingSend("Page.addScriptToEvaluateOnNewDocument", Map{"source": source})
	if err != nil {
		return "", err
	}
	return msg.json().String("identifier"), nil
}

// RemoveScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-removeScriptToEvaluateOnNewDocument
func (session *CDPSession) RemoveScriptToEvaluateOnNewDocument(identifier string) error {
	_, err := session.blockingSend("Page.removeScriptToEvaluateOnNewDocument", Map{"identifier": identifier})
	return err
}

// OnNewTabOpen subscribe to Target.targetCreated event and return channel with targetID
func (session *CDPSession) OnNewTabOpen() chan string {
	var (
		wg                = sync.WaitGroup{}
		targetCreatedChan = make(chan string, 1)
		stopChan          = make(chan struct{})
	)
	unsubscribe := session.subscribe("Target.targetCreated", func(e *Event) {
		wg.Add(1)
		defer wg.Done()
		targetCreated := new(devtool.TargetCreated)
		if err := json.Unmarshal(e.Params, targetCreated); err != nil {
			session.panic(err)
		}
		if targetCreated.TargetInfo.Type == "page" {
			select {
			case targetCreatedChan <- targetCreated.TargetInfo.TargetID:
				stopChan <- struct{}{}
			default:
			}
		}
	})
	go func() {
		select {
		case <-stopChan:
		case <-time.After(session.client.Timeouts.Navigation): // to close channel after timeout to prevent hangs
		}
		unsubscribe()
		wg.Wait()
		close(targetCreatedChan)
	}()

	return targetCreatedChan
}

// Listen subscribe to listen cdp events with methods name
// return channel with incomming events and func to unsubscribe
// channel will be closed after unsubscribe func call
func (session *CDPSession) Listen(methods ...string) (chan *Event, func()) {
	var (
		wg          = sync.WaitGroup{}
		eventsChan  = make(chan *Event, 1)
		interrupt   = make(chan struct{})
		unsubscribe = make([]func(), 0)
	)
	callback := func(e *Event) {
		wg.Add(1)
		defer wg.Done()
		select {
		case eventsChan <- e:
		case <-interrupt:
		}
	}
	for _, m := range methods {
		unsubscribe = append(unsubscribe, session.subscribe(m, callback))
	}
	return eventsChan, func() {
		close(interrupt)
		for _, un := range unsubscribe {
			un()
		}
		wg.Wait()
		close(eventsChan)
	}
}
