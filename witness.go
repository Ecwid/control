package witness

import (
	"github.com/ecwid/witness/pkg/devtool"
	"github.com/ecwid/witness/pkg/mobile"
)

// Session entry point session
type Session struct {
	Network   Network
	Input     Input
	Runtime   Runtime
	Page      Page
	Message   Message
	Tabs      Tab
	Emulation Emulation
}

// Network network domain
type Network interface {
	SetCookies(...*devtool.Cookie) error
	ClearBrowserCookies() error
	GetCookies(...string) ([]*devtool.Cookie, error)
	Intercept([]*devtool.RequestPattern, func(*devtool.RequestPaused, Interceptor)) func()
	SetOffline(e bool) error
	SetThrottling(int, int, int) error
	SetBlockedURLs([]string) error
	GetRequestPostData(string) (string, error)
	GetResponseBody(string) (string, error)
}

// Tab pages manage
type Tab interface {
	NewTab(string) (string, error)
	SwitchToTab(string) (*Session, error)
	GetTabs() ([]string, error)
	OnNewTabOpen() chan string
}

// Input input domain
type Input interface {
	MouseMove(float64, float64) error
	SendKeys(...rune) error
	InsertText(string) error
}

// Runtime runtime domain
type Runtime interface {
	Evaluate(string, bool, bool) (interface{}, error)
	TerminateExecution() error
}

// Message internal CDP methods
type Message interface {
	BlockingSend(method string, send interface{}) ([]byte, error)
	Listen(...string) (chan *Event, func())
}

// Emulation Emulation domain
type Emulation interface {
	SetCPUThrottlingRate(rate int) error
	SetUserAgent(userAgent string) error
	Emulate(*mobile.Device) error
}

// Page page domain
type Page interface {
	Findable

	Navigate(string) error
	Reload() error
	GetNavigationEntry() (*devtool.NavigationEntry, error)
	Close() error
	IsClosed() bool
	MainFrame() error
	SwitchToFrame(string) error
	Back() error
	Forward() error

	ID() string

	AddScriptToEvaluateOnNewDocument(string) (string, error)
	RemoveScriptToEvaluateOnNewDocument(string) error
	SetDownloadBehavior(devtool.DownloadBehavior, string) error
	CaptureScreenshot(string, int8, bool, func() error) ([]byte, error)
	Activate() error

	Ticker(call TickerFunc) (interface{}, error)

	GetLayoutMetrics() (*devtool.LayoutMetrics, error)
}

// Findable interface to find element
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
