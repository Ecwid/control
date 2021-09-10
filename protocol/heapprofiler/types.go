package heapprofiler

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Heap snapshot object id.
*/
type HeapSnapshotObjectId string

/*
	Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
*/
type SamplingHeapProfileNode struct {
	CallFrame *runtime.CallFrame         `json:"callFrame"`
	SelfSize  float64                    `json:"selfSize"`
	Id        int                        `json:"id"`
	Children  []*SamplingHeapProfileNode `json:"children"`
}

/*
	A single sample from a sampling profile.
*/
type SamplingHeapProfileSample struct {
	Size    float64 `json:"size"`
	NodeId  int     `json:"nodeId"`
	Ordinal float64 `json:"ordinal"`
}

/*
	Sampling profile.
*/
type SamplingHeapProfile struct {
	Head    *SamplingHeapProfileNode     `json:"head"`
	Samples []*SamplingHeapProfileSample `json:"samples"`
}

type AddInspectedHeapObjectArgs struct {
	HeapObjectId HeapSnapshotObjectId `json:"heapObjectId"`
}

type GetHeapObjectIdArgs struct {
	ObjectId runtime.RemoteObjectId `json:"objectId"`
}

type GetHeapObjectIdVal struct {
	HeapSnapshotObjectId HeapSnapshotObjectId `json:"heapSnapshotObjectId"`
}

type GetObjectByHeapObjectIdArgs struct {
	ObjectId    HeapSnapshotObjectId `json:"objectId"`
	ObjectGroup string               `json:"objectGroup,omitempty"`
}

type GetObjectByHeapObjectIdVal struct {
	Result *runtime.RemoteObject `json:"result"`
}

type GetSamplingProfileVal struct {
	Profile *SamplingHeapProfile `json:"profile"`
}

type StartSamplingArgs struct {
	SamplingInterval float64 `json:"samplingInterval,omitempty"`
}

type StartTrackingHeapObjectsArgs struct {
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

type StopSamplingVal struct {
	Profile *SamplingHeapProfile `json:"profile"`
}

type StopTrackingHeapObjectsArgs struct {
	ReportProgress            bool `json:"reportProgress,omitempty"`
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
}

type TakeHeapSnapshotArgs struct {
	ReportProgress            bool `json:"reportProgress,omitempty"`
	TreatGlobalObjectsAsRoots bool `json:"treatGlobalObjectsAsRoots,omitempty"`
}
