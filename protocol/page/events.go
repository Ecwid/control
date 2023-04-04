package page

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/runtime"
)

/*
 */
type DomContentEventFired struct {
	Timestamp network.MonotonicTime `json:"timestamp"`
}

/*
Emitted only when `page.interceptFileChooser` is enabled.
*/
type FileChooserOpened struct {
	FrameId       common.FrameId    `json:"frameId"`
	Mode          string            `json:"mode"`
	BackendNodeId dom.BackendNodeId `json:"backendNodeId,omitempty"`
}

/*
Fired when frame has been attached to its parent.
*/
type FrameAttached struct {
	FrameId       common.FrameId      `json:"frameId"`
	ParentFrameId common.FrameId      `json:"parentFrameId"`
	Stack         *runtime.StackTrace `json:"stack,omitempty"`
}

/*
Fired when frame has been detached from its parent.
*/
type FrameDetached struct {
	FrameId common.FrameId `json:"frameId"`
	Reason  string         `json:"reason"`
}

/*
Fired once navigation of the frame has completed. Frame is now associated with the new loader.
*/
type FrameNavigated struct {
	Frame *Frame         `json:"frame"`
	Type  NavigationType `json:"type"`
}

/*
Fired when opening document to write to.
*/
type DocumentOpened struct {
	Frame *Frame `json:"frame"`
}

/*
 */
type FrameResized interface{}

/*
	Fired when a renderer-initiated navigation is requested.

Navigation may still be cancelled after the event is issued.
*/
type FrameRequestedNavigation struct {
	FrameId     common.FrameId              `json:"frameId"`
	Reason      ClientNavigationReason      `json:"reason"`
	Url         string                      `json:"url"`
	Disposition ClientNavigationDisposition `json:"disposition"`
}

/*
Fired when frame has started loading.
*/
type FrameStartedLoading struct {
	FrameId common.FrameId `json:"frameId"`
}

/*
Fired when frame has stopped loading.
*/
type FrameStoppedLoading struct {
	FrameId common.FrameId `json:"frameId"`
}

/*
Fired when interstitial page was hidden
*/
type InterstitialHidden interface{}

/*
Fired when interstitial page was shown
*/
type InterstitialShown interface{}

/*
	Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been

closed.
*/
type JavascriptDialogClosed struct {
	Result    bool   `json:"result"`
	UserInput string `json:"userInput"`
}

/*
	Fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to

open.
*/
type JavascriptDialogOpening struct {
	Url               string     `json:"url"`
	Message           string     `json:"message"`
	Type              DialogType `json:"type"`
	HasBrowserHandler bool       `json:"hasBrowserHandler"`
	DefaultPrompt     string     `json:"defaultPrompt,omitempty"`
}

/*
Fired for top level page lifecycle events such as navigation, load, paint, etc.
*/
type LifecycleEvent struct {
	FrameId   common.FrameId        `json:"frameId"`
	LoaderId  network.LoaderId      `json:"loaderId"`
	Name      string                `json:"name"`
	Timestamp network.MonotonicTime `json:"timestamp"`
}

/*
	Fired for failed bfcache history navigations if BackForwardCache feature is enabled. Do

not assume any ordering with the Page.frameNavigated event. This event is fired only for
main-frame history navigation where the document changes (non-same-document navigations),
when bfcache navigation fails.
*/
type BackForwardCacheNotUsed struct {
	LoaderId                    network.LoaderId                            `json:"loaderId"`
	FrameId                     common.FrameId                              `json:"frameId"`
	NotRestoredExplanations     []*BackForwardCacheNotRestoredExplanation   `json:"notRestoredExplanations"`
	NotRestoredExplanationsTree *BackForwardCacheNotRestoredExplanationTree `json:"notRestoredExplanationsTree,omitempty"`
}

/*
Fired when a prerender attempt is completed.
*/
type PrerenderAttemptCompleted struct {
	InitiatingFrameId   common.FrameId       `json:"initiatingFrameId"`
	PrerenderingUrl     string               `json:"prerenderingUrl"`
	FinalStatus         PrerenderFinalStatus `json:"finalStatus"`
	DisallowedApiMethod string               `json:"disallowedApiMethod,omitempty"`
}

/*
 */
type LoadEventFired struct {
	Timestamp network.MonotonicTime `json:"timestamp"`
}

/*
Fired when same-document navigation happens, e.g. due to history API usage or anchor navigation.
*/
type NavigatedWithinDocument struct {
	FrameId common.FrameId `json:"frameId"`
	Url     string         `json:"url"`
}

/*
Compressed image data requested by the `startScreencast`.
*/
type ScreencastFrame struct {
	Data      []byte                   `json:"data"`
	Metadata  *ScreencastFrameMetadata `json:"metadata"`
	SessionId int                      `json:"sessionId"`
}

/*
Fired when the page with currently enabled screencast was shown or hidden `.
*/
type ScreencastVisibilityChanged struct {
	Visible bool `json:"visible"`
}

/*
	Fired when a new window is going to be opened, via window.open(), link click, form submission,

etc.
*/
type WindowOpen struct {
	Url            string   `json:"url"`
	WindowName     string   `json:"windowName"`
	WindowFeatures []string `json:"windowFeatures"`
	UserGesture    bool     `json:"userGesture"`
}

/*
	Issued for every compilation cache generated. Is only available

if Page.setGenerateCompilationCache is enabled.
*/
type CompilationCacheProduced struct {
	Url  string `json:"url"`
	Data []byte `json:"data"`
}
