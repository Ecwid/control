package layertree

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
)

/*
Unique Layer identifier.
*/
type LayerId string

/*
Unique snapshot identifier.
*/
type SnapshotId string

/*
Rectangle where scrolling happens on the main thread.
*/
type ScrollRect struct {
	Rect *common.Rect `json:"rect"`
	Type string       `json:"type"`
}

/*
Sticky position constraints.
*/
type StickyPositionConstraint struct {
	StickyBoxRect                       *common.Rect `json:"stickyBoxRect"`
	ContainingBlockRect                 *common.Rect `json:"containingBlockRect"`
	NearestLayerShiftingStickyBox       LayerId      `json:"nearestLayerShiftingStickyBox,omitempty"`
	NearestLayerShiftingContainingBlock LayerId      `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

/*
Serialized fragment of layer picture along with its offset within the layer.
*/
type PictureTile struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Picture []byte  `json:"picture"`
}

/*
Information about a compositing layer.
*/
type Layer struct {
	LayerId                  LayerId                   `json:"layerId"`
	ParentLayerId            LayerId                   `json:"parentLayerId,omitempty"`
	BackendNodeId            dom.BackendNodeId         `json:"backendNodeId,omitempty"`
	OffsetX                  float64                   `json:"offsetX"`
	OffsetY                  float64                   `json:"offsetY"`
	Width                    float64                   `json:"width"`
	Height                   float64                   `json:"height"`
	Transform                []float64                 `json:"transform,omitempty"`
	AnchorX                  float64                   `json:"anchorX,omitempty"`
	AnchorY                  float64                   `json:"anchorY,omitempty"`
	AnchorZ                  float64                   `json:"anchorZ,omitempty"`
	PaintCount               int                       `json:"paintCount"`
	DrawsContent             bool                      `json:"drawsContent"`
	Invisible                bool                      `json:"invisible,omitempty"`
	ScrollRects              []*ScrollRect             `json:"scrollRects,omitempty"`
	StickyPositionConstraint *StickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}

/*
Array of timings, one per paint step.
*/
type PaintProfile []float64

type CompositingReasonsArgs struct {
	LayerId LayerId `json:"layerId"`
}

type CompositingReasonsVal struct {
	CompositingReasonIds []string `json:"compositingReasonIds"`
}

type LoadSnapshotArgs struct {
	Tiles []*PictureTile `json:"tiles"`
}

type LoadSnapshotVal struct {
	SnapshotId SnapshotId `json:"snapshotId"`
}

type MakeSnapshotArgs struct {
	LayerId LayerId `json:"layerId"`
}

type MakeSnapshotVal struct {
	SnapshotId SnapshotId `json:"snapshotId"`
}

type ProfileSnapshotArgs struct {
	SnapshotId     SnapshotId   `json:"snapshotId"`
	MinRepeatCount int          `json:"minRepeatCount,omitempty"`
	MinDuration    float64      `json:"minDuration,omitempty"`
	ClipRect       *common.Rect `json:"clipRect,omitempty"`
}

type ProfileSnapshotVal struct {
	Timings []PaintProfile `json:"timings"`
}

type ReleaseSnapshotArgs struct {
	SnapshotId SnapshotId `json:"snapshotId"`
}

type ReplaySnapshotArgs struct {
	SnapshotId SnapshotId `json:"snapshotId"`
	FromStep   int        `json:"fromStep,omitempty"`
	ToStep     int        `json:"toStep,omitempty"`
	Scale      float64    `json:"scale,omitempty"`
}

type ReplaySnapshotVal struct {
	DataURL string `json:"dataURL"`
}

type SnapshotCommandLogArgs struct {
	SnapshotId SnapshotId `json:"snapshotId"`
}

type SnapshotCommandLogVal struct {
	CommandLog []interface{} `json:"commandLog"`
}
