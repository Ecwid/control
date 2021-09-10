package tracing

import (
	"github.com/ecwid/control/protocol"
)

/*
	Stop trace events collection.
*/
func End(c protocol.Caller) error {
	return c.Call("Tracing.end", nil, nil)
}

/*
	Gets supported tracing categories.
*/
func GetCategories(c protocol.Caller) (*GetCategoriesVal, error) {
	var val = &GetCategoriesVal{}
	return val, c.Call("Tracing.getCategories", nil, val)
}

/*
	Record a clock sync marker in the trace.
*/
func RecordClockSyncMarker(c protocol.Caller, args RecordClockSyncMarkerArgs) error {
	return c.Call("Tracing.recordClockSyncMarker", args, nil)
}

/*
	Request a global memory dump.
*/
func RequestMemoryDump(c protocol.Caller, args RequestMemoryDumpArgs) (*RequestMemoryDumpVal, error) {
	var val = &RequestMemoryDumpVal{}
	return val, c.Call("Tracing.requestMemoryDump", args, val)
}

/*
	Start trace events collection.
*/
func Start(c protocol.Caller, args StartArgs) error {
	return c.Call("Tracing.start", args, nil)
}
