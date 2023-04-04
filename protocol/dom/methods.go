package dom

import (
	"github.com/ecwid/control/protocol"
)

/*
Collects class names for the node with given id and all of it's child nodes.
*/
func CollectClassNamesFromSubtree(c protocol.Caller, args CollectClassNamesFromSubtreeArgs) (*CollectClassNamesFromSubtreeVal, error) {
	var val = &CollectClassNamesFromSubtreeVal{}
	return val, c.Call("DOM.collectClassNamesFromSubtree", args, val)
}

/*
	Creates a deep copy of the specified node and places it into the target container before the

given anchor.
*/
func CopyTo(c protocol.Caller, args CopyToArgs) (*CopyToVal, error) {
	var val = &CopyToVal{}
	return val, c.Call("DOM.copyTo", args, val)
}

/*
	Describes node given its id, does not require domain to be enabled. Does not start tracking any

objects, can be used for automation.
*/
func DescribeNode(c protocol.Caller, args DescribeNodeArgs) (*DescribeNodeVal, error) {
	var val = &DescribeNodeVal{}
	return val, c.Call("DOM.describeNode", args, val)
}

/*
	Scrolls the specified rect of the given node into view if not already visible.

Note: exactly one between nodeId, backendNodeId and objectId should be passed
to identify the node.
*/
func ScrollIntoViewIfNeeded(c protocol.Caller, args ScrollIntoViewIfNeededArgs) error {
	return c.Call("DOM.scrollIntoViewIfNeeded", args, nil)
}

/*
Disables DOM agent for the given page.
*/
func Disable(c protocol.Caller) error {
	return c.Call("DOM.disable", nil, nil)
}

/*
	Discards search results from the session with the given id. `getSearchResults` should no longer

be called for that search.
*/
func DiscardSearchResults(c protocol.Caller, args DiscardSearchResultsArgs) error {
	return c.Call("DOM.discardSearchResults", args, nil)
}

/*
Enables DOM agent for the given page.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("DOM.enable", args, nil)
}

/*
Focuses the given element.
*/
func Focus(c protocol.Caller, args FocusArgs) error {
	return c.Call("DOM.focus", args, nil)
}

/*
Returns attributes for the specified node.
*/
func GetAttributes(c protocol.Caller, args GetAttributesArgs) (*GetAttributesVal, error) {
	var val = &GetAttributesVal{}
	return val, c.Call("DOM.getAttributes", args, val)
}

/*
Returns boxes for the given node.
*/
func GetBoxModel(c protocol.Caller, args GetBoxModelArgs) (*GetBoxModelVal, error) {
	var val = &GetBoxModelVal{}
	return val, c.Call("DOM.getBoxModel", args, val)
}

/*
	Returns quads that describe node position on the page. This method

might return multiple quads for inline nodes.
*/
func GetContentQuads(c protocol.Caller, args GetContentQuadsArgs) (*GetContentQuadsVal, error) {
	var val = &GetContentQuadsVal{}
	return val, c.Call("DOM.getContentQuads", args, val)
}

/*
Returns the root DOM node (and optionally the subtree) to the caller.
*/
func GetDocument(c protocol.Caller, args GetDocumentArgs) (*GetDocumentVal, error) {
	var val = &GetDocumentVal{}
	return val, c.Call("DOM.getDocument", args, val)
}

/*
Finds nodes with a given computed style in a subtree.
*/
func GetNodesForSubtreeByStyle(c protocol.Caller, args GetNodesForSubtreeByStyleArgs) (*GetNodesForSubtreeByStyleVal, error) {
	var val = &GetNodesForSubtreeByStyleVal{}
	return val, c.Call("DOM.getNodesForSubtreeByStyle", args, val)
}

/*
	Returns node id at given location. Depending on whether DOM domain is enabled, nodeId is

either returned or not.
*/
func GetNodeForLocation(c protocol.Caller, args GetNodeForLocationArgs) (*GetNodeForLocationVal, error) {
	var val = &GetNodeForLocationVal{}
	return val, c.Call("DOM.getNodeForLocation", args, val)
}

/*
Returns node's HTML markup.
*/
func GetOuterHTML(c protocol.Caller, args GetOuterHTMLArgs) (*GetOuterHTMLVal, error) {
	var val = &GetOuterHTMLVal{}
	return val, c.Call("DOM.getOuterHTML", args, val)
}

/*
Returns the id of the nearest ancestor that is a relayout boundary.
*/
func GetRelayoutBoundary(c protocol.Caller, args GetRelayoutBoundaryArgs) (*GetRelayoutBoundaryVal, error) {
	var val = &GetRelayoutBoundaryVal{}
	return val, c.Call("DOM.getRelayoutBoundary", args, val)
}

/*
	Returns search results from given `fromIndex` to given `toIndex` from the search with the given

identifier.
*/
func GetSearchResults(c protocol.Caller, args GetSearchResultsArgs) (*GetSearchResultsVal, error) {
	var val = &GetSearchResultsVal{}
	return val, c.Call("DOM.getSearchResults", args, val)
}

/*
Hides any highlight.
*/
func HideHighlight(c protocol.Caller) error {
	return c.Call("DOM.hideHighlight", nil, nil)
}

/*
Highlights DOM node.
*/
func HighlightNode(c protocol.Caller) error {
	return c.Call("DOM.highlightNode", nil, nil)
}

/*
Highlights given rectangle.
*/
func HighlightRect(c protocol.Caller) error {
	return c.Call("DOM.highlightRect", nil, nil)
}

/*
Marks last undoable state.
*/
func MarkUndoableState(c protocol.Caller) error {
	return c.Call("DOM.markUndoableState", nil, nil)
}

/*
Moves node into the new container, places it before the given anchor.
*/
func MoveTo(c protocol.Caller, args MoveToArgs) (*MoveToVal, error) {
	var val = &MoveToVal{}
	return val, c.Call("DOM.moveTo", args, val)
}

/*
	Searches for a given string in the DOM tree. Use `getSearchResults` to access search results or

`cancelSearch` to end this search session.
*/
func PerformSearch(c protocol.Caller, args PerformSearchArgs) (*PerformSearchVal, error) {
	var val = &PerformSearchVal{}
	return val, c.Call("DOM.performSearch", args, val)
}

/*
Requests that the node is sent to the caller given its path. // FIXME, use XPath
*/
func PushNodeByPathToFrontend(c protocol.Caller, args PushNodeByPathToFrontendArgs) (*PushNodeByPathToFrontendVal, error) {
	var val = &PushNodeByPathToFrontendVal{}
	return val, c.Call("DOM.pushNodeByPathToFrontend", args, val)
}

/*
Requests that a batch of nodes is sent to the caller given their backend node ids.
*/
func PushNodesByBackendIdsToFrontend(c protocol.Caller, args PushNodesByBackendIdsToFrontendArgs) (*PushNodesByBackendIdsToFrontendVal, error) {
	var val = &PushNodesByBackendIdsToFrontendVal{}
	return val, c.Call("DOM.pushNodesByBackendIdsToFrontend", args, val)
}

/*
Executes `querySelector` on a given node.
*/
func QuerySelector(c protocol.Caller, args QuerySelectorArgs) (*QuerySelectorVal, error) {
	var val = &QuerySelectorVal{}
	return val, c.Call("DOM.querySelector", args, val)
}

/*
Executes `querySelectorAll` on a given node.
*/
func QuerySelectorAll(c protocol.Caller, args QuerySelectorAllArgs) (*QuerySelectorAllVal, error) {
	var val = &QuerySelectorAllVal{}
	return val, c.Call("DOM.querySelectorAll", args, val)
}

/*
	Returns NodeIds of current top layer elements.

Top layer is rendered closest to the user within a viewport, therefore its elements always
appear on top of all other content.
*/
func GetTopLayerElements(c protocol.Caller) (*GetTopLayerElementsVal, error) {
	var val = &GetTopLayerElementsVal{}
	return val, c.Call("DOM.getTopLayerElements", nil, val)
}

/*
Re-does the last undone action.
*/
func Redo(c protocol.Caller) error {
	return c.Call("DOM.redo", nil, nil)
}

/*
Removes attribute with given name from an element with given id.
*/
func RemoveAttribute(c protocol.Caller, args RemoveAttributeArgs) error {
	return c.Call("DOM.removeAttribute", args, nil)
}

/*
Removes node with given id.
*/
func RemoveNode(c protocol.Caller, args RemoveNodeArgs) error {
	return c.Call("DOM.removeNode", args, nil)
}

/*
	Requests that children of the node with given id are returned to the caller in form of

`setChildNodes` events where not only immediate children are retrieved, but all children down to
the specified depth.
*/
func RequestChildNodes(c protocol.Caller, args RequestChildNodesArgs) error {
	return c.Call("DOM.requestChildNodes", args, nil)
}

/*
	Requests that the node is sent to the caller given the JavaScript node object reference. All

nodes that form the path from the node to the root are also sent to the client as a series of
`setChildNodes` notifications.
*/
func RequestNode(c protocol.Caller, args RequestNodeArgs) (*RequestNodeVal, error) {
	var val = &RequestNodeVal{}
	return val, c.Call("DOM.requestNode", args, val)
}

/*
Resolves the JavaScript node object for a given NodeId or BackendNodeId.
*/
func ResolveNode(c protocol.Caller, args ResolveNodeArgs) (*ResolveNodeVal, error) {
	var val = &ResolveNodeVal{}
	return val, c.Call("DOM.resolveNode", args, val)
}

/*
Sets attribute for an element with given id.
*/
func SetAttributeValue(c protocol.Caller, args SetAttributeValueArgs) error {
	return c.Call("DOM.setAttributeValue", args, nil)
}

/*
	Sets attributes on element with given id. This method is useful when user edits some existing

attribute value and types in several attribute name/value pairs.
*/
func SetAttributesAsText(c protocol.Caller, args SetAttributesAsTextArgs) error {
	return c.Call("DOM.setAttributesAsText", args, nil)
}

/*
Sets files for the given file input element.
*/
func SetFileInputFiles(c protocol.Caller, args SetFileInputFilesArgs) error {
	return c.Call("DOM.setFileInputFiles", args, nil)
}

/*
Sets if stack traces should be captured for Nodes. See `Node.getNodeStackTraces`. Default is disabled.
*/
func SetNodeStackTracesEnabled(c protocol.Caller, args SetNodeStackTracesEnabledArgs) error {
	return c.Call("DOM.setNodeStackTracesEnabled", args, nil)
}

/*
Gets stack traces associated with a Node. As of now, only provides stack trace for Node creation.
*/
func GetNodeStackTraces(c protocol.Caller, args GetNodeStackTracesArgs) (*GetNodeStackTracesVal, error) {
	var val = &GetNodeStackTracesVal{}
	return val, c.Call("DOM.getNodeStackTraces", args, val)
}

/*
	Returns file information for the given

File wrapper.
*/
func GetFileInfo(c protocol.Caller, args GetFileInfoArgs) (*GetFileInfoVal, error) {
	var val = &GetFileInfoVal{}
	return val, c.Call("DOM.getFileInfo", args, val)
}

/*
	Enables console to refer to the node with given id via $x (see Command Line API for more details

$x functions).
*/
func SetInspectedNode(c protocol.Caller, args SetInspectedNodeArgs) error {
	return c.Call("DOM.setInspectedNode", args, nil)
}

/*
Sets node name for a node with given id.
*/
func SetNodeName(c protocol.Caller, args SetNodeNameArgs) (*SetNodeNameVal, error) {
	var val = &SetNodeNameVal{}
	return val, c.Call("DOM.setNodeName", args, val)
}

/*
Sets node value for a node with given id.
*/
func SetNodeValue(c protocol.Caller, args SetNodeValueArgs) error {
	return c.Call("DOM.setNodeValue", args, nil)
}

/*
Sets node HTML markup, returns new node id.
*/
func SetOuterHTML(c protocol.Caller, args SetOuterHTMLArgs) error {
	return c.Call("DOM.setOuterHTML", args, nil)
}

/*
Undoes the last performed action.
*/
func Undo(c protocol.Caller) error {
	return c.Call("DOM.undo", nil, nil)
}

/*
Returns iframe node that owns iframe with the given domain.
*/
func GetFrameOwner(c protocol.Caller, args GetFrameOwnerArgs) (*GetFrameOwnerVal, error) {
	var val = &GetFrameOwnerVal{}
	return val, c.Call("DOM.getFrameOwner", args, val)
}

/*
	Returns the query container of the given node based on container query

conditions: containerName, physical, and logical axes. If no axes are
provided, the style container is returned, which is the direct parent or the
closest element with a matching container-name.
*/
func GetContainerForNode(c protocol.Caller, args GetContainerForNodeArgs) (*GetContainerForNodeVal, error) {
	var val = &GetContainerForNodeVal{}
	return val, c.Call("DOM.getContainerForNode", args, val)
}

/*
	Returns the descendants of a container query container that have

container queries against this container.
*/
func GetQueryingDescendantsForContainer(c protocol.Caller, args GetQueryingDescendantsForContainerArgs) (*GetQueryingDescendantsForContainerVal, error) {
	var val = &GetQueryingDescendantsForContainerVal{}
	return val, c.Call("DOM.getQueryingDescendantsForContainer", args, val)
}
