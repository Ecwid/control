package browser

import (
	"github.com/ecwid/control/protocol"
)

/*
Set permission settings for given embedding and embedded origins.
*/
func SetPermission(c protocol.Caller, args SetPermissionArgs) error {
	return c.Call("Browser.setPermission", args, nil)
}

/*
Reset all permission management for all origins.
*/
func ResetPermissions(c protocol.Caller, args ResetPermissionsArgs) error {
	return c.Call("Browser.resetPermissions", args, nil)
}

/*
Set the behavior when downloading a file.
*/
func SetDownloadBehavior(c protocol.Caller, args SetDownloadBehaviorArgs) error {
	return c.Call("Browser.setDownloadBehavior", args, nil)
}

/*
Cancel a download if in progress
*/
func CancelDownload(c protocol.Caller, args CancelDownloadArgs) error {
	return c.Call("Browser.cancelDownload", args, nil)
}

/*
Close browser gracefully.
*/
func Close(c protocol.Caller) error {
	return c.Call("Browser.close", nil, nil)
}

/*
Crashes browser on the main thread.
*/
func Crash(c protocol.Caller) error {
	return c.Call("Browser.crash", nil, nil)
}

/*
Crashes GPU process.
*/
func CrashGpuProcess(c protocol.Caller) error {
	return c.Call("Browser.crashGpuProcess", nil, nil)
}

/*
Returns version information.
*/
func GetVersion(c protocol.Caller) (*GetVersionVal, error) {
	var val = &GetVersionVal{}
	return val, c.Call("Browser.getVersion", nil, val)
}

/*
	Returns the command line switches for the browser process if, and only if

--enable-automation is on the commandline.
*/
func GetBrowserCommandLine(c protocol.Caller) (*GetBrowserCommandLineVal, error) {
	var val = &GetBrowserCommandLineVal{}
	return val, c.Call("Browser.getBrowserCommandLine", nil, val)
}

/*
Get Chrome histograms.
*/
func GetHistograms(c protocol.Caller, args GetHistogramsArgs) (*GetHistogramsVal, error) {
	var val = &GetHistogramsVal{}
	return val, c.Call("Browser.getHistograms", args, val)
}

/*
Get a Chrome histogram by name.
*/
func GetHistogram(c protocol.Caller, args GetHistogramArgs) (*GetHistogramVal, error) {
	var val = &GetHistogramVal{}
	return val, c.Call("Browser.getHistogram", args, val)
}

/*
Get position and size of the browser window.
*/
func GetWindowBounds(c protocol.Caller, args GetWindowBoundsArgs) (*GetWindowBoundsVal, error) {
	var val = &GetWindowBoundsVal{}
	return val, c.Call("Browser.getWindowBounds", args, val)
}

/*
Get the browser window that contains the devtools target.
*/
func GetWindowForTarget(c protocol.Caller, args GetWindowForTargetArgs) (*GetWindowForTargetVal, error) {
	var val = &GetWindowForTargetVal{}
	return val, c.Call("Browser.getWindowForTarget", args, val)
}

/*
Set position and/or size of the browser window.
*/
func SetWindowBounds(c protocol.Caller, args SetWindowBoundsArgs) error {
	return c.Call("Browser.setWindowBounds", args, nil)
}

/*
Set size of the browser contents resizing browser window as necessary.
*/
func SetContentsSize(c protocol.Caller, args SetContentsSizeArgs) error {
	return c.Call("Browser.setContentsSize", args, nil)
}

/*
Set dock tile details, platform-specific.
*/
func SetDockTile(c protocol.Caller, args SetDockTileArgs) error {
	return c.Call("Browser.setDockTile", args, nil)
}

/*
Invoke custom browser commands used by telemetry.
*/
func ExecuteBrowserCommand(c protocol.Caller, args ExecuteBrowserCommandArgs) error {
	return c.Call("Browser.executeBrowserCommand", args, nil)
}

/*
	Allows a site to use privacy sandbox features that require enrollment

without the site actually being enrolled. Only supported on page targets.
*/
func AddPrivacySandboxEnrollmentOverride(c protocol.Caller, args AddPrivacySandboxEnrollmentOverrideArgs) error {
	return c.Call("Browser.addPrivacySandboxEnrollmentOverride", args, nil)
}

/*
	Configures encryption keys used with a given privacy sandbox API to talk

to a trusted coordinator.  Since this is intended for test automation only,
coordinatorOrigin must be a .test domain. No existing coordinator
configuration for the origin may exist.
*/
func AddPrivacySandboxCoordinatorKeyConfig(c protocol.Caller, args AddPrivacySandboxCoordinatorKeyConfigArgs) error {
	return c.Call("Browser.addPrivacySandboxCoordinatorKeyConfig", args, nil)
}
