package target

import (
	"github.com/ecwid/control/protocol"
)

/*
	Activates (focuses) the target.
*/
func ActivateTarget(c protocol.Caller, args ActivateTargetArgs) error {
	return c.Call("Target.activateTarget", args, nil)
}

/*
	Attaches to the target with given id.
*/
func AttachToTarget(c protocol.Caller, args AttachToTargetArgs) (*AttachToTargetVal, error) {
	var val = &AttachToTargetVal{}
	return val, c.Call("Target.attachToTarget", args, val)
}

/*
	Attaches to the browser target, only uses flat sessionId mode.
*/
func AttachToBrowserTarget(c protocol.Caller) (*AttachToBrowserTargetVal, error) {
	var val = &AttachToBrowserTargetVal{}
	return val, c.Call("Target.attachToBrowserTarget", nil, val)
}

/*
	Closes the target. If the target is a page that gets closed too.
*/
func CloseTarget(c protocol.Caller, args CloseTargetArgs) error {
	return c.Call("Target.closeTarget", args, nil)
}

/*
	Inject object to the target's main frame that provides a communication
channel with browser target.

Injected object will be available as `window[bindingName]`.

The object has the follwing API:
- `binding.send(json)` - a method to send messages over the remote debugging protocol
- `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
*/
func ExposeDevToolsProtocol(c protocol.Caller, args ExposeDevToolsProtocolArgs) error {
	return c.Call("Target.exposeDevToolsProtocol", args, nil)
}

/*
	Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than
one.
*/
func CreateBrowserContext(c protocol.Caller, args CreateBrowserContextArgs) (*CreateBrowserContextVal, error) {
	var val = &CreateBrowserContextVal{}
	return val, c.Call("Target.createBrowserContext", args, val)
}

/*
	Returns all browser contexts created with `Target.createBrowserContext` method.
*/
func GetBrowserContexts(c protocol.Caller) (*GetBrowserContextsVal, error) {
	var val = &GetBrowserContextsVal{}
	return val, c.Call("Target.getBrowserContexts", nil, val)
}

/*
	Creates a new page.
*/
func CreateTarget(c protocol.Caller, args CreateTargetArgs) (*CreateTargetVal, error) {
	var val = &CreateTargetVal{}
	return val, c.Call("Target.createTarget", args, val)
}

/*
	Detaches session with given id.
*/
func DetachFromTarget(c protocol.Caller, args DetachFromTargetArgs) error {
	return c.Call("Target.detachFromTarget", args, nil)
}

/*
	Deletes a BrowserContext. All the belonging pages will be closed without calling their
beforeunload hooks.
*/
func DisposeBrowserContext(c protocol.Caller, args DisposeBrowserContextArgs) error {
	return c.Call("Target.disposeBrowserContext", args, nil)
}

/*
	Returns information about a target.
*/
func GetTargetInfo(c protocol.Caller, args GetTargetInfoArgs) (*GetTargetInfoVal, error) {
	var val = &GetTargetInfoVal{}
	return val, c.Call("Target.getTargetInfo", args, val)
}

/*
	Retrieves a list of available targets.
*/
func GetTargets(c protocol.Caller) (*GetTargetsVal, error) {
	var val = &GetTargetsVal{}
	return val, c.Call("Target.getTargets", nil, val)
}

/*
	Controls whether to automatically attach to new targets which are considered to be related to
this one. When turned on, attaches to all existing related targets as well. When turned off,
automatically detaches from all currently attached targets.
*/
func SetAutoAttach(c protocol.Caller, args SetAutoAttachArgs) error {
	return c.Call("Target.setAutoAttach", args, nil)
}

/*
	Controls whether to discover available targets and notify via
`targetCreated/targetInfoChanged/targetDestroyed` events.
*/
func SetDiscoverTargets(c protocol.Caller, args SetDiscoverTargetsArgs) error {
	return c.Call("Target.setDiscoverTargets", args, nil)
}

/*
	Enables target discovery for the specified locations, when `setDiscoverTargets` was set to
`true`.
*/
func SetRemoteLocations(c protocol.Caller, args SetRemoteLocationsArgs) error {
	return c.Call("Target.setRemoteLocations", args, nil)
}
