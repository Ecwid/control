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
	page, err := chrome.CDP.DefaultPage()
	if err != nil {
		panic(err)
	}

	chrome.CDP.Logging.Level = witness.LevelProtocolMessage
	// Implicitly affected only C(selector, visible) function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	page.Navigate("https://my.ecwid.com")

	page.C("[name='email']", true).Type("test@example.com")
	page.C("[name='password']", true).Type("xxxxxx")
	page.C("button.btn-primary", true).Click()
}
```

Implemented element methods:
```go
type Element interface {
	C(string, bool) Element // Select by CSS selector
	Query(string) (Element, error)
	QueryAll(string) []Element

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

	Release() error
}
```

Page's methods:
```go
type Page interface {
	C(string, bool) Element // Select by CSS selector
	Query(string) (Element, error)
	QueryAll(string) []Element

	Navigate(string) error
	Reload() error
	GetNavigationEntry() (*devtool.NavigationEntry, error)
	Close() error
	IsClosed() bool
	MainFrame() error
	SwitchToFrame(string) error
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

	SetCookies(...*devtool.Cookie) error
	ClearBrowserCookies() error
	Fetch([]*devtool.RequestPattern, func(*devtool.RequestPaused, *Proceed)) func()

	MouseMove(float64, float64) error
	SendKeys(...rune) error

	Ticker(call TickerFunc) (interface{}, error)
}
```


See https://github.com/Ecwid/witness/tree/master/examples for more examples