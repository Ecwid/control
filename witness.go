package witness

import "github.com/ecwid/witness/pkg/devtool"

// Session entry point session
type Session struct {
	Network   Network
	Input     Input
	Runtime   Runtime
	Page      Page
	Core      Core
	Tabs      Tab
	Emulation Emulation
}

// Network network domain
type Network interface {
	SetCookies(...*devtool.Cookie) error
	ClearBrowserCookies() error
	Intercept([]*devtool.RequestPattern, func(*devtool.RequestPaused, *Intercepted)) func()
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
}

// Runtime runtime domain
type Runtime interface {
	Evaluate(string, bool) (interface{}, error)
}

// Core internal CDP methods
type Core interface {
	Listen(string) (chan []byte, func())
}

// Emulation Emulation domain
type Emulation interface {
	SetCPUThrottlingRate(rate int) error
}

// Page page domain
type Page interface {
	selectable

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

// selectable interface to find element
type selectable interface {
	C(string, bool) Element // Select by CSS selector
	Query(string) (Element, error)
	QueryAll(string) []Element
}

// Element element interface
type Element interface {
	selectable

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
