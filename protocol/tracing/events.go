package tracing

import (
	"github.com/ecwid/control/protocol/io"
)

/*
 */
type BufferUsage struct {
	PercentFull float64 `json:"percentFull,omitempty"`
	EventCount  float64 `json:"eventCount,omitempty"`
	Value       float64 `json:"value,omitempty"`
}

/*
	Contains a bucket of collected trace events. When tracing is stopped collected events will be

sent as a sequence of dataCollected events followed by tracingComplete event.
*/
type DataCollected struct {
	Value []interface{} `json:"value"`
}

/*
	Signals that tracing is stopped and there is no trace buffers pending flush, all data were

delivered via dataCollected events.
*/
type TracingComplete struct {
	DataLossOccurred  bool              `json:"dataLossOccurred"`
	Stream            io.StreamHandle   `json:"stream,omitempty"`
	TraceFormat       StreamFormat      `json:"traceFormat,omitempty"`
	StreamCompression StreamCompression `json:"streamCompression,omitempty"`
}
