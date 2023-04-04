package domsnapshot

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/domdebugger"
)

/*
A Node in the DOM tree.
*/
type DOMNode struct {
	NodeType             int                          `json:"nodeType"`
	NodeName             string                       `json:"nodeName"`
	NodeValue            string                       `json:"nodeValue"`
	TextValue            string                       `json:"textValue,omitempty"`
	InputValue           string                       `json:"inputValue,omitempty"`
	InputChecked         bool                         `json:"inputChecked,omitempty"`
	OptionSelected       bool                         `json:"optionSelected,omitempty"`
	BackendNodeId        dom.BackendNodeId            `json:"backendNodeId"`
	ChildNodeIndexes     []int                        `json:"childNodeIndexes,omitempty"`
	Attributes           []*NameValue                 `json:"attributes,omitempty"`
	PseudoElementIndexes []int                        `json:"pseudoElementIndexes,omitempty"`
	LayoutNodeIndex      int                          `json:"layoutNodeIndex,omitempty"`
	DocumentURL          string                       `json:"documentURL,omitempty"`
	BaseURL              string                       `json:"baseURL,omitempty"`
	ContentLanguage      string                       `json:"contentLanguage,omitempty"`
	DocumentEncoding     string                       `json:"documentEncoding,omitempty"`
	PublicId             string                       `json:"publicId,omitempty"`
	SystemId             string                       `json:"systemId,omitempty"`
	FrameId              common.FrameId               `json:"frameId,omitempty"`
	ContentDocumentIndex int                          `json:"contentDocumentIndex,omitempty"`
	PseudoType           dom.PseudoType               `json:"pseudoType,omitempty"`
	ShadowRootType       dom.ShadowRootType           `json:"shadowRootType,omitempty"`
	IsClickable          bool                         `json:"isClickable,omitempty"`
	EventListeners       []*domdebugger.EventListener `json:"eventListeners,omitempty"`
	CurrentSourceURL     string                       `json:"currentSourceURL,omitempty"`
	OriginURL            string                       `json:"originURL,omitempty"`
	ScrollOffsetX        float64                      `json:"scrollOffsetX,omitempty"`
	ScrollOffsetY        float64                      `json:"scrollOffsetY,omitempty"`
}

/*
	Details of post layout rendered text positions. The exact layout should not be regarded as

stable and may change between versions.
*/
type InlineTextBox struct {
	BoundingBox         *common.Rect `json:"boundingBox"`
	StartCharacterIndex int          `json:"startCharacterIndex"`
	NumCharacters       int          `json:"numCharacters"`
}

/*
Details of an element in the DOM tree with a LayoutObject.
*/
type LayoutTreeNode struct {
	DomNodeIndex      int              `json:"domNodeIndex"`
	BoundingBox       *common.Rect     `json:"boundingBox"`
	LayoutText        string           `json:"layoutText,omitempty"`
	InlineTextNodes   []*InlineTextBox `json:"inlineTextNodes,omitempty"`
	StyleIndex        int              `json:"styleIndex,omitempty"`
	PaintOrder        int              `json:"paintOrder,omitempty"`
	IsStackingContext bool             `json:"isStackingContext,omitempty"`
}

/*
A subset of the full ComputedStyle as defined by the request whitelist.
*/
type ComputedStyle struct {
	Properties []*NameValue `json:"properties"`
}

/*
A name/value pair.
*/
type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
Index of the string in the strings table.
*/
type StringIndex int

/*
Index of the string in the strings table.
*/
type ArrayOfStrings []StringIndex

/*
Data that is only present on rare nodes.
*/
type RareStringData struct {
	Index []int         `json:"index"`
	Value []StringIndex `json:"value"`
}

/*
 */
type RareBooleanData struct {
	Index []int `json:"index"`
}

/*
 */
type RareIntegerData struct {
	Index []int `json:"index"`
	Value []int `json:"value"`
}

/*
 */
type Rectangle []float64

/*
Document snapshot.
*/
type DocumentSnapshot struct {
	DocumentURL     StringIndex         `json:"documentURL"`
	Title           StringIndex         `json:"title"`
	BaseURL         StringIndex         `json:"baseURL"`
	ContentLanguage StringIndex         `json:"contentLanguage"`
	EncodingName    StringIndex         `json:"encodingName"`
	PublicId        StringIndex         `json:"publicId"`
	SystemId        StringIndex         `json:"systemId"`
	FrameId         StringIndex         `json:"frameId"`
	Nodes           *NodeTreeSnapshot   `json:"nodes"`
	Layout          *LayoutTreeSnapshot `json:"layout"`
	TextBoxes       *TextBoxSnapshot    `json:"textBoxes"`
	ScrollOffsetX   float64             `json:"scrollOffsetX,omitempty"`
	ScrollOffsetY   float64             `json:"scrollOffsetY,omitempty"`
	ContentWidth    float64             `json:"contentWidth,omitempty"`
	ContentHeight   float64             `json:"contentHeight,omitempty"`
}

/*
Table containing nodes.
*/
type NodeTreeSnapshot struct {
	ParentIndex          []int               `json:"parentIndex,omitempty"`
	NodeType             []int               `json:"nodeType,omitempty"`
	ShadowRootType       *RareStringData     `json:"shadowRootType,omitempty"`
	NodeName             []StringIndex       `json:"nodeName,omitempty"`
	NodeValue            []StringIndex       `json:"nodeValue,omitempty"`
	BackendNodeId        []dom.BackendNodeId `json:"backendNodeId,omitempty"`
	Attributes           []ArrayOfStrings    `json:"attributes,omitempty"`
	TextValue            *RareStringData     `json:"textValue,omitempty"`
	InputValue           *RareStringData     `json:"inputValue,omitempty"`
	InputChecked         *RareBooleanData    `json:"inputChecked,omitempty"`
	OptionSelected       *RareBooleanData    `json:"optionSelected,omitempty"`
	ContentDocumentIndex *RareIntegerData    `json:"contentDocumentIndex,omitempty"`
	PseudoType           *RareStringData     `json:"pseudoType,omitempty"`
	PseudoIdentifier     *RareStringData     `json:"pseudoIdentifier,omitempty"`
	IsClickable          *RareBooleanData    `json:"isClickable,omitempty"`
	CurrentSourceURL     *RareStringData     `json:"currentSourceURL,omitempty"`
	OriginURL            *RareStringData     `json:"originURL,omitempty"`
}

/*
Table of details of an element in the DOM tree with a LayoutObject.
*/
type LayoutTreeSnapshot struct {
	NodeIndex               []int            `json:"nodeIndex"`
	Styles                  []ArrayOfStrings `json:"styles"`
	Bounds                  []Rectangle      `json:"bounds"`
	Text                    []StringIndex    `json:"text"`
	StackingContexts        *RareBooleanData `json:"stackingContexts"`
	PaintOrders             []int            `json:"paintOrders,omitempty"`
	OffsetRects             []Rectangle      `json:"offsetRects,omitempty"`
	ScrollRects             []Rectangle      `json:"scrollRects,omitempty"`
	ClientRects             []Rectangle      `json:"clientRects,omitempty"`
	BlendedBackgroundColors []StringIndex    `json:"blendedBackgroundColors,omitempty"`
	TextColorOpacities      []float64        `json:"textColorOpacities,omitempty"`
}

/*
	Table of details of the post layout rendered text positions. The exact layout should not be regarded as

stable and may change between versions.
*/
type TextBoxSnapshot struct {
	LayoutIndex []int       `json:"layoutIndex"`
	Bounds      []Rectangle `json:"bounds"`
	Start       []int       `json:"start"`
	Length      []int       `json:"length"`
}

type CaptureSnapshotArgs struct {
	ComputedStyles                 []string `json:"computedStyles"`
	IncludePaintOrder              bool     `json:"includePaintOrder,omitempty"`
	IncludeDOMRects                bool     `json:"includeDOMRects,omitempty"`
	IncludeBlendedBackgroundColors bool     `json:"includeBlendedBackgroundColors,omitempty"`
	IncludeTextColorOpacities      bool     `json:"includeTextColorOpacities,omitempty"`
}

type CaptureSnapshotVal struct {
	Documents []*DocumentSnapshot `json:"documents"`
	Strings   []string            `json:"strings"`
}
