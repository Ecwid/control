package memory

import (
	"github.com/ecwid/control/protocol"
)

/*
Retruns current DOM object counters.
*/
func GetDOMCounters(c protocol.Caller) (*GetDOMCountersVal, error) {
	var val = &GetDOMCountersVal{}
	return val, c.Call("Memory.getDOMCounters", nil, val)
}

/*
Retruns DOM object counters after preparing renderer for leak detection.
*/
func GetDOMCountersForLeakDetection(c protocol.Caller) (*GetDOMCountersForLeakDetectionVal, error) {
	var val = &GetDOMCountersForLeakDetectionVal{}
	return val, c.Call("Memory.getDOMCountersForLeakDetection", nil, val)
}

/*
	Prepares for leak detection by terminating workers, stopping spellcheckers,

dropping non-essential internal caches, running garbage collections, etc.
*/
func PrepareForLeakDetection(c protocol.Caller) error {
	return c.Call("Memory.prepareForLeakDetection", nil, nil)
}

/*
Simulate OomIntervention by purging V8 memory.
*/
func ForciblyPurgeJavaScriptMemory(c protocol.Caller) error {
	return c.Call("Memory.forciblyPurgeJavaScriptMemory", nil, nil)
}

/*
Enable/disable suppressing memory pressure notifications in all processes.
*/
func SetPressureNotificationsSuppressed(c protocol.Caller, args SetPressureNotificationsSuppressedArgs) error {
	return c.Call("Memory.setPressureNotificationsSuppressed", args, nil)
}

/*
Simulate a memory pressure notification in all processes.
*/
func SimulatePressureNotification(c protocol.Caller, args SimulatePressureNotificationArgs) error {
	return c.Call("Memory.simulatePressureNotification", args, nil)
}

/*
Start collecting native memory profile.
*/
func StartSampling(c protocol.Caller, args StartSamplingArgs) error {
	return c.Call("Memory.startSampling", args, nil)
}

/*
Stop collecting native memory profile.
*/
func StopSampling(c protocol.Caller) error {
	return c.Call("Memory.stopSampling", nil, nil)
}

/*
	Retrieve native memory allocations profile

collected since renderer process startup.
*/
func GetAllTimeSamplingProfile(c protocol.Caller) (*GetAllTimeSamplingProfileVal, error) {
	var val = &GetAllTimeSamplingProfileVal{}
	return val, c.Call("Memory.getAllTimeSamplingProfile", nil, val)
}

/*
	Retrieve native memory allocations profile

collected since browser process startup.
*/
func GetBrowserSamplingProfile(c protocol.Caller) (*GetBrowserSamplingProfileVal, error) {
	var val = &GetBrowserSamplingProfileVal{}
	return val, c.Call("Memory.getBrowserSamplingProfile", nil, val)
}

/*
	Retrieve native memory allocations profile collected since last

`startSampling` call.
*/
func GetSamplingProfile(c protocol.Caller) (*GetSamplingProfileVal, error) {
	var val = &GetSamplingProfileVal{}
	return val, c.Call("Memory.getSamplingProfile", nil, val)
}
