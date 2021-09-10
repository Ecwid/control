package performancetimeline

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
)

/*
	See https://github.com/WICG/LargestContentfulPaint and largest_contentful_paint.idl
*/
type LargestContentfulPaint struct {
	RenderTime common.TimeSinceEpoch `json:"renderTime"`
	LoadTime   common.TimeSinceEpoch `json:"loadTime"`
	Size       float64               `json:"size"`
	ElementId  string                `json:"elementId,omitempty"`
	Url        string                `json:"url,omitempty"`
	NodeId     dom.BackendNodeId     `json:"nodeId,omitempty"`
}

/*

 */
type LayoutShiftAttribution struct {
	PreviousRect *common.Rect      `json:"previousRect"`
	CurrentRect  *common.Rect      `json:"currentRect"`
	NodeId       dom.BackendNodeId `json:"nodeId,omitempty"`
}

/*
	See https://wicg.github.io/layout-instability/#sec-layout-shift and layout_shift.idl
*/
type LayoutShift struct {
	Value          float64                   `json:"value"`
	HadRecentInput bool                      `json:"hadRecentInput"`
	LastInputTime  common.TimeSinceEpoch     `json:"lastInputTime"`
	Sources        []*LayoutShiftAttribution `json:"sources"`
}

/*

 */
type TimelineEvent struct {
	FrameId            common.FrameId          `json:"frameId"`
	Type               string                  `json:"type"`
	Name               string                  `json:"name"`
	Time               common.TimeSinceEpoch   `json:"time"`
	Duration           float64                 `json:"duration,omitempty"`
	LcpDetails         *LargestContentfulPaint `json:"lcpDetails,omitempty"`
	LayoutShiftDetails *LayoutShift            `json:"layoutShiftDetails,omitempty"`
}

type EnableArgs struct {
	EventTypes []string `json:"eventTypes"`
}
