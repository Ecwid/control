package devtool

// OrientationType https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type OrientationType string

// https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
const (
	PortraitPrimary    OrientationType = "portraitPrimary"
	PortraitSecondary                  = "portraitSecondary"
	LandscapePrimary                   = "landscapePrimary"
	LandscapeSecondary                 = "landscapeSecondary"
)

// portraitPrimary, portraitSecondary, landscapePrimary, landscapeSecondary

// ScreenOrientation https://chromedevtools.github.io/devtools-protocol/tot/Emulation#type-ScreenOrientation
type ScreenOrientation struct {
	Type  OrientationType `json:"type"`
	Angle int64           `json:"angle"`
}

// DeviceMetrics https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setDeviceMetricsOverride
type DeviceMetrics struct {
	Width              int64              `json:"width"`
	Height             int64              `json:"height"`
	DeviceScaleFactor  float64            `json:"deviceScaleFactor"`
	Mobile             bool               `json:"mobile"`
	Scale              float64            `json:"scale,omitempty"`
	ScreenWidth        int64              `json:"screenWidth,omitempty"`
	ScreenHeight       int64              `json:"screenHeight,omitempty"`
	PositionX          int64              `json:"positionX,omitempty"`
	PositionY          int64              `json:"positionY,omitempty"`
	DontSetVisibleSize *bool              `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation  *ScreenOrientation `json:"screenOrientation,omitempty"`
	Viewport           *Viewport          `json:"viewport,omitempty"`
}
