package profiler

import (
	"github.com/ecwid/control/protocol"
)

/*

 */
func Disable(c protocol.Caller) error {
	return c.Call("Profiler.disable", nil, nil)
}

/*

 */
func Enable(c protocol.Caller) error {
	return c.Call("Profiler.enable", nil, nil)
}

/*
	Collect coverage data for the current isolate. The coverage data may be incomplete due to
garbage collection.
*/
func GetBestEffortCoverage(c protocol.Caller) (*GetBestEffortCoverageVal, error) {
	var val = &GetBestEffortCoverageVal{}
	return val, c.Call("Profiler.getBestEffortCoverage", nil, val)
}

/*
	Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
*/
func SetSamplingInterval(c protocol.Caller, args SetSamplingIntervalArgs) error {
	return c.Call("Profiler.setSamplingInterval", args, nil)
}

/*

 */
func Start(c protocol.Caller) error {
	return c.Call("Profiler.start", nil, nil)
}

/*
	Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code
coverage may be incomplete. Enabling prevents running optimized code and resets execution
counters.
*/
func StartPreciseCoverage(c protocol.Caller, args StartPreciseCoverageArgs) (*StartPreciseCoverageVal, error) {
	var val = &StartPreciseCoverageVal{}
	return val, c.Call("Profiler.startPreciseCoverage", args, val)
}

/*
	Enable type profile.
*/
func StartTypeProfile(c protocol.Caller) error {
	return c.Call("Profiler.startTypeProfile", nil, nil)
}

/*

 */
func Stop(c protocol.Caller) (*StopVal, error) {
	var val = &StopVal{}
	return val, c.Call("Profiler.stop", nil, val)
}

/*
	Disable precise code coverage. Disabling releases unnecessary execution count records and allows
executing optimized code.
*/
func StopPreciseCoverage(c protocol.Caller) error {
	return c.Call("Profiler.stopPreciseCoverage", nil, nil)
}

/*
	Disable type profile. Disabling releases type profile data collected so far.
*/
func StopTypeProfile(c protocol.Caller) error {
	return c.Call("Profiler.stopTypeProfile", nil, nil)
}

/*
	Collect coverage data for the current isolate, and resets execution counters. Precise code
coverage needs to have started.
*/
func TakePreciseCoverage(c protocol.Caller) (*TakePreciseCoverageVal, error) {
	var val = &TakePreciseCoverageVal{}
	return val, c.Call("Profiler.takePreciseCoverage", nil, val)
}

/*
	Collect type profile.
*/
func TakeTypeProfile(c protocol.Caller) (*TakeTypeProfileVal, error) {
	var val = &TakeTypeProfileVal{}
	return val, c.Call("Profiler.takeTypeProfile", nil, val)
}

/*
	Enable counters collection.
*/
func EnableCounters(c protocol.Caller) error {
	return c.Call("Profiler.enableCounters", nil, nil)
}

/*
	Disable counters collection.
*/
func DisableCounters(c protocol.Caller) error {
	return c.Call("Profiler.disableCounters", nil, nil)
}

/*
	Retrieve counters.
*/
func GetCounters(c protocol.Caller) (*GetCountersVal, error) {
	var val = &GetCountersVal{}
	return val, c.Call("Profiler.getCounters", nil, val)
}

/*
	Enable run time call stats collection.
*/
func EnableRuntimeCallStats(c protocol.Caller) error {
	return c.Call("Profiler.enableRuntimeCallStats", nil, nil)
}

/*
	Disable run time call stats collection.
*/
func DisableRuntimeCallStats(c protocol.Caller) error {
	return c.Call("Profiler.disableRuntimeCallStats", nil, nil)
}

/*
	Retrieve run time call stats.
*/
func GetRuntimeCallStats(c protocol.Caller) (*GetRuntimeCallStatsVal, error) {
	var val = &GetRuntimeCallStatsVal{}
	return val, c.Call("Profiler.getRuntimeCallStats", nil, val)
}
