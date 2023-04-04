package dom

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/runtime"
)

/*
Unique DOM node identifier.
*/
type NodeId int

/*
	Unique DOM node identifier used to reference a node that may not have been pushed to the

front-end.
*/
type BackendNodeId int

/*
Backend node with a friendly name.
*/
type BackendNode struct {
	NodeType      int           `json:"nodeType"`
	NodeName      string        `json:"nodeName"`
	BackendNodeId BackendNodeId `json:"backendNodeId"`
}

/*
Pseudo element type.
*/
type PseudoType string

/*
Shadow root type.
*/
type ShadowRootType string

/*
Document compatibility mode.
*/
type CompatibilityMode string

/*
ContainerSelector physical axes
*/
type PhysicalAxes string

/*
ContainerSelector logical axes
*/
type LogicalAxes string

/*
	DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes.

DOMNode is a base node mirror type.
*/
type Node struct {
	NodeId            NodeId            `json:"nodeId"`
	ParentId          NodeId            `json:"parentId,omitempty"`
	BackendNodeId     BackendNodeId     `json:"backendNodeId"`
	NodeType          int               `json:"nodeType"`
	NodeName          string            `json:"nodeName"`
	LocalName         string            `json:"localName"`
	NodeValue         string            `json:"nodeValue"`
	ChildNodeCount    int               `json:"childNodeCount,omitempty"`
	Children          []*Node           `json:"children,omitempty"`
	Attributes        []string          `json:"attributes,omitempty"`
	DocumentURL       string            `json:"documentURL,omitempty"`
	BaseURL           string            `json:"baseURL,omitempty"`
	PublicId          string            `json:"publicId,omitempty"`
	SystemId          string            `json:"systemId,omitempty"`
	InternalSubset    string            `json:"internalSubset,omitempty"`
	XmlVersion        string            `json:"xmlVersion,omitempty"`
	Name              string            `json:"name,omitempty"`
	Value             string            `json:"value,omitempty"`
	PseudoType        PseudoType        `json:"pseudoType,omitempty"`
	PseudoIdentifier  string            `json:"pseudoIdentifier,omitempty"`
	ShadowRootType    ShadowRootType    `json:"shadowRootType,omitempty"`
	FrameId           common.FrameId    `json:"frameId,omitempty"`
	ContentDocument   *Node             `json:"contentDocument,omitempty"`
	ShadowRoots       []*Node           `json:"shadowRoots,omitempty"`
	TemplateContent   *Node             `json:"templateContent,omitempty"`
	PseudoElements    []*Node           `json:"pseudoElements,omitempty"`
	DistributedNodes  []*BackendNode    `json:"distributedNodes,omitempty"`
	IsSVG             bool              `json:"isSVG,omitempty"`
	CompatibilityMode CompatibilityMode `json:"compatibilityMode,omitempty"`
	AssignedSlot      *BackendNode      `json:"assignedSlot,omitempty"`
}

/*
A structure holding an RGBA color.
*/
type RGBA struct {
	R int     `json:"r"`
	G int     `json:"g"`
	B int     `json:"b"`
	A float64 `json:"a,omitempty"`
}

/*
An array of quad vertices, x immediately followed by y for each point, points clock-wise.
*/
type Quad []float64

/*
Box model.
*/
type BoxModel struct {
	Content      Quad              `json:"content"`
	Padding      Quad              `json:"padding"`
	Border       Quad              `json:"border"`
	Margin       Quad              `json:"margin"`
	Width        int               `json:"width"`
	Height       int               `json:"height"`
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

/*
CSS Shape Outside details.
*/
type ShapeOutsideInfo struct {
	Bounds      Quad          `json:"bounds"`
	Shape       []interface{} `json:"shape"`
	MarginShape []interface{} `json:"marginShape"`
}

/*
Rectangle.
*/
type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

/*
 */
type CSSComputedStyleProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CollectClassNamesFromSubtreeArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type CollectClassNamesFromSubtreeVal struct {
	ClassNames []string `json:"classNames"`
}

type CopyToArgs struct {
	NodeId             NodeId `json:"nodeId"`
	TargetNodeId       NodeId `json:"targetNodeId"`
	InsertBeforeNodeId NodeId `json:"insertBeforeNodeId,omitempty"`
}

type CopyToVal struct {
	NodeId NodeId `json:"nodeId"`
}

type DescribeNodeArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
	Depth         int                    `json:"depth,omitempty"`
	Pierce        bool                   `json:"pierce,omitempty"`
}

type DescribeNodeVal struct {
	Node *Node `json:"node"`
}

type ScrollIntoViewIfNeededArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
	Rect          *common.Rect           `json:"rect,omitempty"`
}

type DiscardSearchResultsArgs struct {
	SearchId string `json:"searchId"`
}

type EnableArgs struct {
	IncludeWhitespace string `json:"includeWhitespace,omitempty"`
}

type FocusArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type GetAttributesArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type GetAttributesVal struct {
	Attributes []string `json:"attributes"`
}

type GetBoxModelArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type GetBoxModelVal struct {
	Model *BoxModel `json:"model"`
}

type GetContentQuadsArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type GetContentQuadsVal struct {
	Quads []Quad `json:"quads"`
}

type GetDocumentArgs struct {
	Depth  int  `json:"depth,omitempty"`
	Pierce bool `json:"pierce,omitempty"`
}

type GetDocumentVal struct {
	Root *Node `json:"root"`
}

type GetNodesForSubtreeByStyleArgs struct {
	NodeId         NodeId                      `json:"nodeId"`
	ComputedStyles []*CSSComputedStyleProperty `json:"computedStyles"`
	Pierce         bool                        `json:"pierce,omitempty"`
}

type GetNodesForSubtreeByStyleVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}

type GetNodeForLocationArgs struct {
	X                         int  `json:"x"`
	Y                         int  `json:"y"`
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
	IgnorePointerEventsNone   bool `json:"ignorePointerEventsNone,omitempty"`
}

type GetNodeForLocationVal struct {
	BackendNodeId BackendNodeId  `json:"backendNodeId"`
	FrameId       common.FrameId `json:"frameId"`
	NodeId        NodeId         `json:"nodeId,omitempty"`
}

type GetOuterHTMLArgs struct {
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type GetOuterHTMLVal struct {
	OuterHTML string `json:"outerHTML"`
}

type GetRelayoutBoundaryArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type GetRelayoutBoundaryVal struct {
	NodeId NodeId `json:"nodeId"`
}

type GetSearchResultsArgs struct {
	SearchId  string `json:"searchId"`
	FromIndex int    `json:"fromIndex"`
	ToIndex   int    `json:"toIndex"`
}

type GetSearchResultsVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}

type MoveToArgs struct {
	NodeId             NodeId `json:"nodeId"`
	TargetNodeId       NodeId `json:"targetNodeId"`
	InsertBeforeNodeId NodeId `json:"insertBeforeNodeId,omitempty"`
}

type MoveToVal struct {
	NodeId NodeId `json:"nodeId"`
}

type PerformSearchArgs struct {
	Query                     string `json:"query"`
	IncludeUserAgentShadowDOM bool   `json:"includeUserAgentShadowDOM,omitempty"`
}

type PerformSearchVal struct {
	SearchId    string `json:"searchId"`
	ResultCount int    `json:"resultCount"`
}

type PushNodeByPathToFrontendArgs struct {
	Path string `json:"path"`
}

type PushNodeByPathToFrontendVal struct {
	NodeId NodeId `json:"nodeId"`
}

type PushNodesByBackendIdsToFrontendArgs struct {
	BackendNodeIds []BackendNodeId `json:"backendNodeIds"`
}

type PushNodesByBackendIdsToFrontendVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}

type QuerySelectorArgs struct {
	NodeId   NodeId `json:"nodeId"`
	Selector string `json:"selector"`
}

type QuerySelectorVal struct {
	NodeId NodeId `json:"nodeId"`
}

type QuerySelectorAllArgs struct {
	NodeId   NodeId `json:"nodeId"`
	Selector string `json:"selector"`
}

type QuerySelectorAllVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}

type GetTopLayerElementsVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}

type RemoveAttributeArgs struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
}

type RemoveNodeArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type RequestChildNodesArgs struct {
	NodeId NodeId `json:"nodeId"`
	Depth  int    `json:"depth,omitempty"`
	Pierce bool   `json:"pierce,omitempty"`
}

type RequestNodeArgs struct {
	ObjectId runtime.RemoteObjectId `json:"objectId"`
}

type RequestNodeVal struct {
	NodeId NodeId `json:"nodeId"`
}

type ResolveNodeArgs struct {
	NodeId             NodeId                     `json:"nodeId,omitempty"`
	BackendNodeId      BackendNodeId              `json:"backendNodeId,omitempty"`
	ObjectGroup        string                     `json:"objectGroup,omitempty"`
	ExecutionContextId runtime.ExecutionContextId `json:"executionContextId,omitempty"`
}

type ResolveNodeVal struct {
	Object *runtime.RemoteObject `json:"object"`
}

type SetAttributeValueArgs struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type SetAttributesAsTextArgs struct {
	NodeId NodeId `json:"nodeId"`
	Text   string `json:"text"`
	Name   string `json:"name,omitempty"`
}

type SetFileInputFilesArgs struct {
	Files         []string               `json:"files"`
	NodeId        NodeId                 `json:"nodeId,omitempty"`
	BackendNodeId BackendNodeId          `json:"backendNodeId,omitempty"`
	ObjectId      runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type SetNodeStackTracesEnabledArgs struct {
	Enable bool `json:"enable"`
}

type GetNodeStackTracesArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type GetNodeStackTracesVal struct {
	Creation *runtime.StackTrace `json:"creation,omitempty"`
}

type GetFileInfoArgs struct {
	ObjectId runtime.RemoteObjectId `json:"objectId"`
}

type GetFileInfoVal struct {
	Path string `json:"path"`
}

type SetInspectedNodeArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type SetNodeNameArgs struct {
	NodeId NodeId `json:"nodeId"`
	Name   string `json:"name"`
}

type SetNodeNameVal struct {
	NodeId NodeId `json:"nodeId"`
}

type SetNodeValueArgs struct {
	NodeId NodeId `json:"nodeId"`
	Value  string `json:"value"`
}

type SetOuterHTMLArgs struct {
	NodeId    NodeId `json:"nodeId"`
	OuterHTML string `json:"outerHTML"`
}

type GetFrameOwnerArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetFrameOwnerVal struct {
	BackendNodeId BackendNodeId `json:"backendNodeId"`
	NodeId        NodeId        `json:"nodeId,omitempty"`
}

type GetContainerForNodeArgs struct {
	NodeId        NodeId       `json:"nodeId"`
	ContainerName string       `json:"containerName,omitempty"`
	PhysicalAxes  PhysicalAxes `json:"physicalAxes,omitempty"`
	LogicalAxes   LogicalAxes  `json:"logicalAxes,omitempty"`
}

type GetContainerForNodeVal struct {
	NodeId NodeId `json:"nodeId,omitempty"`
}

type GetQueryingDescendantsForContainerArgs struct {
	NodeId NodeId `json:"nodeId"`
}

type GetQueryingDescendantsForContainerVal struct {
	NodeIds []NodeId `json:"nodeIds"`
}
