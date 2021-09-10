package browser

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/target"
)

/*

 */
type BrowserContextID string

/*

 */
type WindowID int

/*
	The state of the browser window.
*/
type WindowState string

/*
	Browser window bounds information
*/
type Bounds struct {
	Left        int         `json:"left,omitempty"`
	Top         int         `json:"top,omitempty"`
	Width       int         `json:"width,omitempty"`
	Height      int         `json:"height,omitempty"`
	WindowState WindowState `json:"windowState,omitempty"`
}

/*

 */
type PermissionType string

/*

 */
type PermissionSetting string

/*
	Definition of PermissionDescriptor defined in the Permissions API:
https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
*/
type PermissionDescriptor struct {
	Name                     string `json:"name"`
	Sysex                    bool   `json:"sysex,omitempty"`
	UserVisibleOnly          bool   `json:"userVisibleOnly,omitempty"`
	AllowWithoutSanitization bool   `json:"allowWithoutSanitization,omitempty"`
	PanTiltZoom              bool   `json:"panTiltZoom,omitempty"`
}

/*
	Browser command ids used by executeBrowserCommand.
*/
type BrowserCommandId string

/*
	Chrome histogram bucket.
*/
type Bucket struct {
	Low   int `json:"low"`
	High  int `json:"high"`
	Count int `json:"count"`
}

/*
	Chrome histogram.
*/
type Histogram struct {
	Name    string    `json:"name"`
	Sum     int       `json:"sum"`
	Count   int       `json:"count"`
	Buckets []*Bucket `json:"buckets"`
}

type SetPermissionArgs struct {
	Permission       *PermissionDescriptor   `json:"permission"`
	Setting          PermissionSetting       `json:"setting"`
	Origin           string                  `json:"origin,omitempty"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type GrantPermissionsArgs struct {
	Permissions      []PermissionType        `json:"permissions"`
	Origin           string                  `json:"origin,omitempty"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type ResetPermissionsArgs struct {
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type SetDownloadBehaviorArgs struct {
	Behavior         string                  `json:"behavior"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
	DownloadPath     string                  `json:"downloadPath,omitempty"`
	EventsEnabled    bool                    `json:"eventsEnabled,omitempty"`
}

type CancelDownloadArgs struct {
	Guid             string                  `json:"guid"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type GetVersionVal struct {
	ProtocolVersion string `json:"protocolVersion"`
	Product         string `json:"product"`
	Revision        string `json:"revision"`
	UserAgent       string `json:"userAgent"`
	JsVersion       string `json:"jsVersion"`
}

type GetBrowserCommandLineVal struct {
	Arguments []string `json:"arguments"`
}

type GetHistogramsArgs struct {
	Query string `json:"query,omitempty"`
	Delta bool   `json:"delta,omitempty"`
}

type GetHistogramsVal struct {
	Histograms []*Histogram `json:"histograms"`
}

type GetHistogramArgs struct {
	Name  string `json:"name"`
	Delta bool   `json:"delta,omitempty"`
}

type GetHistogramVal struct {
	Histogram *Histogram `json:"histogram"`
}

type GetWindowBoundsArgs struct {
	WindowId WindowID `json:"windowId"`
}

type GetWindowBoundsVal struct {
	Bounds *Bounds `json:"bounds"`
}

type GetWindowForTargetArgs struct {
	TargetId target.TargetID `json:"targetId,omitempty"`
}

type GetWindowForTargetVal struct {
	WindowId WindowID `json:"windowId"`
	Bounds   *Bounds  `json:"bounds"`
}

type SetWindowBoundsArgs struct {
	WindowId WindowID `json:"windowId"`
	Bounds   *Bounds  `json:"bounds"`
}

type SetDockTileArgs struct {
	BadgeLabel string `json:"badgeLabel,omitempty"`
	Image      []byte `json:"image,omitempty"`
}

type ExecuteBrowserCommandArgs struct {
	CommandId BrowserCommandId `json:"commandId"`
}
