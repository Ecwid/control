package network

import (
	"github.com/ecwid/control/protocol"
)

/*
	Sets a list of content encodings that will be accepted. Empty list means no encoding is accepted.
*/
func SetAcceptedEncodings(c protocol.Caller, args SetAcceptedEncodingsArgs) error {
	return c.Call("Network.setAcceptedEncodings", args, nil)
}

/*
	Clears accepted encodings set by setAcceptedEncodings
*/
func ClearAcceptedEncodingsOverride(c protocol.Caller) error {
	return c.Call("Network.clearAcceptedEncodingsOverride", nil, nil)
}

/*
	Clears browser cache.
*/
func ClearBrowserCache(c protocol.Caller) error {
	return c.Call("Network.clearBrowserCache", nil, nil)
}

/*
	Clears browser cookies.
*/
func ClearBrowserCookies(c protocol.Caller) error {
	return c.Call("Network.clearBrowserCookies", nil, nil)
}

/*
	Deletes browser cookies with matching name and url or domain/path pair.
*/
func DeleteCookies(c protocol.Caller, args DeleteCookiesArgs) error {
	return c.Call("Network.deleteCookies", args, nil)
}

/*
	Disables network tracking, prevents network events from being sent to the client.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Network.disable", nil, nil)
}

/*
	Activates emulation of network conditions.
*/
func EmulateNetworkConditions(c protocol.Caller, args EmulateNetworkConditionsArgs) error {
	return c.Call("Network.emulateNetworkConditions", args, nil)
}

/*
	Enables network tracking, network events will now be delivered to the client.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("Network.enable", args, nil)
}

/*
	Returns all browser cookies. Depending on the backend support, will return detailed cookie
information in the `cookies` field.
*/
func GetAllCookies(c protocol.Caller) (*GetAllCookiesVal, error) {
	var val = &GetAllCookiesVal{}
	return val, c.Call("Network.getAllCookies", nil, val)
}

/*
	Returns the DER-encoded certificate.
*/
func GetCertificate(c protocol.Caller, args GetCertificateArgs) (*GetCertificateVal, error) {
	var val = &GetCertificateVal{}
	return val, c.Call("Network.getCertificate", args, val)
}

/*
	Returns all browser cookies for the current URL. Depending on the backend support, will return
detailed cookie information in the `cookies` field.
*/
func GetCookies(c protocol.Caller, args GetCookiesArgs) (*GetCookiesVal, error) {
	var val = &GetCookiesVal{}
	return val, c.Call("Network.getCookies", args, val)
}

/*
	Returns content served for the given request.
*/
func GetResponseBody(c protocol.Caller, args GetResponseBodyArgs) (*GetResponseBodyVal, error) {
	var val = &GetResponseBodyVal{}
	return val, c.Call("Network.getResponseBody", args, val)
}

/*
	Returns post data sent with the request. Returns an error when no data was sent with the request.
*/
func GetRequestPostData(c protocol.Caller, args GetRequestPostDataArgs) (*GetRequestPostDataVal, error) {
	var val = &GetRequestPostDataVal{}
	return val, c.Call("Network.getRequestPostData", args, val)
}

/*
	Returns content served for the given currently intercepted request.
*/
func GetResponseBodyForInterception(c protocol.Caller, args GetResponseBodyForInterceptionArgs) (*GetResponseBodyForInterceptionVal, error) {
	var val = &GetResponseBodyForInterceptionVal{}
	return val, c.Call("Network.getResponseBodyForInterception", args, val)
}

/*
	Returns a handle to the stream representing the response body. Note that after this command,
the intercepted request can't be continued as is -- you either need to cancel it or to provide
the response body. The stream only supports sequential read, IO.read will fail if the position
is specified.
*/
func TakeResponseBodyForInterceptionAsStream(c protocol.Caller, args TakeResponseBodyForInterceptionAsStreamArgs) (*TakeResponseBodyForInterceptionAsStreamVal, error) {
	var val = &TakeResponseBodyForInterceptionAsStreamVal{}
	return val, c.Call("Network.takeResponseBodyForInterceptionAsStream", args, val)
}

/*
	This method sends a new XMLHttpRequest which is identical to the original one. The following
parameters should be identical: method, url, async, request body, extra headers, withCredentials
attribute, user, password.
*/
func ReplayXHR(c protocol.Caller, args ReplayXHRArgs) error {
	return c.Call("Network.replayXHR", args, nil)
}

/*
	Searches for given string in response content.
*/
func SearchInResponseBody(c protocol.Caller, args SearchInResponseBodyArgs) (*SearchInResponseBodyVal, error) {
	var val = &SearchInResponseBodyVal{}
	return val, c.Call("Network.searchInResponseBody", args, val)
}

/*
	Blocks URLs from loading.
*/
func SetBlockedURLs(c protocol.Caller, args SetBlockedURLsArgs) error {
	return c.Call("Network.setBlockedURLs", args, nil)
}

/*
	Toggles ignoring of service worker for each request.
*/
func SetBypassServiceWorker(c protocol.Caller, args SetBypassServiceWorkerArgs) error {
	return c.Call("Network.setBypassServiceWorker", args, nil)
}

/*
	Toggles ignoring cache for each request. If `true`, cache will not be used.
*/
func SetCacheDisabled(c protocol.Caller, args SetCacheDisabledArgs) error {
	return c.Call("Network.setCacheDisabled", args, nil)
}

/*
	Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
*/
func SetCookie(c protocol.Caller, args SetCookieArgs) error {
	return c.Call("Network.setCookie", args, nil)
}

/*
	Sets given cookies.
*/
func SetCookies(c protocol.Caller, args SetCookiesArgs) error {
	return c.Call("Network.setCookies", args, nil)
}

/*
	For testing.
*/
func SetDataSizeLimitsForTest(c protocol.Caller, args SetDataSizeLimitsForTestArgs) error {
	return c.Call("Network.setDataSizeLimitsForTest", args, nil)
}

/*
	Specifies whether to always send extra HTTP headers with the requests from this page.
*/
func SetExtraHTTPHeaders(c protocol.Caller, args SetExtraHTTPHeadersArgs) error {
	return c.Call("Network.setExtraHTTPHeaders", args, nil)
}

/*
	Specifies whether to attach a page script stack id in requests
*/
func SetAttachDebugStack(c protocol.Caller, args SetAttachDebugStackArgs) error {
	return c.Call("Network.setAttachDebugStack", args, nil)
}

/*
	Allows overriding user agent with the given string.
*/
func SetUserAgentOverride(c protocol.Caller, args SetUserAgentOverrideArgs) error {
	return c.Call("Network.setUserAgentOverride", args, nil)
}

/*
	Returns information about the COEP/COOP isolation status.
*/
func GetSecurityIsolationStatus(c protocol.Caller, args GetSecurityIsolationStatusArgs) (*GetSecurityIsolationStatusVal, error) {
	var val = &GetSecurityIsolationStatusVal{}
	return val, c.Call("Network.getSecurityIsolationStatus", args, val)
}

/*
	Fetches the resource and returns the content.
*/
func LoadNetworkResource(c protocol.Caller, args LoadNetworkResourceArgs) (*LoadNetworkResourceVal, error) {
	var val = &LoadNetworkResourceVal{}
	return val, c.Call("Network.loadNetworkResource", args, val)
}
