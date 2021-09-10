package accessibility

import (
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Unique accessibility node identifier.
*/
type AXNodeId string

/*
	Enum of possible property types.
*/
type AXValueType string

/*
	Enum of possible property sources.
*/
type AXValueSourceType string

/*
	Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
*/
type AXValueNativeSourceType string

/*
	A single source for a computed AX property.
*/
type AXValueSource struct {
	Type              AXValueSourceType       `json:"type"`
	Value             *AXValue                `json:"value,omitempty"`
	Attribute         string                  `json:"attribute,omitempty"`
	AttributeValue    *AXValue                `json:"attributeValue,omitempty"`
	Superseded        bool                    `json:"superseded,omitempty"`
	NativeSource      AXValueNativeSourceType `json:"nativeSource,omitempty"`
	NativeSourceValue *AXValue                `json:"nativeSourceValue,omitempty"`
	Invalid           bool                    `json:"invalid,omitempty"`
	InvalidReason     string                  `json:"invalidReason,omitempty"`
}

/*

 */
type AXRelatedNode struct {
	BackendDOMNodeId dom.BackendNodeId `json:"backendDOMNodeId"`
	Idref            string            `json:"idref,omitempty"`
	Text             string            `json:"text,omitempty"`
}

/*

 */
type AXProperty struct {
	Name  AXPropertyName `json:"name"`
	Value *AXValue       `json:"value"`
}

/*
	A single computed AX property.
*/
type AXValue struct {
	Type         AXValueType      `json:"type"`
	Value        interface{}      `json:"value,omitempty"`
	RelatedNodes []*AXRelatedNode `json:"relatedNodes,omitempty"`
	Sources      []*AXValueSource `json:"sources,omitempty"`
}

/*
	Values of AXProperty name:
- from 'busy' to 'roledescription': states which apply to every AX node
- from 'live' to 'root': attributes which apply to nodes in live regions
- from 'autocomplete' to 'valuetext': attributes which apply to widgets
- from 'checked' to 'selected': states which apply to widgets
- from 'activedescendant' to 'owns' - relationships between elements other than parent/child/sibling.
*/
type AXPropertyName string

/*
	A node in the accessibility tree.
*/
type AXNode struct {
	NodeId           AXNodeId          `json:"nodeId"`
	Ignored          bool              `json:"ignored"`
	IgnoredReasons   []*AXProperty     `json:"ignoredReasons,omitempty"`
	Role             *AXValue          `json:"role,omitempty"`
	Name             *AXValue          `json:"name,omitempty"`
	Description      *AXValue          `json:"description,omitempty"`
	Value            *AXValue          `json:"value,omitempty"`
	Properties       []*AXProperty     `json:"properties,omitempty"`
	ChildIds         []AXNodeId        `json:"childIds,omitempty"`
	BackendDOMNodeId dom.BackendNodeId `json:"backendDOMNodeId,omitempty"`
}

type GetPartialAXTreeArgs struct {
	NodeId         dom.NodeId             `json:"nodeId,omitempty"`
	BackendNodeId  dom.BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId       runtime.RemoteObjectId `json:"objectId,omitempty"`
	FetchRelatives bool                   `json:"fetchRelatives,omitempty"`
}

type GetPartialAXTreeVal struct {
	Nodes []*AXNode `json:"nodes"`
}

type GetFullAXTreeArgs struct {
	Max_depth int `json:"max_depth,omitempty"`
}

type GetFullAXTreeVal struct {
	Nodes []*AXNode `json:"nodes"`
}

type GetChildAXNodesArgs struct {
	Id AXNodeId `json:"id"`
}

type GetChildAXNodesVal struct {
	Nodes []*AXNode `json:"nodes"`
}

type QueryAXTreeArgs struct {
	NodeId         dom.NodeId             `json:"nodeId,omitempty"`
	BackendNodeId  dom.BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId       runtime.RemoteObjectId `json:"objectId,omitempty"`
	AccessibleName string                 `json:"accessibleName,omitempty"`
	Role           string                 `json:"role,omitempty"`
}

type QueryAXTreeVal struct {
	Nodes []*AXNode `json:"nodes"`
}
