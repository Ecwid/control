package accessibility

/*
	The loadComplete event mirrors the load complete event sent by the browser to assistive

technology when the web page has finished loading.
*/
type LoadComplete struct {
	Root *AXNode `json:"root"`
}

/*
The nodesUpdated event is sent every time a previously requested node has changed the in tree.
*/
type NodesUpdated struct {
	Nodes []*AXNode `json:"nodes"`
}
