package pwa

import (
	"github.com/ecwid/control/protocol"
)

/*
Returns the following OS state for the given manifest id.
*/
func GetOsAppState(c protocol.Caller, args GetOsAppStateArgs) (*GetOsAppStateVal, error) {
	var val = &GetOsAppStateVal{}
	return val, c.Call("PWA.getOsAppState", args, val)
}

/*
	Installs the given manifest identity, optionally using the given installUrlOrBundleUrl

IWA-specific install description:
manifestId corresponds to isolated-app:// + web_package::SignedWebBundleId

File installation mode:
The installUrlOrBundleUrl can be either file:// or http(s):// pointing
to a signed web bundle (.swbn). In this case SignedWebBundleId must correspond to
The .swbn file's signing key.

Dev proxy installation mode:
installUrlOrBundleUrl must be http(s):// that serves dev mode IWA.
web_package::SignedWebBundleId must be of type dev proxy.

The advantage of dev proxy mode is that all changes to IWA
automatically will be reflected in the running app without
reinstallation.

To generate bundle id for proxy mode:
 1. Generate 32 random bytes.
 2. Add a specific suffix at the end following the documentation
    https://github.com/WICG/isolated-web-apps/blob/main/Scheme.md#suffix
 3. Encode the entire sequence using Base32 without padding.

If Chrome is not in IWA dev
mode, the installation will fail, regardless of the state of the allowlist.
*/
func Install(c protocol.Caller, args InstallArgs) error {
	return c.Call("PWA.install", args, nil)
}

/*
Uninstalls the given manifest_id and closes any opened app windows.
*/
func Uninstall(c protocol.Caller, args UninstallArgs) error {
	return c.Call("PWA.uninstall", args, nil)
}

/*
	Launches the installed web app, or an url in the same web app instead of the

default start url if it is provided. Returns a page Target.TargetID which
can be used to attach to via Target.attachToTarget or similar APIs.
*/
func Launch(c protocol.Caller, args LaunchArgs) (*LaunchVal, error) {
	var val = &LaunchVal{}
	return val, c.Call("PWA.launch", args, val)
}

/*
	Opens one or more local files from an installed web app identified by its

manifestId. The web app needs to have file handlers registered to process
the files. The API returns one or more page Target.TargetIDs which can be
used to attach to via Target.attachToTarget or similar APIs.
If some files in the parameters cannot be handled by the web app, they will
be ignored. If none of the files can be handled, this API returns an error.
If no files are provided as the parameter, this API also returns an error.

According to the definition of the file handlers in the manifest file, one
Target.TargetID may represent a page handling one or more files. The order
of the returned Target.TargetIDs is not guaranteed.

TODO(crbug.com/339454034): Check the existences of the input files.
*/
func LaunchFilesInApp(c protocol.Caller, args LaunchFilesInAppArgs) (*LaunchFilesInAppVal, error) {
	var val = &LaunchFilesInAppVal{}
	return val, c.Call("PWA.launchFilesInApp", args, val)
}

/*
	Opens the current page in its web app identified by the manifest id, needs

to be called on a page target. This function returns immediately without
waiting for the app to finish loading.
*/
func OpenCurrentPageInApp(c protocol.Caller, args OpenCurrentPageInAppArgs) error {
	return c.Call("PWA.openCurrentPageInApp", args, nil)
}

/*
	Changes user settings of the web app identified by its manifestId. If the

app was not installed, this command returns an error. Unset parameters will
be ignored; unrecognized values will cause an error.

Unlike the ones defined in the manifest files of the web apps, these
settings are provided by the browser and controlled by the users, they
impact the way the browser handling the web apps.

See the comment of each parameter.
*/
func ChangeAppUserSettings(c protocol.Caller, args ChangeAppUserSettingsArgs) error {
	return c.Call("PWA.changeAppUserSettings", args, nil)
}
