package target

import (
	"github.com/ecwid/control/protocol/common"
)

/*
 */
type TargetID string

/*
Unique identifier of attached debugging session.
*/
type SessionID string

/*
 */
type TargetInfo struct {
	TargetId         TargetID                `json:"targetId"`
	Type             string                  `json:"type"`
	Title            string                  `json:"title"`
	Url              string                  `json:"url"`
	Attached         bool                    `json:"attached"`
	OpenerId         TargetID                `json:"openerId,omitempty"`
	CanAccessOpener  bool                    `json:"canAccessOpener"`
	OpenerFrameId    common.FrameId          `json:"openerFrameId,omitempty"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
	Subtype          string                  `json:"subtype,omitempty"`
}

/*
A filter used by target query/discovery/auto-attach operations.
*/
type FilterEntry struct {
	Exclude bool   `json:"exclude,omitempty"`
	Type    string `json:"type,omitempty"`
}

/*
	The entries in TargetFilter are matched sequentially against targets and

the first entry that matches determines if the target is included or not,
depending on the value of `exclude` field in the entry.
If filter is not specified, the one assumed is
[{type: "browser", exclude: true}, {type: "tab", exclude: true}, {}]
(i.e. include everything but `browser` and `tab`).
*/
type TargetFilter []*FilterEntry

/*
 */
type RemoteLocation struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ActivateTargetArgs struct {
	TargetId TargetID `json:"targetId"`
}

type AttachToTargetArgs struct {
	TargetId TargetID `json:"targetId"`
	Flatten  bool     `json:"flatten,omitempty"`
}

type AttachToTargetVal struct {
	SessionId SessionID `json:"sessionId"`
}

type AttachToBrowserTargetVal struct {
	SessionId SessionID `json:"sessionId"`
}

type CloseTargetArgs struct {
	TargetId TargetID `json:"targetId"`
}

type ExposeDevToolsProtocolArgs struct {
	TargetId    TargetID `json:"targetId"`
	BindingName string   `json:"bindingName,omitempty"`
}

type CreateBrowserContextArgs struct {
	DisposeOnDetach                   bool     `json:"disposeOnDetach,omitempty"`
	ProxyServer                       string   `json:"proxyServer,omitempty"`
	ProxyBypassList                   string   `json:"proxyBypassList,omitempty"`
	OriginsWithUniversalNetworkAccess []string `json:"originsWithUniversalNetworkAccess,omitempty"`
}

type CreateBrowserContextVal struct {
	BrowserContextId common.BrowserContextID `json:"browserContextId"`
}

type GetBrowserContextsVal struct {
	BrowserContextIds []common.BrowserContextID `json:"browserContextIds"`
}

type CreateTargetArgs struct {
	Url                     string                  `json:"url"`
	Width                   int                     `json:"width,omitempty"`
	Height                  int                     `json:"height,omitempty"`
	BrowserContextId        common.BrowserContextID `json:"browserContextId,omitempty"`
	EnableBeginFrameControl bool                    `json:"enableBeginFrameControl,omitempty"`
	NewWindow               bool                    `json:"newWindow,omitempty"`
	Background              bool                    `json:"background,omitempty"`
	ForTab                  bool                    `json:"forTab,omitempty"`
}

type CreateTargetVal struct {
	TargetId TargetID `json:"targetId"`
}

type DetachFromTargetArgs struct {
	SessionId SessionID `json:"sessionId,omitempty"`
}

type DisposeBrowserContextArgs struct {
	BrowserContextId common.BrowserContextID `json:"browserContextId"`
}

type GetTargetInfoArgs struct {
	TargetId TargetID `json:"targetId,omitempty"`
}

type GetTargetInfoVal struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}

type GetTargetsArgs struct {
	Filter TargetFilter `json:"filter,omitempty"`
}

type GetTargetsVal struct {
	TargetInfos []*TargetInfo `json:"targetInfos"`
}

type SetAutoAttachArgs struct {
	AutoAttach             bool         `json:"autoAttach"`
	WaitForDebuggerOnStart bool         `json:"waitForDebuggerOnStart"`
	Flatten                bool         `json:"flatten,omitempty"`
	Filter                 TargetFilter `json:"filter,omitempty"`
}

type AutoAttachRelatedArgs struct {
	TargetId               TargetID     `json:"targetId"`
	WaitForDebuggerOnStart bool         `json:"waitForDebuggerOnStart"`
	Filter                 TargetFilter `json:"filter,omitempty"`
}

type SetDiscoverTargetsArgs struct {
	Discover bool         `json:"discover"`
	Filter   TargetFilter `json:"filter,omitempty"`
}

type SetRemoteLocationsArgs struct {
	Locations []*RemoteLocation `json:"locations"`
}
