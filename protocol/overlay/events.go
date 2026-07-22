package overlay

import (
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/page"
)

/*
	Fired when the node should be inspected. This happens after call to `setInspectMode` or when

user manually inspects an element.
*/
type InspectNodeRequested struct {
	BackendNodeId dom.BackendNodeId `json:"backendNodeId"`
}

/*
Fired when the node should be highlighted. This happens after call to `setInspectMode`.
*/
type NodeHighlightRequested struct {
	NodeId dom.NodeId `json:"nodeId"`
}

/*
Fired when user asks to capture screenshot of some area on the page.
*/
type ScreenshotRequested struct {
	Viewport *page.Viewport `json:"viewport"`
}

/*
Fired when user asks to show the Inspect panel.
*/
type InspectPanelShowRequested struct {
	BackendNodeId dom.BackendNodeId `json:"backendNodeId"`
}

/*
Fired when user asks to restore the Inspected Element floating window.
*/
type InspectedElementWindowRestored struct {
	BackendNodeId dom.BackendNodeId `json:"backendNodeId"`
}

/*
Fired when user cancels the inspect mode.
*/
type InspectModeCanceled any
