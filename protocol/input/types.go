package input

import (
	"github.com/ecwid/control/protocol/common"
)

/*
 */
type TouchPoint struct {
	X                  float64 `json:"x"`
	Y                  float64 `json:"y"`
	RadiusX            float64 `json:"radiusX,omitempty"`
	RadiusY            float64 `json:"radiusY,omitempty"`
	RotationAngle      float64 `json:"rotationAngle,omitempty"`
	Force              float64 `json:"force,omitempty"`
	TangentialPressure float64 `json:"tangentialPressure,omitempty"`
	TiltX              int     `json:"tiltX,omitempty"`
	TiltY              int     `json:"tiltY,omitempty"`
	Twist              int     `json:"twist,omitempty"`
	Id                 float64 `json:"id,omitempty"`
}

/*
 */
type GestureSourceType string

/*
 */
type MouseButton string

/*
UTC time in seconds, counted from January 1, 1970.
*/
type TimeSinceEpoch float64

/*
 */
type DragDataItem struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
	Title    string `json:"title,omitempty"`
	BaseURL  string `json:"baseURL,omitempty"`
}

/*
 */
type DragData struct {
	Items              []*DragDataItem `json:"items"`
	Files              []string        `json:"files,omitempty"`
	DragOperationsMask int             `json:"dragOperationsMask"`
}

type DispatchDragEventArgs struct {
	Type      string    `json:"type"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Data      *DragData `json:"data"`
	Modifiers int       `json:"modifiers,omitempty"`
}

type DispatchKeyEventArgs struct {
	Type                  string                `json:"type"`
	Modifiers             int                   `json:"modifiers,omitempty"`
	Timestamp             common.TimeSinceEpoch `json:"timestamp,omitempty"`
	Text                  string                `json:"text,omitempty"`
	UnmodifiedText        string                `json:"unmodifiedText,omitempty"`
	KeyIdentifier         string                `json:"keyIdentifier,omitempty"`
	Code                  string                `json:"code,omitempty"`
	Key                   string                `json:"key,omitempty"`
	WindowsVirtualKeyCode int                   `json:"windowsVirtualKeyCode,omitempty"`
	NativeVirtualKeyCode  int                   `json:"nativeVirtualKeyCode,omitempty"`
	AutoRepeat            bool                  `json:"autoRepeat,omitempty"`
	IsKeypad              bool                  `json:"isKeypad,omitempty"`
	IsSystemKey           bool                  `json:"isSystemKey,omitempty"`
	Location              int                   `json:"location,omitempty"`
	Commands              []string              `json:"commands,omitempty"`
}

type InsertTextArgs struct {
	Text string `json:"text"`
}

type ImeSetCompositionArgs struct {
	Text             string `json:"text"`
	SelectionStart   int    `json:"selectionStart"`
	SelectionEnd     int    `json:"selectionEnd"`
	ReplacementStart int    `json:"replacementStart,omitempty"`
	ReplacementEnd   int    `json:"replacementEnd,omitempty"`
}

type DispatchMouseEventArgs struct {
	Type               string                `json:"type"`
	X                  float64               `json:"x"`
	Y                  float64               `json:"y"`
	Modifiers          int                   `json:"modifiers,omitempty"`
	Timestamp          common.TimeSinceEpoch `json:"timestamp,omitempty"`
	Button             MouseButton           `json:"button,omitempty"`
	Buttons            int                   `json:"buttons,omitempty"`
	ClickCount         int                   `json:"clickCount,omitempty"`
	Force              float64               `json:"force,omitempty"`
	TangentialPressure float64               `json:"tangentialPressure,omitempty"`
	TiltX              int                   `json:"tiltX,omitempty"`
	TiltY              int                   `json:"tiltY,omitempty"`
	Twist              int                   `json:"twist,omitempty"`
	DeltaX             float64               `json:"deltaX,omitempty"`
	DeltaY             float64               `json:"deltaY,omitempty"`
	PointerType        string                `json:"pointerType,omitempty"`
}

type DispatchTouchEventArgs struct {
	Type        string                `json:"type"`
	TouchPoints []*TouchPoint         `json:"touchPoints"`
	Modifiers   int                   `json:"modifiers,omitempty"`
	Timestamp   common.TimeSinceEpoch `json:"timestamp,omitempty"`
}

type EmulateTouchFromMouseEventArgs struct {
	Type       string                `json:"type"`
	X          int                   `json:"x"`
	Y          int                   `json:"y"`
	Button     MouseButton           `json:"button"`
	Timestamp  common.TimeSinceEpoch `json:"timestamp,omitempty"`
	DeltaX     float64               `json:"deltaX,omitempty"`
	DeltaY     float64               `json:"deltaY,omitempty"`
	Modifiers  int                   `json:"modifiers,omitempty"`
	ClickCount int                   `json:"clickCount,omitempty"`
}

type SetIgnoreInputEventsArgs struct {
	Ignore bool `json:"ignore"`
}

type SetInterceptDragsArgs struct {
	Enabled bool `json:"enabled"`
}

type SynthesizePinchGestureArgs struct {
	X                 float64           `json:"x"`
	Y                 float64           `json:"y"`
	ScaleFactor       float64           `json:"scaleFactor"`
	RelativeSpeed     int               `json:"relativeSpeed,omitempty"`
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"`
}

type SynthesizeScrollGestureArgs struct {
	X                     float64           `json:"x"`
	Y                     float64           `json:"y"`
	XDistance             float64           `json:"xDistance,omitempty"`
	YDistance             float64           `json:"yDistance,omitempty"`
	XOverscroll           float64           `json:"xOverscroll,omitempty"`
	YOverscroll           float64           `json:"yOverscroll,omitempty"`
	PreventFling          bool              `json:"preventFling,omitempty"`
	Speed                 int               `json:"speed,omitempty"`
	GestureSourceType     GestureSourceType `json:"gestureSourceType,omitempty"`
	RepeatCount           int               `json:"repeatCount,omitempty"`
	RepeatDelayMs         int               `json:"repeatDelayMs,omitempty"`
	InteractionMarkerName string            `json:"interactionMarkerName,omitempty"`
}

type SynthesizeTapGestureArgs struct {
	X                 float64           `json:"x"`
	Y                 float64           `json:"y"`
	Duration          int               `json:"duration,omitempty"`
	TapCount          int               `json:"tapCount,omitempty"`
	GestureSourceType GestureSourceType `json:"gestureSourceType,omitempty"`
}
