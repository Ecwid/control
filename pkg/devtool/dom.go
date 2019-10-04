package devtool

import (
	"math"
)

// DescribeNode https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-describeNode
type DescribeNode struct {
	Node *Node `json:"node"`
}

// Node https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Node
type Node struct {
	NodeID           int64    `json:"nodeId"`
	ParentID         int64    `json:"parentId"`
	BackendNodeID    int64    `json:"backendNodeId"`
	NodeType         int64    `json:"nodeType"`
	NodeName         string   `json:"nodeName"`
	LocalName        string   `json:"localName"`
	NodeValue        string   `json:"nodeValue"`
	ChildNodeCount   int64    `json:"childNodeCount"`
	Children         []*Node  `json:"children"`
	Attributes       []string `json:"attributes"`
	DocumentURL      string   `json:"documentURL"`
	BaseURL          string   `json:"baseURL"`
	PublicID         string   `json:"publicId"`
	SystemID         string   `json:"systemId"`
	InternalSubset   string   `json:"internalSubset"`
	XMLVersion       string   `json:"xmlVersion"`
	Name             string   `json:"name"`
	Value            string   `json:"value"`
	PseudoType       string   `json:"pseudoType"`
	ShadowRootType   string   `json:"shadowRootType"`
	FrameID          string   `json:"frameId"`
	ContentDocument  *Node    `json:"contentDocument"`
	ShadowRoots      []*Node  `json:"shadowRoots"`
	TemplateContent  *Node    `json:"templateContent"`
	PseudoElements   []*Node  `json:"pseudoElements"`
	ImportedDocument *Node    `json:"importedDocument"`
	IsSVG            bool     `json:"isSVG"`
}

// EventListeners https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#type-EventListener
type EventListeners struct {
	Listeners []*EventListener `json:"listeners"`
}

// EventListener https://chromedevtools.github.io/devtools-protocol/tot/DOMDebugger#type-EventListener
type EventListener struct {
	Type            string        `json:"type"`
	UseCapture      bool          `json:"useCapture"`
	Passive         bool          `json:"passive"`
	Once            bool          `json:"once"`
	ScriptID        string        `json:"scriptId"`
	LineNumber      int64         `json:"lineNumber"`
	ColumnNumber    int64         `json:"columnNumber"`
	Handler         *RemoteObject `json:"handler"`
	OriginalHandler *RemoteObject `json:"originalHandler"`
	BackendNodeID   int64         `json:"backendNodeId"`
}

// Quad https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Quad
// type Quad []float64

// Point point
type Point struct {
	X float64
	Y float64
}

// Quad quad
type Quad []*Point

// ContentQuads https://chromedevtools.github.io/devtools-protocol/tot/DOM#method-getContentQuads
type ContentQuads struct {
	Quads [][]float64 `json:"quads"`
}

// Calc coverts to quads
func (c ContentQuads) Calc() []Quad {
	p := make([]Quad, len(c.Quads))
	for n, q := range c.Quads {
		p[n] = Quad{
			&Point{q[0], q[1]},
			&Point{q[2], q[3]},
			&Point{q[4], q[5]},
			&Point{q[6], q[7]},
		}
	}
	return p
}

// Rect https://chromedevtools.github.io/devtools-protocol/tot/DOM#type-Rect
type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Middle calc middle of quad
func (q Quad) Middle() (float64, float64) {
	x := 0.0
	y := 0.0
	for i := 0; i < 4; i++ {
		x += q[i].X
		y += q[i].Y
	}
	return x / 4, y / 4
}

// Area calc area of quad
func (q Quad) Area() float64 {
	var area float64
	var x1, x2, y1, y2 float64
	vertices := len(q)
	for i := 0; i < vertices; i++ {
		x1 = q[i].X
		y1 = q[i].Y
		x2 = q[(i+1)%vertices].X
		y2 = q[(i+1)%vertices].Y
		area += (x1*y2 - x2*y1) / 2
	}
	return math.Abs(area)
}
