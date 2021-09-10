package dom

/*
	Fired when `Element`'s attribute is modified.
*/
type AttributeModified struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

/*
	Fired when `Element`'s attribute is removed.
*/
type AttributeRemoved struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
}

/*
	Mirrors `DOMCharacterDataModified` event.
*/
type CharacterDataModified struct {
	NodeId        NodeId `json:"nodeId"`
	CharacterData string `json:"characterData"`
}

/*
	Fired when `Container`'s child node count has changed.
*/
type ChildNodeCountUpdated struct {
	NodeId         NodeId `json:"nodeId"`
	ChildNodeCount int    `json:"childNodeCount"`
}

/*
	Mirrors `DOMNodeInserted` event.
*/
type ChildNodeInserted struct {
	ParentNodeId   NodeId `json:"parentNodeId"`
	PreviousNodeId NodeId `json:"previousNodeId"`
	Node           *Node  `json:"node"`
}

/*
	Mirrors `DOMNodeRemoved` event.
*/
type ChildNodeRemoved struct {
	ParentNodeId NodeId `json:"parentNodeId"`
	NodeId       NodeId `json:"nodeId"`
}

/*
	Called when distrubution is changed.
*/
type DistributedNodesUpdated struct {
	InsertionPointId NodeId         `json:"insertionPointId"`
	DistributedNodes []*BackendNode `json:"distributedNodes"`
}

/*
	Fired when `Document` has been totally updated. Node ids are no longer valid.
*/
type DocumentUpdated interface{}

/*
	Fired when `Element`'s inline style is modified via a CSS property modification.
*/
type InlineStyleInvalidated struct {
	NodeIds []NodeId `json:"nodeIds"`
}

/*
	Called when a pseudo element is added to an element.
*/
type PseudoElementAdded struct {
	ParentId      NodeId `json:"parentId"`
	PseudoElement *Node  `json:"pseudoElement"`
}

/*
	Called when a pseudo element is removed from an element.
*/
type PseudoElementRemoved struct {
	ParentId        NodeId `json:"parentId"`
	PseudoElementId NodeId `json:"pseudoElementId"`
}

/*
	Fired when backend wants to provide client with the missing DOM structure. This happens upon
most of the calls requesting node ids.
*/
type SetChildNodes struct {
	ParentId NodeId  `json:"parentId"`
	Nodes    []*Node `json:"nodes"`
}

/*
	Called when shadow root is popped from the element.
*/
type ShadowRootPopped struct {
	HostId NodeId `json:"hostId"`
	RootId NodeId `json:"rootId"`
}

/*
	Called when shadow root is pushed into the element.
*/
type ShadowRootPushed struct {
	HostId NodeId `json:"hostId"`
	Root   *Node  `json:"root"`
}
