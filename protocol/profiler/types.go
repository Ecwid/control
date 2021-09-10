package profiler

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Profile node. Holds callsite information, execution statistics and child nodes.
*/
type ProfileNode struct {
	Id            int                 `json:"id"`
	CallFrame     *runtime.CallFrame  `json:"callFrame"`
	HitCount      int                 `json:"hitCount,omitempty"`
	Children      []int               `json:"children,omitempty"`
	DeoptReason   string              `json:"deoptReason,omitempty"`
	PositionTicks []*PositionTickInfo `json:"positionTicks,omitempty"`
}

/*
	Profile.
*/
type Profile struct {
	Nodes      []*ProfileNode `json:"nodes"`
	StartTime  float64        `json:"startTime"`
	EndTime    float64        `json:"endTime"`
	Samples    []int          `json:"samples,omitempty"`
	TimeDeltas []int          `json:"timeDeltas,omitempty"`
}

/*
	Specifies a number of samples attributed to a certain source position.
*/
type PositionTickInfo struct {
	Line  int `json:"line"`
	Ticks int `json:"ticks"`
}

/*
	Coverage data for a source range.
*/
type CoverageRange struct {
	StartOffset int `json:"startOffset"`
	EndOffset   int `json:"endOffset"`
	Count       int `json:"count"`
}

/*
	Coverage data for a JavaScript function.
*/
type FunctionCoverage struct {
	FunctionName    string           `json:"functionName"`
	Ranges          []*CoverageRange `json:"ranges"`
	IsBlockCoverage bool             `json:"isBlockCoverage"`
}

/*
	Coverage data for a JavaScript script.
*/
type ScriptCoverage struct {
	ScriptId  runtime.ScriptId    `json:"scriptId"`
	Url       string              `json:"url"`
	Functions []*FunctionCoverage `json:"functions"`
}

/*
	Describes a type collected during runtime.
*/
type TypeObject struct {
	Name string `json:"name"`
}

/*
	Source offset and types for a parameter or return value.
*/
type TypeProfileEntry struct {
	Offset int           `json:"offset"`
	Types  []*TypeObject `json:"types"`
}

/*
	Type profile data collected during runtime for a JavaScript script.
*/
type ScriptTypeProfile struct {
	ScriptId runtime.ScriptId    `json:"scriptId"`
	Url      string              `json:"url"`
	Entries  []*TypeProfileEntry `json:"entries"`
}

/*
	Collected counter information.
*/
type CounterInfo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

/*
	Runtime call counter information.
*/
type RuntimeCallCounterInfo struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Time  float64 `json:"time"`
}

type GetBestEffortCoverageVal struct {
	Result []*ScriptCoverage `json:"result"`
}

type SetSamplingIntervalArgs struct {
	Interval int `json:"interval"`
}

type StartPreciseCoverageArgs struct {
	CallCount             bool `json:"callCount,omitempty"`
	Detailed              bool `json:"detailed,omitempty"`
	AllowTriggeredUpdates bool `json:"allowTriggeredUpdates,omitempty"`
}

type StartPreciseCoverageVal struct {
	Timestamp float64 `json:"timestamp"`
}

type StopVal struct {
	Profile *Profile `json:"profile"`
}

type TakePreciseCoverageVal struct {
	Result    []*ScriptCoverage `json:"result"`
	Timestamp float64           `json:"timestamp"`
}

type TakeTypeProfileVal struct {
	Result []*ScriptTypeProfile `json:"result"`
}

type GetCountersVal struct {
	Result []*CounterInfo `json:"result"`
}

type GetRuntimeCallStatsVal struct {
	Result []*RuntimeCallCounterInfo `json:"result"`
}
