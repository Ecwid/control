package emulation

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/page"
)

/*
	Screen orientation.
*/
type ScreenOrientation struct {
	Type  string `json:"type"`
	Angle int    `json:"angle"`
}

/*

 */
type DisplayFeature struct {
	Orientation string `json:"orientation"`
	Offset      int    `json:"offset"`
	MaskLength  int    `json:"maskLength"`
}

/*

 */
type MediaFeature struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
	advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to
allow the next delayed task (if any) to run; pause: The virtual time base may not advance;
pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending
resource fetches.
*/
type VirtualTimePolicy string

/*
	Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
*/
type UserAgentBrandVersion struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

/*
	Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
Missing optional values will be filled in by the target with what it would normally use.
*/
type UserAgentMetadata struct {
	Brands          []*common.UserAgentBrandVersion `json:"brands,omitempty"`
	FullVersion     string                          `json:"fullVersion,omitempty"`
	Platform        string                          `json:"platform"`
	PlatformVersion string                          `json:"platformVersion"`
	Architecture    string                          `json:"architecture"`
	Model           string                          `json:"model"`
	Mobile          bool                            `json:"mobile"`
}

/*
	Enum of image types that can be disabled.
*/
type DisabledImageType string

type CanEmulateVal struct {
	Result bool `json:"result"`
}

type SetFocusEmulationEnabledArgs struct {
	Enabled bool `json:"enabled"`
}

type SetCPUThrottlingRateArgs struct {
	Rate float64 `json:"rate"`
}

type SetDefaultBackgroundColorOverrideArgs struct {
	Color *dom.RGBA `json:"color,omitempty"`
}

type SetDeviceMetricsOverrideArgs struct {
	Width              int                `json:"width"`
	Height             int                `json:"height"`
	DeviceScaleFactor  float64            `json:"deviceScaleFactor"`
	Mobile             bool               `json:"mobile"`
	Scale              float64            `json:"scale,omitempty"`
	ScreenWidth        int                `json:"screenWidth,omitempty"`
	ScreenHeight       int                `json:"screenHeight,omitempty"`
	PositionX          int                `json:"positionX,omitempty"`
	PositionY          int                `json:"positionY,omitempty"`
	DontSetVisibleSize bool               `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation  *ScreenOrientation `json:"screenOrientation,omitempty"`
	Viewport           *page.Viewport     `json:"viewport,omitempty"`
	DisplayFeature     *DisplayFeature    `json:"displayFeature,omitempty"`
}

type SetScrollbarsHiddenArgs struct {
	Hidden bool `json:"hidden"`
}

type SetDocumentCookieDisabledArgs struct {
	Disabled bool `json:"disabled"`
}

type SetEmitTouchEventsForMouseArgs struct {
	Enabled       bool   `json:"enabled"`
	Configuration string `json:"configuration,omitempty"`
}

type SetEmulatedMediaArgs struct {
	Media    string          `json:"media,omitempty"`
	Features []*MediaFeature `json:"features,omitempty"`
}

type SetEmulatedVisionDeficiencyArgs struct {
	Type string `json:"type"`
}

type SetGeolocationOverrideArgs struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Accuracy  float64 `json:"accuracy,omitempty"`
}

type SetIdleOverrideArgs struct {
	IsUserActive     bool `json:"isUserActive"`
	IsScreenUnlocked bool `json:"isScreenUnlocked"`
}

type SetPageScaleFactorArgs struct {
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

type SetScriptExecutionDisabledArgs struct {
	Value bool `json:"value"`
}

type SetTouchEmulationEnabledArgs struct {
	Enabled        bool `json:"enabled"`
	MaxTouchPoints int  `json:"maxTouchPoints,omitempty"`
}

type SetVirtualTimePolicyArgs struct {
	Policy                            VirtualTimePolicy     `json:"policy"`
	Budget                            float64               `json:"budget,omitempty"`
	MaxVirtualTimeTaskStarvationCount int                   `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
	WaitForNavigation                 bool                  `json:"waitForNavigation,omitempty"`
	InitialVirtualTime                common.TimeSinceEpoch `json:"initialVirtualTime,omitempty"`
}

type SetVirtualTimePolicyVal struct {
	VirtualTimeTicksBase float64 `json:"virtualTimeTicksBase"`
}

type SetLocaleOverrideArgs struct {
	Locale string `json:"locale,omitempty"`
}

type SetTimezoneOverrideArgs struct {
	TimezoneId string `json:"timezoneId"`
}

type SetDisabledImageTypesArgs struct {
	ImageTypes []DisabledImageType `json:"imageTypes"`
}

type SetUserAgentOverrideArgs struct {
	UserAgent         string                    `json:"userAgent"`
	AcceptLanguage    string                    `json:"acceptLanguage,omitempty"`
	Platform          string                    `json:"platform,omitempty"`
	UserAgentMetadata *common.UserAgentMetadata `json:"userAgentMetadata,omitempty"`
}
