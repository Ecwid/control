# witness (PoC)
witness is golang client driving Chrome browser using the Chrome DevTools Protocol.
Witness has Selenium like interface. It is experimental project, backward compatibility is not guaranteed!

## Installation
`go get -u github.com/ecwid/witness`

## How to use

Here is an example of using:
```go
package main

import (
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func main() {
	chrome, _ := chrome.New("--headless")
	defer chrome.Close()
	session, err := chrome.CDP.DefaultSession()
	if err != nil {
		panic(err)
	}

	chrome.CDP.Logging.Level = witness.LevelProtocolMessage
	// Implicitly affected only C() function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	session.Page.Navigate("https://my.ecwid.com")

	session.Page.C("[name='email']", true).Type("test@example.com")
	session.Page.C("[name='password']", true).Type("xxxxxx")
	session.Page.C("button.btn-primary", true).Click()
}
```

Implemented methods:
```go
type Session struct {
	Network   Network
	Input     Input
	Runtime   Runtime
	Page      Page
	Message   Message
	Tabs      Tab
	Emulation Emulation
}

type Network interface {
	SetCookies(...*devtool.Cookie) error
	ClearBrowserCookies() error
	Intercept([]*devtool.RequestPattern, func(*devtool.RequestPaused, Interceptor)) func()
	SetOffline(bool) error
	SetThrottling(int, int, int) error
}

type Tab interface {
	NewTab(string) (string, error)
	SwitchToTab(string) (*Session, error)
	GetTabs() ([]string, error)
	OnNewTabOpen() chan string
}

type Input interface {
	MouseMove(float64, float64) error
	SendKeys(...rune) error
	InsertText(string) error
}

type Runtime interface {
	Evaluate(string, bool) (interface{}, error)
}

type Message interface {
	BlockingSend(method string, send interface{}) ([]byte, error)
	Listen(...string) (chan *Event, func())
}

type Emulation interface {
	SetCPUThrottlingRate(rate int) error
}

type Page interface {
	Findable

	Navigate(string) error
	Reload() error
	GetNavigationEntry() (*devtool.NavigationEntry, error)
	Close() error
	IsClosed() bool
	MainFrame() error
	SwitchToFrame(string) error

	ID() string

	AddScriptToEvaluateOnNewDocument(string) (string, error)
	RemoveScriptToEvaluateOnNewDocument(string) error
	TakeScreenshot(string, int8, *devtool.Viewport, bool) ([]byte, error)

	Ticker(call TickerFunc) (interface{}, error)
}

type Findable interface {
	C(string, bool) Element // Select by CSS selector
	Query(string) (Element, error)
	QueryAll(string) []Element
}

// Element element interface
type Element interface {
	Findable

	Click() error
	Hover() error
	Type(string, ...rune) error
	Upload(...string) error
	Clear() error
	Select(...string) error
	Checkbox(bool) error
	SetAttr(string, string) error
	Call(string, ...interface{}) (interface{}, error)
	Focus() error

	IsVisible() (bool, error)
	GetText() (string, error)
	GetAttr(attr string) (string, error)
	GetRectangle() (*devtool.Rect, error)
	GetComputedStyle(string) (string, error)
	GetSelected(bool) ([]string, error)
	IsChecked() (bool, error)
	GetEventListeners() ([]string, error)
	GetFrameID() (string, error)

	ObserveMutation(attributes, childList, subtree bool) (chan string, chan error)
	Release() error
}
```

See https://github.com/Ecwid/witness/tree/master/examples for more examples