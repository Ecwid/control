package witness

import (
	"encoding/base64"
	"fmt"
	"math"
	"time"

	"github.com/ecwid/witness/pkg/devtool"
)

// Page ...
type Page interface {
	Doc() Element

	Navigate(string) error
	Reload() error
	GetNavigationEntry() (*devtool.NavigationEntry, error)
	Close() error
	IsClosed() bool
	MainFrame() error
	Listen(string) (chan []byte, func())
	ID() string

	AddScriptToEvaluateOnNewDocument(string) (string, error)
	RemoveScriptToEvaluateOnNewDocument(string) error
	TakeScreenshot(string, int8, *devtool.Viewport, bool) ([]byte, error)

	NewPage(string) (string, error)
	SwitchToPage(string) (Page, error)
	GetPages() ([]string, error)
	SubscribeOnWindowOpen() chan string

	Evaluate(string, bool) (interface{}, error)

	SetCookies(...devtool.Cookie) error
	ClearBrowserCookies() error
	Fetch([]*devtool.RequestPattern, func(*devtool.RequestPaused, *Proceed)) func()

	MouseMove(float64, float64) error
	SendKeys(...rune) error

	Ticker(call TickerFunc) (interface{}, error)
}

// ID session's ID
func (session *Session) ID() string {
	return session.id
}

// Doc ...
func (session *Session) Doc() Element {
	return session.document
}

// Close close this sessions
func (session *Session) Close() error {
	_, err := session.blockingSend("Target.closeTarget", Map{"targetId": session.targetID})
	// event 'Target.targetDestroyed' can be received early than message response
	if err != nil && err != ErrSessionClosed {
		return err
	}
	return nil
}

// Navigate navigate to url
func (session *Session) Navigate(urlStr string) error {
	eventFired := make(chan bool, 1)
	unsubscribe := session.subscribe("Page.domContentEventFired", func([]byte) {
		eventFired <- true
	})
	defer unsubscribe()
	msg, err := session.blockingSend("Page.navigate", Map{
		"url":            urlStr,
		"transitionType": "typed",
		"frameId":        session.frameID,
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
		close(eventFired)
	}
	select {
	case <-eventFired:
	case <-time.After(session.client.Timeouts.Navigation):
		return ErrNavigateTimeout
	}
	return session.createIsolatedWorld(nav.FrameID)
}

// Reload refresh current page ignores cache
func (session *Session) Reload() error {
	eventFired := make(chan bool, 1)
	unsubscribe := session.subscribe("Page.domContentEventFired", func([]byte) {
		eventFired <- true
	})
	defer unsubscribe()
	_, err := session.blockingSend("Page.reload", Map{"ignoreCache": true})
	if err != nil {
		return err
	}
	select {
	case <-eventFired:
	case <-time.After(session.client.Timeouts.Navigation):
		return ErrNavigateTimeout
	}
	session.MainFrame()
	return nil
}

// Evaluate evaluate javascript code at context of web page
func (session *Session) Evaluate(code string, async bool) (interface{}, error) {
	result, err := session.evaluate(code, 0, async)
	if err != nil {
		return "", err
	}
	return result.Value, nil
}

// GetNavigationEntry get current tab info
func (session *Session) GetNavigationEntry() (*devtool.NavigationEntry, error) {
	history, err := session.getNavigationHistory()
	if err != nil {
		return nil, err
	}
	if history.CurrentIndex == -1 {
		return &devtool.NavigationEntry{URL: "about:blank"}, nil
	}
	return history.Entries[history.CurrentIndex], nil
}

// TakeScreenshot get screen of current page
func (session *Session) TakeScreenshot(format string, quality int8, clip *devtool.Viewport, fullPage bool) ([]byte, error) {
	_, err := session.blockingSend("Target.activateTarget", Map{"targetId": session.targetID})
	if fullPage {
		view, err := session.getLayoutMetrics()
		if err != nil {
			return nil, err
		}
		defer session.blockingSend("Emulation.clearDeviceMetricsOverride", Map{})
		_, err = session.blockingSend("Emulation.setDeviceMetricsOverride", Map{
			"width":             int64(math.Ceil(view.ContentSize.Width)),
			"height":            int64(math.Ceil(view.ContentSize.Height)),
			"deviceScaleFactor": 1,
			"mobile":            false,
		})
		if err != nil {
			return nil, err
		}
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

// NewPage ...
func (session *Session) NewPage(url string) (string, error) {
	msg, err := session.blockingSend("Target.createTarget", Map{
		"url": url,
	})
	if err != nil {
		return "", err
	}
	return msg.json().String("targetId"), nil
}

// SwitchToPage switch to another tab (new independent session will be created)
func (session *Session) SwitchToPage(id string) (Page, error) {
	return session.client.newSession(id)
}

// GetPages list of opened tabs in browser (targetID)
func (session *Session) GetPages() ([]string, error) {
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
func (session *Session) IsClosed() bool {
	select {
	case <-session.closed:
		return true
	default:
		return false
	}
}

// MainFrame switch context to main frame of page
func (session *Session) MainFrame() error {
	return session.createIsolatedWorld(session.targetID)
}

// AddScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addScriptToEvaluateOnNewDocument
func (session *Session) AddScriptToEvaluateOnNewDocument(source string) (string, error) {
	msg, err := session.blockingSend("Page.addScriptToEvaluateOnNewDocument", Map{"source": source})
	if err != nil {
		return "", err
	}
	return msg.json().String("identifier"), nil
}

// RemoveScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-removeScriptToEvaluateOnNewDocument
func (session *Session) RemoveScriptToEvaluateOnNewDocument(identifier string) error {
	_, err := session.blockingSend("Page.removeScriptToEvaluateOnNewDocument", Map{"identifier": identifier})
	return err
}

// SetCPUThrottlingRate https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setCPUThrottlingRate
func (session *Session) SetCPUThrottlingRate(rate int) error {
	_, err := session.blockingSend("Emulation.setCPUThrottlingRate", Map{"rate": rate})
	return err
}

// SubscribeOnWindowOpen subscribe to Target.targetCreated event and return channel with targetID
func (session *Session) SubscribeOnWindowOpen() chan string {
	message := make(chan string, 1)
	var unsubscribe func()
	close := time.AfterFunc(session.client.Timeouts.Navigation, func() {
		close(message)
	})
	unsubscribe = session.subscribe("Target.targetCreated", func(msg []byte) {
		targetCreated := new(devtool.TargetCreated)
		if err := bytes(msg).Unmarshal(targetCreated); err != nil {
			session.panic(err)
		}
		if targetCreated.TargetInfo.Type == "page" {
			message <- targetCreated.TargetInfo.TargetID
			unsubscribe()
			close.Stop()
		}
	})
	return message
}

// Listen subscribe to listen cdp event with name
// return channel with incomming events and func to unsubscribe
func (session *Session) Listen(name string) (chan []byte, func()) {
	message := make(chan []byte, 1)
	unsubscribe := session.subscribe(name, func(msg []byte) {
		message <- msg
	})
	return message, func() {
		close(message)
		unsubscribe()
	}
}
