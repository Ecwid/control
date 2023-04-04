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
	Collect coverage data for the current isolate, and resets execution counters. Precise code

coverage needs to have started.
*/
func TakePreciseCoverage(c protocol.Caller) (*TakePreciseCoverageVal, error) {
	var val = &TakePreciseCoverageVal{}
	return val, c.Call("Profiler.takePreciseCoverage", nil, val)
}
