package accessibility

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables the accessibility domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Accessibility.disable", nil, nil)
}

/*
	Enables the accessibility domain which causes `AXNodeId`s to remain consistent between method calls.

This turns on accessibility for the page, which can impact performance until accessibility is disabled.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Accessibility.enable", nil, nil)
}

/*
Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
*/
func GetPartialAXTree(c protocol.Caller, args GetPartialAXTreeArgs) (*GetPartialAXTreeVal, error) {
	var val = &GetPartialAXTreeVal{}
	return val, c.Call("Accessibility.getPartialAXTree", args, val)
}

/*
Fetches the entire accessibility tree for the root Document
*/
func GetFullAXTree(c protocol.Caller, args GetFullAXTreeArgs) (*GetFullAXTreeVal, error) {
	var val = &GetFullAXTreeVal{}
	return val, c.Call("Accessibility.getFullAXTree", args, val)
}

/*
	Fetches the root node.

Requires `enable()` to have been called previously.
*/
func GetRootAXNode(c protocol.Caller, args GetRootAXNodeArgs) (*GetRootAXNodeVal, error) {
	var val = &GetRootAXNodeVal{}
	return val, c.Call("Accessibility.getRootAXNode", args, val)
}

/*
	Fetches a node and all ancestors up to and including the root.

Requires `enable()` to have been called previously.
*/
func GetAXNodeAndAncestors(c protocol.Caller, args GetAXNodeAndAncestorsArgs) (*GetAXNodeAndAncestorsVal, error) {
	var val = &GetAXNodeAndAncestorsVal{}
	return val, c.Call("Accessibility.getAXNodeAndAncestors", args, val)
}

/*
	Fetches a particular accessibility node by AXNodeId.

Requires `enable()` to have been called previously.
*/
func GetChildAXNodes(c protocol.Caller, args GetChildAXNodesArgs) (*GetChildAXNodesVal, error) {
	var val = &GetChildAXNodesVal{}
	return val, c.Call("Accessibility.getChildAXNodes", args, val)
}

/*
	Query a DOM node's accessibility subtree for accessible name and role.

This command computes the name and role for all nodes in the subtree, including those that are
ignored for accessibility, and returns those that mactch the specified name and role. If no DOM
node is specified, or the DOM node does not exist, the command returns an error. If neither
`accessibleName` or `role` is specified, it returns all the accessibility nodes in the subtree.
*/
func QueryAXTree(c protocol.Caller, args QueryAXTreeArgs) (*QueryAXTreeVal, error) {
	var val = &QueryAXTreeVal{}
	return val, c.Call("Accessibility.queryAXTree", args, val)
}
