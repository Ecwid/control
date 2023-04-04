package serviceworker

import (
	"github.com/ecwid/control/protocol"
)

/*
 */
func DeliverPushMessage(c protocol.Caller, args DeliverPushMessageArgs) error {
	return c.Call("ServiceWorker.deliverPushMessage", args, nil)
}

/*
 */
func Disable(c protocol.Caller) error {
	return c.Call("ServiceWorker.disable", nil, nil)
}

/*
 */
func DispatchSyncEvent(c protocol.Caller, args DispatchSyncEventArgs) error {
	return c.Call("ServiceWorker.dispatchSyncEvent", args, nil)
}

/*
 */
func DispatchPeriodicSyncEvent(c protocol.Caller, args DispatchPeriodicSyncEventArgs) error {
	return c.Call("ServiceWorker.dispatchPeriodicSyncEvent", args, nil)
}

/*
 */
func Enable(c protocol.Caller) error {
	return c.Call("ServiceWorker.enable", nil, nil)
}

/*
 */
func InspectWorker(c protocol.Caller, args InspectWorkerArgs) error {
	return c.Call("ServiceWorker.inspectWorker", args, nil)
}

/*
 */
func SetForceUpdateOnPageLoad(c protocol.Caller, args SetForceUpdateOnPageLoadArgs) error {
	return c.Call("ServiceWorker.setForceUpdateOnPageLoad", args, nil)
}

/*
 */
func SkipWaiting(c protocol.Caller, args SkipWaitingArgs) error {
	return c.Call("ServiceWorker.skipWaiting", args, nil)
}

/*
 */
func StartWorker(c protocol.Caller, args StartWorkerArgs) error {
	return c.Call("ServiceWorker.startWorker", args, nil)
}

/*
 */
func StopAllWorkers(c protocol.Caller) error {
	return c.Call("ServiceWorker.stopAllWorkers", nil, nil)
}

/*
 */
func StopWorker(c protocol.Caller, args StopWorkerArgs) error {
	return c.Call("ServiceWorker.stopWorker", args, nil)
}

/*
 */
func Unregister(c protocol.Caller, args UnregisterArgs) error {
	return c.Call("ServiceWorker.unregister", args, nil)
}

/*
 */
func UpdateRegistration(c protocol.Caller, args UpdateRegistrationArgs) error {
	return c.Call("ServiceWorker.updateRegistration", args, nil)
}
