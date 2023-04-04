package heapprofiler

import (
	"github.com/ecwid/control/protocol"
)

/*
	Enables console to refer to the node with given id via $x (see Command Line API for more details

$x functions).
*/
func AddInspectedHeapObject(c protocol.Caller, args AddInspectedHeapObjectArgs) error {
	return c.Call("HeapProfiler.addInspectedHeapObject", args, nil)
}

/*
 */
func CollectGarbage(c protocol.Caller) error {
	return c.Call("HeapProfiler.collectGarbage", nil, nil)
}

/*
 */
func Disable(c protocol.Caller) error {
	return c.Call("HeapProfiler.disable", nil, nil)
}

/*
 */
func Enable(c protocol.Caller) error {
	return c.Call("HeapProfiler.enable", nil, nil)
}

/*
 */
func GetHeapObjectId(c protocol.Caller, args GetHeapObjectIdArgs) (*GetHeapObjectIdVal, error) {
	var val = &GetHeapObjectIdVal{}
	return val, c.Call("HeapProfiler.getHeapObjectId", args, val)
}

/*
 */
func GetObjectByHeapObjectId(c protocol.Caller, args GetObjectByHeapObjectIdArgs) (*GetObjectByHeapObjectIdVal, error) {
	var val = &GetObjectByHeapObjectIdVal{}
	return val, c.Call("HeapProfiler.getObjectByHeapObjectId", args, val)
}

/*
 */
func GetSamplingProfile(c protocol.Caller) (*GetSamplingProfileVal, error) {
	var val = &GetSamplingProfileVal{}
	return val, c.Call("HeapProfiler.getSamplingProfile", nil, val)
}

/*
 */
func StartSampling(c protocol.Caller, args StartSamplingArgs) error {
	return c.Call("HeapProfiler.startSampling", args, nil)
}

/*
 */
func StartTrackingHeapObjects(c protocol.Caller, args StartTrackingHeapObjectsArgs) error {
	return c.Call("HeapProfiler.startTrackingHeapObjects", args, nil)
}

/*
 */
func StopSampling(c protocol.Caller) (*StopSamplingVal, error) {
	var val = &StopSamplingVal{}
	return val, c.Call("HeapProfiler.stopSampling", nil, val)
}

/*
 */
func StopTrackingHeapObjects(c protocol.Caller, args StopTrackingHeapObjectsArgs) error {
	return c.Call("HeapProfiler.stopTrackingHeapObjects", args, nil)
}

/*
 */
func TakeHeapSnapshot(c protocol.Caller, args TakeHeapSnapshotArgs) error {
	return c.Call("HeapProfiler.takeHeapSnapshot", args, nil)
}
