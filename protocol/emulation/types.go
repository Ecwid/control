package emulation

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/page"
)

/*
 */
type SafeAreaInsets struct {
	Top       int `json:"top,omitempty"`
	TopMax    int `json:"topMax,omitempty"`
	Left      int `json:"left,omitempty"`
	LeftMax   int `json:"leftMax,omitempty"`
	Bottom    int `json:"bottom,omitempty"`
	BottomMax int `json:"bottomMax,omitempty"`
	Right     int `json:"right,omitempty"`
	RightMax  int `json:"rightMax,omitempty"`
}

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
type DevicePosture struct {
	Type string `json:"type"`
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
Used to specify User Agent Client Hints to emulate. See https://wicg.github.io/ua-client-hints
*/
type UserAgentBrandVersion struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

/*
	Used to specify User Agent Client Hints to emulate. See https://wicg.github.io/ua-client-hints

Missing optional values will be filled in by the target with what it would normally use.
*/
type UserAgentMetadata struct {
	Brands          []*common.UserAgentBrandVersion `json:"brands,omitempty"`
	FullVersionList []*common.UserAgentBrandVersion `json:"fullVersionList,omitempty"`
	Platform        string                          `json:"platform"`
	PlatformVersion string                          `json:"platformVersion"`
	Architecture    string                          `json:"architecture"`
	Model           string                          `json:"model"`
	Mobile          bool                            `json:"mobile"`
	Bitness         string                          `json:"bitness,omitempty"`
	Wow64           bool                            `json:"wow64,omitempty"`
	FormFactors     []string                        `json:"formFactors,omitempty"`
}

/*
	Used to specify sensor types to emulate.

See https://w3c.github.io/sensors/#automation for more information.
*/
type SensorType string

/*
 */
type SensorMetadata struct {
	Available        bool    `json:"available,omitempty"`
	MinimumFrequency float64 `json:"minimumFrequency,omitempty"`
	MaximumFrequency float64 `json:"maximumFrequency,omitempty"`
}

/*
 */
type SensorReadingSingle struct {
	Value float64 `json:"value"`
}

/*
 */
type SensorReadingXYZ struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

/*
 */
type SensorReadingQuaternion struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
	W float64 `json:"w"`
}

/*
 */
type SensorReading struct {
	Single     *SensorReadingSingle     `json:"single,omitempty"`
	Xyz        *SensorReadingXYZ        `json:"xyz,omitempty"`
	Quaternion *SensorReadingQuaternion `json:"quaternion,omitempty"`
}

/*
 */
type PressureSource string

/*
 */
type PressureState string

/*
 */
type PressureMetadata struct {
	Available bool `json:"available,omitempty"`
}

/*
 */
type WorkAreaInsets struct {
	Top    int `json:"top,omitempty"`
	Left   int `json:"left,omitempty"`
	Bottom int `json:"bottom,omitempty"`
	Right  int `json:"right,omitempty"`
}

/*
 */
type ScreenId string

/*
	Screen information similar to the one returned by window.getScreenDetails() method,

see https://w3c.github.io/window-management/#screendetailed.
*/
type ScreenInfo struct {
	Left             int                `json:"left"`
	Top              int                `json:"top"`
	Width            int                `json:"width"`
	Height           int                `json:"height"`
	AvailLeft        int                `json:"availLeft"`
	AvailTop         int                `json:"availTop"`
	AvailWidth       int                `json:"availWidth"`
	AvailHeight      int                `json:"availHeight"`
	DevicePixelRatio float64            `json:"devicePixelRatio"`
	Orientation      *ScreenOrientation `json:"orientation"`
	ColorDepth       int                `json:"colorDepth"`
	IsExtended       bool               `json:"isExtended"`
	IsInternal       bool               `json:"isInternal"`
	IsPrimary        bool               `json:"isPrimary"`
	Label            string             `json:"label"`
	Id               ScreenId           `json:"id"`
}

/*
Enum of image types that can be disabled.
*/
type DisabledImageType string

type SetFocusEmulationEnabledArgs struct {
	Enabled bool `json:"enabled"`
}

type SetAutoDarkModeOverrideArgs struct {
	Enabled bool `json:"enabled,omitempty"`
}

type SetCPUThrottlingRateArgs struct {
	Rate float64 `json:"rate"`
}

type SetDefaultBackgroundColorOverrideArgs struct {
	Color *dom.RGBA `json:"color,omitempty"`
}

type SetSafeAreaInsetsOverrideArgs struct {
	Insets *SafeAreaInsets `json:"insets"`
}

type SetDeviceMetricsOverrideArgs struct {
	Width                          int                `json:"width"`
	Height                         int                `json:"height"`
	DeviceScaleFactor              float64            `json:"deviceScaleFactor"`
	Mobile                         bool               `json:"mobile"`
	Scale                          float64            `json:"scale,omitempty"`
	ScreenWidth                    int                `json:"screenWidth,omitempty"`
	ScreenHeight                   int                `json:"screenHeight,omitempty"`
	PositionX                      int                `json:"positionX,omitempty"`
	PositionY                      int                `json:"positionY,omitempty"`
	DontSetVisibleSize             bool               `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation              *ScreenOrientation `json:"screenOrientation,omitempty"`
	Viewport                       *page.Viewport     `json:"viewport,omitempty"`
	ScrollbarType                  string             `json:"scrollbarType,omitempty"`
	ScreenOrientationLockEmulation bool               `json:"screenOrientationLockEmulation,omitempty"`
}

type SetDevicePostureOverrideArgs struct {
	Posture *DevicePosture `json:"posture"`
}

type SetDisplayFeaturesOverrideArgs struct {
	Features []*DisplayFeature `json:"features"`
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

type SetEmulatedOSTextScaleArgs struct {
	Scale float64 `json:"scale,omitempty"`
}

type SetGeolocationOverrideArgs struct {
	Latitude         float64 `json:"latitude,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
	Accuracy         float64 `json:"accuracy,omitempty"`
	Altitude         float64 `json:"altitude,omitempty"`
	AltitudeAccuracy float64 `json:"altitudeAccuracy,omitempty"`
	Heading          float64 `json:"heading,omitempty"`
	Speed            float64 `json:"speed,omitempty"`
}

type GetOverriddenSensorInformationArgs struct {
	Type SensorType `json:"type"`
}

type GetOverriddenSensorInformationVal struct {
	RequestedSamplingFrequency float64 `json:"requestedSamplingFrequency"`
}

type SetSensorOverrideEnabledArgs struct {
	Enabled  bool            `json:"enabled"`
	Type     SensorType      `json:"type"`
	Metadata *SensorMetadata `json:"metadata,omitempty"`
}

type SetSensorOverrideReadingsArgs struct {
	Type    SensorType     `json:"type"`
	Reading *SensorReading `json:"reading"`
}

type SetPressureSourceOverrideEnabledArgs struct {
	Enabled  bool              `json:"enabled"`
	Source   PressureSource    `json:"source"`
	Metadata *PressureMetadata `json:"metadata,omitempty"`
}

type SetPressureStateOverrideArgs struct {
	Source PressureSource `json:"source"`
	State  PressureState  `json:"state"`
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

type SetDataSaverOverrideArgs struct {
	DataSaverEnabled bool `json:"dataSaverEnabled,omitempty"`
}

type SetHardwareConcurrencyOverrideArgs struct {
	HardwareConcurrency int `json:"hardwareConcurrency"`
}

type SetUserAgentOverrideArgs struct {
	UserAgent         string                    `json:"userAgent"`
	AcceptLanguage    string                    `json:"acceptLanguage,omitempty"`
	Platform          string                    `json:"platform,omitempty"`
	UserAgentMetadata *common.UserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

type SetAutomationOverrideArgs struct {
	Enabled bool `json:"enabled"`
}

type SetSmallViewportHeightDifferenceOverrideArgs struct {
	Difference int `json:"difference"`
}

type GetScreenInfosVal struct {
	ScreenInfos []*ScreenInfo `json:"screenInfos"`
}

type AddScreenArgs struct {
	Left             int             `json:"left"`
	Top              int             `json:"top"`
	Width            int             `json:"width"`
	Height           int             `json:"height"`
	WorkAreaInsets   *WorkAreaInsets `json:"workAreaInsets,omitempty"`
	DevicePixelRatio float64         `json:"devicePixelRatio,omitempty"`
	Rotation         int             `json:"rotation,omitempty"`
	ColorDepth       int             `json:"colorDepth,omitempty"`
	Label            string          `json:"label,omitempty"`
	IsInternal       bool            `json:"isInternal,omitempty"`
}

type AddScreenVal struct {
	ScreenInfo *ScreenInfo `json:"screenInfo"`
}

type UpdateScreenArgs struct {
	ScreenId         ScreenId        `json:"screenId"`
	Left             int             `json:"left,omitempty"`
	Top              int             `json:"top,omitempty"`
	Width            int             `json:"width,omitempty"`
	Height           int             `json:"height,omitempty"`
	WorkAreaInsets   *WorkAreaInsets `json:"workAreaInsets,omitempty"`
	DevicePixelRatio float64         `json:"devicePixelRatio,omitempty"`
	Rotation         int             `json:"rotation,omitempty"`
	ColorDepth       int             `json:"colorDepth,omitempty"`
	Label            string          `json:"label,omitempty"`
	IsInternal       bool            `json:"isInternal,omitempty"`
}

type UpdateScreenVal struct {
	ScreenInfo *ScreenInfo `json:"screenInfo"`
}

type RemoveScreenArgs struct {
	ScreenId ScreenId `json:"screenId"`
}

type SetPrimaryScreenArgs struct {
	ScreenId ScreenId `json:"screenId"`
}
