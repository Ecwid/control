package overlay

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

/*
Configuration data for drawing the source order of an elements children.
*/
type SourceOrderConfig struct {
	ParentOutlineColor *dom.RGBA `json:"parentOutlineColor"`
	ChildOutlineColor  *dom.RGBA `json:"childOutlineColor"`
}

/*
Configuration data for the highlighting of Grid elements.
*/
type GridHighlightConfig struct {
	ShowGridExtensionLines  bool      `json:"showGridExtensionLines,omitempty"`
	ShowPositiveLineNumbers bool      `json:"showPositiveLineNumbers,omitempty"`
	ShowNegativeLineNumbers bool      `json:"showNegativeLineNumbers,omitempty"`
	ShowAreaNames           bool      `json:"showAreaNames,omitempty"`
	ShowLineNames           bool      `json:"showLineNames,omitempty"`
	ShowTrackSizes          bool      `json:"showTrackSizes,omitempty"`
	GridBorderColor         *dom.RGBA `json:"gridBorderColor,omitempty"`
	RowLineColor            *dom.RGBA `json:"rowLineColor,omitempty"`
	ColumnLineColor         *dom.RGBA `json:"columnLineColor,omitempty"`
	GridBorderDash          bool      `json:"gridBorderDash,omitempty"`
	RowLineDash             bool      `json:"rowLineDash,omitempty"`
	ColumnLineDash          bool      `json:"columnLineDash,omitempty"`
	RowGapColor             *dom.RGBA `json:"rowGapColor,omitempty"`
	RowHatchColor           *dom.RGBA `json:"rowHatchColor,omitempty"`
	ColumnGapColor          *dom.RGBA `json:"columnGapColor,omitempty"`
	ColumnHatchColor        *dom.RGBA `json:"columnHatchColor,omitempty"`
	AreaBorderColor         *dom.RGBA `json:"areaBorderColor,omitempty"`
	GridBackgroundColor     *dom.RGBA `json:"gridBackgroundColor,omitempty"`
}

/*
Configuration data for the highlighting of Flex container elements.
*/
type FlexContainerHighlightConfig struct {
	ContainerBorder       *LineStyle `json:"containerBorder,omitempty"`
	LineSeparator         *LineStyle `json:"lineSeparator,omitempty"`
	ItemSeparator         *LineStyle `json:"itemSeparator,omitempty"`
	MainDistributedSpace  *BoxStyle  `json:"mainDistributedSpace,omitempty"`
	CrossDistributedSpace *BoxStyle  `json:"crossDistributedSpace,omitempty"`
	RowGapSpace           *BoxStyle  `json:"rowGapSpace,omitempty"`
	ColumnGapSpace        *BoxStyle  `json:"columnGapSpace,omitempty"`
	CrossAlignment        *LineStyle `json:"crossAlignment,omitempty"`
}

/*
Configuration data for the highlighting of Flex item elements.
*/
type FlexItemHighlightConfig struct {
	BaseSizeBox      *BoxStyle  `json:"baseSizeBox,omitempty"`
	BaseSizeBorder   *LineStyle `json:"baseSizeBorder,omitempty"`
	FlexibilityArrow *LineStyle `json:"flexibilityArrow,omitempty"`
}

/*
Style information for drawing a line.
*/
type LineStyle struct {
	Color   *dom.RGBA `json:"color,omitempty"`
	Pattern string    `json:"pattern,omitempty"`
}

/*
Style information for drawing a box.
*/
type BoxStyle struct {
	FillColor  *dom.RGBA `json:"fillColor,omitempty"`
	HatchColor *dom.RGBA `json:"hatchColor,omitempty"`
}

/*
 */
type ContrastAlgorithm string

/*
Configuration data for the highlighting of page elements.
*/
type HighlightConfig struct {
	ShowInfo                               bool                                    `json:"showInfo,omitempty"`
	ShowStyles                             bool                                    `json:"showStyles,omitempty"`
	ShowRulers                             bool                                    `json:"showRulers,omitempty"`
	ShowAccessibilityInfo                  bool                                    `json:"showAccessibilityInfo,omitempty"`
	ShowExtensionLines                     bool                                    `json:"showExtensionLines,omitempty"`
	ContentColor                           *dom.RGBA                               `json:"contentColor,omitempty"`
	PaddingColor                           *dom.RGBA                               `json:"paddingColor,omitempty"`
	BorderColor                            *dom.RGBA                               `json:"borderColor,omitempty"`
	MarginColor                            *dom.RGBA                               `json:"marginColor,omitempty"`
	EventTargetColor                       *dom.RGBA                               `json:"eventTargetColor,omitempty"`
	ShapeColor                             *dom.RGBA                               `json:"shapeColor,omitempty"`
	ShapeMarginColor                       *dom.RGBA                               `json:"shapeMarginColor,omitempty"`
	CssGridColor                           *dom.RGBA                               `json:"cssGridColor,omitempty"`
	ColorFormat                            ColorFormat                             `json:"colorFormat,omitempty"`
	GridHighlightConfig                    *GridHighlightConfig                    `json:"gridHighlightConfig,omitempty"`
	FlexContainerHighlightConfig           *FlexContainerHighlightConfig           `json:"flexContainerHighlightConfig,omitempty"`
	FlexItemHighlightConfig                *FlexItemHighlightConfig                `json:"flexItemHighlightConfig,omitempty"`
	ContrastAlgorithm                      ContrastAlgorithm                       `json:"contrastAlgorithm,omitempty"`
	ContainerQueryContainerHighlightConfig *ContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig,omitempty"`
}

/*
 */
type ColorFormat string

/*
Configurations for Persistent Grid Highlight
*/
type GridNodeHighlightConfig struct {
	GridHighlightConfig *GridHighlightConfig `json:"gridHighlightConfig"`
	NodeId              dom.NodeId           `json:"nodeId"`
}

/*
 */
type FlexNodeHighlightConfig struct {
	FlexContainerHighlightConfig *FlexContainerHighlightConfig `json:"flexContainerHighlightConfig"`
	NodeId                       dom.NodeId                    `json:"nodeId"`
}

/*
 */
type ScrollSnapContainerHighlightConfig struct {
	SnapportBorder     *LineStyle `json:"snapportBorder,omitempty"`
	SnapAreaBorder     *LineStyle `json:"snapAreaBorder,omitempty"`
	ScrollMarginColor  *dom.RGBA  `json:"scrollMarginColor,omitempty"`
	ScrollPaddingColor *dom.RGBA  `json:"scrollPaddingColor,omitempty"`
}

/*
 */
type ScrollSnapHighlightConfig struct {
	ScrollSnapContainerHighlightConfig *ScrollSnapContainerHighlightConfig `json:"scrollSnapContainerHighlightConfig"`
	NodeId                             dom.NodeId                          `json:"nodeId"`
}

/*
Configuration for dual screen hinge
*/
type HingeConfig struct {
	Rect         *common.Rect `json:"rect"`
	ContentColor *dom.RGBA    `json:"contentColor,omitempty"`
	OutlineColor *dom.RGBA    `json:"outlineColor,omitempty"`
}

/*
 */
type ContainerQueryHighlightConfig struct {
	ContainerQueryContainerHighlightConfig *ContainerQueryContainerHighlightConfig `json:"containerQueryContainerHighlightConfig"`
	NodeId                                 dom.NodeId                              `json:"nodeId"`
}

/*
 */
type ContainerQueryContainerHighlightConfig struct {
	ContainerBorder  *LineStyle `json:"containerBorder,omitempty"`
	DescendantBorder *LineStyle `json:"descendantBorder,omitempty"`
}

/*
 */
type IsolatedElementHighlightConfig struct {
	IsolationModeHighlightConfig *IsolationModeHighlightConfig `json:"isolationModeHighlightConfig"`
	NodeId                       dom.NodeId                    `json:"nodeId"`
}

/*
 */
type IsolationModeHighlightConfig struct {
	ResizerColor       *dom.RGBA `json:"resizerColor,omitempty"`
	ResizerHandleColor *dom.RGBA `json:"resizerHandleColor,omitempty"`
	MaskColor          *dom.RGBA `json:"maskColor,omitempty"`
}

/*
 */
type InspectMode string

type GetHighlightObjectForTestArgs struct {
	NodeId                dom.NodeId  `json:"nodeId"`
	IncludeDistance       bool        `json:"includeDistance,omitempty"`
	IncludeStyle          bool        `json:"includeStyle,omitempty"`
	ColorFormat           ColorFormat `json:"colorFormat,omitempty"`
	ShowAccessibilityInfo bool        `json:"showAccessibilityInfo,omitempty"`
}

type GetHighlightObjectForTestVal struct {
	Highlight interface{} `json:"highlight"`
}

type GetGridHighlightObjectsForTestArgs struct {
	NodeIds []dom.NodeId `json:"nodeIds"`
}

type GetGridHighlightObjectsForTestVal struct {
	Highlights interface{} `json:"highlights"`
}

type GetSourceOrderHighlightObjectForTestArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetSourceOrderHighlightObjectForTestVal struct {
	Highlight interface{} `json:"highlight"`
}

type HighlightNodeArgs struct {
	HighlightConfig *HighlightConfig       `json:"highlightConfig"`
	NodeId          dom.NodeId             `json:"nodeId,omitempty"`
	BackendNodeId   dom.BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId        runtime.RemoteObjectId `json:"objectId,omitempty"`
	Selector        string                 `json:"selector,omitempty"`
}

type HighlightQuadArgs struct {
	Quad         dom.Quad  `json:"quad"`
	Color        *dom.RGBA `json:"color,omitempty"`
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

type HighlightRectArgs struct {
	X            int       `json:"x"`
	Y            int       `json:"y"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Color        *dom.RGBA `json:"color,omitempty"`
	OutlineColor *dom.RGBA `json:"outlineColor,omitempty"`
}

type HighlightSourceOrderArgs struct {
	SourceOrderConfig *SourceOrderConfig     `json:"sourceOrderConfig"`
	NodeId            dom.NodeId             `json:"nodeId,omitempty"`
	BackendNodeId     dom.BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId          runtime.RemoteObjectId `json:"objectId,omitempty"`
}

type SetInspectModeArgs struct {
	Mode            InspectMode      `json:"mode"`
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`
}

type SetShowAdHighlightsArgs struct {
	Show bool `json:"show"`
}

type SetPausedInDebuggerMessageArgs struct {
	Message string `json:"message,omitempty"`
}

type SetShowDebugBordersArgs struct {
	Show bool `json:"show"`
}

type SetShowFPSCounterArgs struct {
	Show bool `json:"show"`
}

type SetShowGridOverlaysArgs struct {
	GridNodeHighlightConfigs []*GridNodeHighlightConfig `json:"gridNodeHighlightConfigs"`
}

type SetShowFlexOverlaysArgs struct {
	FlexNodeHighlightConfigs []*FlexNodeHighlightConfig `json:"flexNodeHighlightConfigs"`
}

type SetShowScrollSnapOverlaysArgs struct {
	ScrollSnapHighlightConfigs []*ScrollSnapHighlightConfig `json:"scrollSnapHighlightConfigs"`
}

type SetShowContainerQueryOverlaysArgs struct {
	ContainerQueryHighlightConfigs []*ContainerQueryHighlightConfig `json:"containerQueryHighlightConfigs"`
}

type SetShowPaintRectsArgs struct {
	Result bool `json:"result"`
}

type SetShowLayoutShiftRegionsArgs struct {
	Result bool `json:"result"`
}

type SetShowScrollBottleneckRectsArgs struct {
	Show bool `json:"show"`
}

type SetShowWebVitalsArgs struct {
	Show bool `json:"show"`
}

type SetShowViewportSizeOnResizeArgs struct {
	Show bool `json:"show"`
}

type SetShowHingeArgs struct {
	HingeConfig *HingeConfig `json:"hingeConfig,omitempty"`
}

type SetShowIsolatedElementsArgs struct {
	IsolatedElementHighlightConfigs []*IsolatedElementHighlightConfig `json:"isolatedElementHighlightConfigs"`
}
