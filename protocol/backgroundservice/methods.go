package backgroundservice

import (
	"github.com/ecwid/control/protocol"
)

/*
Enables event updates for the service.
*/
func StartObserving(c protocol.Caller, args StartObservingArgs) error {
	return c.Call("BackgroundService.startObserving", args, nil)
}

/*
Disables event updates for the service.
*/
func StopObserving(c protocol.Caller, args StopObservingArgs) error {
	return c.Call("BackgroundService.stopObserving", args, nil)
}

/*
Set the recording state for the service.
*/
func SetRecording(c protocol.Caller, args SetRecordingArgs) error {
	return c.Call("BackgroundService.setRecording", args, nil)
}

/*
Clears all stored data for the service.
*/
func ClearEvents(c protocol.Caller, args ClearEventsArgs) error {
	return c.Call("BackgroundService.clearEvents", args, nil)
}
