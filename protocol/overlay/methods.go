package overlay

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables domain notifications.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Overlay.disable", nil, nil)
}

/*
Enables domain notifications.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Overlay.enable", nil, nil)
}

/*
For testing.
*/
func GetHighlightObjectForTest(c protocol.Caller, args GetHighlightObjectForTestArgs) (*GetHighlightObjectForTestVal, error) {
	var val = &GetHighlightObjectForTestVal{}
	return val, c.Call("Overlay.getHighlightObjectForTest", args, val)
}

/*
For Persistent Grid testing.
*/
func GetGridHighlightObjectsForTest(c protocol.Caller, args GetGridHighlightObjectsForTestArgs) (*GetGridHighlightObjectsForTestVal, error) {
	var val = &GetGridHighlightObjectsForTestVal{}
	return val, c.Call("Overlay.getGridHighlightObjectsForTest", args, val)
}

/*
For Source Order Viewer testing.
*/
func GetSourceOrderHighlightObjectForTest(c protocol.Caller, args GetSourceOrderHighlightObjectForTestArgs) (*GetSourceOrderHighlightObjectForTestVal, error) {
	var val = &GetSourceOrderHighlightObjectForTestVal{}
	return val, c.Call("Overlay.getSourceOrderHighlightObjectForTest", args, val)
}

/*
Hides any highlight.
*/
func HideHighlight(c protocol.Caller) error {
	return c.Call("Overlay.hideHighlight", nil, nil)
}

/*
	Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or

objectId must be specified.
*/
func HighlightNode(c protocol.Caller, args HighlightNodeArgs) error {
	return c.Call("Overlay.highlightNode", args, nil)
}

/*
Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
*/
func HighlightQuad(c protocol.Caller, args HighlightQuadArgs) error {
	return c.Call("Overlay.highlightQuad", args, nil)
}

/*
Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
*/
func HighlightRect(c protocol.Caller, args HighlightRectArgs) error {
	return c.Call("Overlay.highlightRect", args, nil)
}

/*
	Highlights the source order of the children of the DOM node with given id or with the given

JavaScript object wrapper. Either nodeId or objectId must be specified.
*/
func HighlightSourceOrder(c protocol.Caller, args HighlightSourceOrderArgs) error {
	return c.Call("Overlay.highlightSourceOrder", args, nil)
}

/*
	Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted.

Backend then generates 'inspectNodeRequested' event upon element selection.
*/
func SetInspectMode(c protocol.Caller, args SetInspectModeArgs) error {
	return c.Call("Overlay.setInspectMode", args, nil)
}

/*
Highlights owner element of all frames detected to be ads.
*/
func SetShowAdHighlights(c protocol.Caller, args SetShowAdHighlightsArgs) error {
	return c.Call("Overlay.setShowAdHighlights", args, nil)
}

/*
 */
func SetPausedInDebuggerMessage(c protocol.Caller, args SetPausedInDebuggerMessageArgs) error {
	return c.Call("Overlay.setPausedInDebuggerMessage", args, nil)
}

/*
Requests that backend shows debug borders on layers
*/
func SetShowDebugBorders(c protocol.Caller, args SetShowDebugBordersArgs) error {
	return c.Call("Overlay.setShowDebugBorders", args, nil)
}

/*
Requests that backend shows the FPS counter
*/
func SetShowFPSCounter(c protocol.Caller, args SetShowFPSCounterArgs) error {
	return c.Call("Overlay.setShowFPSCounter", args, nil)
}

/*
Highlight multiple elements with the CSS Grid overlay.
*/
func SetShowGridOverlays(c protocol.Caller, args SetShowGridOverlaysArgs) error {
	return c.Call("Overlay.setShowGridOverlays", args, nil)
}

/*
 */
func SetShowFlexOverlays(c protocol.Caller, args SetShowFlexOverlaysArgs) error {
	return c.Call("Overlay.setShowFlexOverlays", args, nil)
}

/*
 */
func SetShowScrollSnapOverlays(c protocol.Caller, args SetShowScrollSnapOverlaysArgs) error {
	return c.Call("Overlay.setShowScrollSnapOverlays", args, nil)
}

/*
 */
func SetShowContainerQueryOverlays(c protocol.Caller, args SetShowContainerQueryOverlaysArgs) error {
	return c.Call("Overlay.setShowContainerQueryOverlays", args, nil)
}

/*
Requests that backend shows paint rectangles
*/
func SetShowPaintRects(c protocol.Caller, args SetShowPaintRectsArgs) error {
	return c.Call("Overlay.setShowPaintRects", args, nil)
}

/*
Requests that backend shows layout shift regions
*/
func SetShowLayoutShiftRegions(c protocol.Caller, args SetShowLayoutShiftRegionsArgs) error {
	return c.Call("Overlay.setShowLayoutShiftRegions", args, nil)
}

/*
Requests that backend shows scroll bottleneck rects
*/
func SetShowScrollBottleneckRects(c protocol.Caller, args SetShowScrollBottleneckRectsArgs) error {
	return c.Call("Overlay.setShowScrollBottleneckRects", args, nil)
}

/*
Request that backend shows an overlay with web vital metrics.
*/
func SetShowWebVitals(c protocol.Caller, args SetShowWebVitalsArgs) error {
	return c.Call("Overlay.setShowWebVitals", args, nil)
}

/*
Paints viewport size upon main frame resize.
*/
func SetShowViewportSizeOnResize(c protocol.Caller, args SetShowViewportSizeOnResizeArgs) error {
	return c.Call("Overlay.setShowViewportSizeOnResize", args, nil)
}

/*
Add a dual screen device hinge
*/
func SetShowHinge(c protocol.Caller, args SetShowHingeArgs) error {
	return c.Call("Overlay.setShowHinge", args, nil)
}

/*
Show elements in isolation mode with overlays.
*/
func SetShowIsolatedElements(c protocol.Caller, args SetShowIsolatedElementsArgs) error {
	return c.Call("Overlay.setShowIsolatedElements", args, nil)
}
