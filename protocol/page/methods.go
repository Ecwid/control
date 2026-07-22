package page

import (
	"github.com/ecwid/control/protocol"
)

/*
Evaluates given script in every frame upon creation (before loading frame's scripts).
*/
func AddScriptToEvaluateOnNewDocument(c protocol.Caller, args AddScriptToEvaluateOnNewDocumentArgs) (*AddScriptToEvaluateOnNewDocumentVal, error) {
	var val = &AddScriptToEvaluateOnNewDocumentVal{}
	return val, c.Call("Page.addScriptToEvaluateOnNewDocument", args, val)
}

/*
Brings page to front (activates tab).
*/
func BringToFront(c protocol.Caller) error {
	return c.Call("Page.bringToFront", nil, nil)
}

/*
Capture page screenshot.
*/
func CaptureScreenshot(c protocol.Caller, args CaptureScreenshotArgs) (*CaptureScreenshotVal, error) {
	var val = &CaptureScreenshotVal{}
	return val, c.Call("Page.captureScreenshot", args, val)
}

/*
	Returns a snapshot of the page as a string. For MHTML format, the serialization includes

iframes, shadow DOM, external resources, and element-inline styles.
*/
func CaptureSnapshot(c protocol.Caller, args CaptureSnapshotArgs) (*CaptureSnapshotVal, error) {
	var val = &CaptureSnapshotVal{}
	return val, c.Call("Page.captureSnapshot", args, val)
}

/*
Creates an isolated world for the given frame.
*/
func CreateIsolatedWorld(c protocol.Caller, args CreateIsolatedWorldArgs) (*CreateIsolatedWorldVal, error) {
	var val = &CreateIsolatedWorldVal{}
	return val, c.Call("Page.createIsolatedWorld", args, val)
}

/*
Disables page domain notifications.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Page.disable", nil, nil)
}

/*
Enables page domain notifications.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("Page.enable", args, nil)
}

/*
		Gets the processed manifest for this current document.
	  This API always waits for the manifest to be loaded.
	  If manifestId is provided, and it does not match the manifest of the
	    current document, this API errors out.
	  If there is not a loaded page, this API errors out immediately.
*/
func GetAppManifest(c protocol.Caller, args GetAppManifestArgs) (*GetAppManifestVal, error) {
	var val = &GetAppManifestVal{}
	return val, c.Call("Page.getAppManifest", args, val)
}

/*
 */
func GetInstallabilityErrors(c protocol.Caller) (*GetInstallabilityErrorsVal, error) {
	var val = &GetInstallabilityErrorsVal{}
	return val, c.Call("Page.getInstallabilityErrors", nil, val)
}

/*
	Returns the unique (PWA) app id.

Only returns values if the feature flag 'WebAppEnableManifestId' is enabled
*/
func GetAppId(c protocol.Caller) (*GetAppIdVal, error) {
	var val = &GetAppIdVal{}
	return val, c.Call("Page.getAppId", nil, val)
}

/*
 */
func GetAdScriptAncestry(c protocol.Caller, args GetAdScriptAncestryArgs) (*GetAdScriptAncestryVal, error) {
	var val = &GetAdScriptAncestryVal{}
	return val, c.Call("Page.getAdScriptAncestry", args, val)
}

/*
Returns present frame tree structure.
*/
func GetFrameTree(c protocol.Caller) (*GetFrameTreeVal, error) {
	var val = &GetFrameTreeVal{}
	return val, c.Call("Page.getFrameTree", nil, val)
}

/*
Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
*/
func GetLayoutMetrics(c protocol.Caller) (*GetLayoutMetricsVal, error) {
	var val = &GetLayoutMetricsVal{}
	return val, c.Call("Page.getLayoutMetrics", nil, val)
}

/*
Returns navigation history for the current page.
*/
func GetNavigationHistory(c protocol.Caller) (*GetNavigationHistoryVal, error) {
	var val = &GetNavigationHistoryVal{}
	return val, c.Call("Page.getNavigationHistory", nil, val)
}

/*
Resets navigation history for the current page.
*/
func ResetNavigationHistory(c protocol.Caller) error {
	return c.Call("Page.resetNavigationHistory", nil, nil)
}

/*
Returns content of the given resource.
*/
func GetResourceContent(c protocol.Caller, args GetResourceContentArgs) (*GetResourceContentVal, error) {
	var val = &GetResourceContentVal{}
	return val, c.Call("Page.getResourceContent", args, val)
}

/*
Returns present frame / resource tree structure.
*/
func GetResourceTree(c protocol.Caller) (*GetResourceTreeVal, error) {
	var val = &GetResourceTreeVal{}
	return val, c.Call("Page.getResourceTree", nil, val)
}

/*
Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
*/
func HandleJavaScriptDialog(c protocol.Caller, args HandleJavaScriptDialogArgs) error {
	return c.Call("Page.handleJavaScriptDialog", args, nil)
}

/*
Navigates current page to the given URL.
*/
func Navigate(c protocol.Caller, args NavigateArgs) (*NavigateVal, error) {
	var val = &NavigateVal{}
	return val, c.Call("Page.navigate", args, val)
}

/*
Navigates current page to the given history entry.
*/
func NavigateToHistoryEntry(c protocol.Caller, args NavigateToHistoryEntryArgs) error {
	return c.Call("Page.navigateToHistoryEntry", args, nil)
}

/*
Print page as PDF.
*/
func PrintToPDF(c protocol.Caller, args PrintToPDFArgs) (*PrintToPDFVal, error) {
	var val = &PrintToPDFVal{}
	return val, c.Call("Page.printToPDF", args, val)
}

/*
Reloads given page optionally ignoring the cache.
*/
func Reload(c protocol.Caller, args ReloadArgs) error {
	return c.Call("Page.reload", args, nil)
}

/*
Removes given script from the list.
*/
func RemoveScriptToEvaluateOnNewDocument(c protocol.Caller, args RemoveScriptToEvaluateOnNewDocumentArgs) error {
	return c.Call("Page.removeScriptToEvaluateOnNewDocument", args, nil)
}

/*
Acknowledges that a screencast frame has been received by the frontend.
*/
func ScreencastFrameAck(c protocol.Caller, args ScreencastFrameAckArgs) error {
	return c.Call("Page.screencastFrameAck", args, nil)
}

/*
Searches for given string in resource content.
*/
func SearchInResource(c protocol.Caller, args SearchInResourceArgs) (*SearchInResourceVal, error) {
	var val = &SearchInResourceVal{}
	return val, c.Call("Page.searchInResource", args, val)
}

/*
Enable Chrome's experimental ad filter on all sites.
*/
func SetAdBlockingEnabled(c protocol.Caller, args SetAdBlockingEnabledArgs) error {
	return c.Call("Page.setAdBlockingEnabled", args, nil)
}

/*
Enable page Content Security Policy by-passing.
*/
func SetBypassCSP(c protocol.Caller, args SetBypassCSPArgs) error {
	return c.Call("Page.setBypassCSP", args, nil)
}

/*
Get Permissions Policy state on given frame.
*/
func GetPermissionsPolicyState(c protocol.Caller, args GetPermissionsPolicyStateArgs) (*GetPermissionsPolicyStateVal, error) {
	var val = &GetPermissionsPolicyStateVal{}
	return val, c.Call("Page.getPermissionsPolicyState", args, val)
}

/*
Get Origin Trials on given frame.
*/
func GetOriginTrials(c protocol.Caller, args GetOriginTrialsArgs) (*GetOriginTrialsVal, error) {
	var val = &GetOriginTrialsVal{}
	return val, c.Call("Page.getOriginTrials", args, val)
}

/*
Set generic font families.
*/
func SetFontFamilies(c protocol.Caller, args SetFontFamiliesArgs) error {
	return c.Call("Page.setFontFamilies", args, nil)
}

/*
Set default font sizes.
*/
func SetFontSizes(c protocol.Caller, args SetFontSizesArgs) error {
	return c.Call("Page.setFontSizes", args, nil)
}

/*
Sets given markup as the document's HTML.
*/
func SetDocumentContent(c protocol.Caller, args SetDocumentContentArgs) error {
	return c.Call("Page.setDocumentContent", args, nil)
}

/*
Controls whether page will emit lifecycle events.
*/
func SetLifecycleEventsEnabled(c protocol.Caller, args SetLifecycleEventsEnabledArgs) error {
	return c.Call("Page.setLifecycleEventsEnabled", args, nil)
}

/*
Starts sending each frame using the `screencastFrame` event.
*/
func StartScreencast(c protocol.Caller, args StartScreencastArgs) error {
	return c.Call("Page.startScreencast", args, nil)
}

/*
Force the page stop all navigations and pending resource fetches.
*/
func StopLoading(c protocol.Caller) error {
	return c.Call("Page.stopLoading", nil, nil)
}

/*
Crashes renderer on the IO thread, generates minidumps.
*/
func Crash(c protocol.Caller) error {
	return c.Call("Page.crash", nil, nil)
}

/*
Tries to close page, running its beforeunload hooks, if any.
*/
func Close(c protocol.Caller) error {
	return c.Call("Page.close", nil, nil)
}

/*
	Tries to update the web lifecycle state of the page.

It will transition the page to the given state according to:
https://github.com/WICG/web-lifecycle/
*/
func SetWebLifecycleState(c protocol.Caller, args SetWebLifecycleStateArgs) error {
	return c.Call("Page.setWebLifecycleState", args, nil)
}

/*
Stops sending each frame in the `screencastFrame`.
*/
func StopScreencast(c protocol.Caller) error {
	return c.Call("Page.stopScreencast", nil, nil)
}

/*
	Requests backend to produce compilation cache for the specified scripts.

`scripts` are appended to the list of scripts for which the cache
would be produced. The list may be reset during page navigation.
When script with a matching URL is encountered, the cache is optionally
produced upon backend discretion, based on internal heuristics.
See also: `Page.compilationCacheProduced`.
*/
func ProduceCompilationCache(c protocol.Caller, args ProduceCompilationCacheArgs) error {
	return c.Call("Page.produceCompilationCache", args, nil)
}

/*
	Seeds compilation cache for given url. Compilation cache does not survive

cross-process navigation.
*/
func AddCompilationCache(c protocol.Caller, args AddCompilationCacheArgs) error {
	return c.Call("Page.addCompilationCache", args, nil)
}

/*
Clears seeded compilation cache.
*/
func ClearCompilationCache(c protocol.Caller) error {
	return c.Call("Page.clearCompilationCache", nil, nil)
}

/*
	Sets the Secure Payment Confirmation transaction mode.

https://w3c.github.io/secure-payment-confirmation/#sctn-automation-set-spc-transaction-mode
*/
func SetSPCTransactionMode(c protocol.Caller, args SetSPCTransactionModeArgs) error {
	return c.Call("Page.setSPCTransactionMode", args, nil)
}

/*
	Extensions for Custom Handlers API:

https://html.spec.whatwg.org/multipage/system-state.html#rph-automation
*/
func SetRPHRegistrationMode(c protocol.Caller, args SetRPHRegistrationModeArgs) error {
	return c.Call("Page.setRPHRegistrationMode", args, nil)
}

/*
Generates a report for testing.
*/
func GenerateTestReport(c protocol.Caller, args GenerateTestReportArgs) error {
	return c.Call("Page.generateTestReport", args, nil)
}

/*
Pauses page execution. Can be resumed using generic Runtime.runIfWaitingForDebugger.
*/
func WaitForDebugger(c protocol.Caller) error {
	return c.Call("Page.waitForDebugger", nil, nil)
}

/*
	Intercept file chooser requests and transfer control to protocol clients.

When file chooser interception is enabled, native file chooser dialog is not shown.
Instead, a protocol event `Page.fileChooserOpened` is emitted.
*/
func SetInterceptFileChooserDialog(c protocol.Caller, args SetInterceptFileChooserDialogArgs) error {
	return c.Call("Page.setInterceptFileChooserDialog", args, nil)
}

/*
	Enable/disable prerendering manually.

This command is a short-term solution for https://crbug.com/1440085.
See https://docs.google.com/document/d/12HVmFxYj5Jc-eJr5OmWsa2bqTJsbgGLKI6ZIyx0_wpA
for more details.

TODO(https://crbug.com/1440085): Remove this once Puppeteer supports tab targets.
*/
func SetPrerenderingAllowed(c protocol.Caller, args SetPrerenderingAllowedArgs) error {
	return c.Call("Page.setPrerenderingAllowed", args, nil)
}

/*
	Get the annotated page content for the main frame.

This is an experimental command that is subject to change.
*/
func GetAnnotatedPageContent(c protocol.Caller, args GetAnnotatedPageContentArgs) (*GetAnnotatedPageContentVal, error) {
	var val = &GetAnnotatedPageContentVal{}
	return val, c.Call("Page.getAnnotatedPageContent", args, val)
}
