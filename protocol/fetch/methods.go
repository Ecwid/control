package fetch

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables the fetch domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Fetch.disable", nil, nil)
}

/*
	Enables issuing of requestPaused events. A request will be paused until client

calls one of failRequest, fulfillRequest or continueRequest/continueWithAuth.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("Fetch.enable", args, nil)
}

/*
Causes the request to fail with specified reason.
*/
func FailRequest(c protocol.Caller, args FailRequestArgs) error {
	return c.Call("Fetch.failRequest", args, nil)
}

/*
Provides response to the request.
*/
func FulfillRequest(c protocol.Caller, args FulfillRequestArgs) error {
	return c.Call("Fetch.fulfillRequest", args, nil)
}

/*
Continues the request, optionally modifying some of its parameters.
*/
func ContinueRequest(c protocol.Caller, args ContinueRequestArgs) error {
	return c.Call("Fetch.continueRequest", args, nil)
}

/*
Continues a request supplying authChallengeResponse following authRequired event.
*/
func ContinueWithAuth(c protocol.Caller, args ContinueWithAuthArgs) error {
	return c.Call("Fetch.continueWithAuth", args, nil)
}

/*
	Continues loading of the paused response, optionally modifying the

response headers. If either responseCode or headers are modified, all of them
must be present.
*/
func ContinueResponse(c protocol.Caller, args ContinueResponseArgs) error {
	return c.Call("Fetch.continueResponse", args, nil)
}

/*
	Causes the body of the response to be received from the server and

returned as a single string. May only be issued for a request that
is paused in the Response stage and is mutually exclusive with
takeResponseBodyForInterceptionAsStream. Calling other methods that
affect the request or disabling fetch domain before body is received
results in an undefined behavior.
*/
func GetResponseBody(c protocol.Caller, args GetResponseBodyArgs) (*GetResponseBodyVal, error) {
	var val = &GetResponseBodyVal{}
	return val, c.Call("Fetch.getResponseBody", args, val)
}

/*
	Returns a handle to the stream representing the response body.

The request must be paused in the HeadersReceived stage.
Note that after this command the request can't be continued
as is -- client either needs to cancel it or to provide the
response body.
The stream only supports sequential read, IO.read will fail if the position
is specified.
This method is mutually exclusive with getResponseBody.
Calling other methods that affect the request or disabling fetch
domain before body is received results in an undefined behavior.
*/
func TakeResponseBodyAsStream(c protocol.Caller, args TakeResponseBodyAsStreamArgs) (*TakeResponseBodyAsStreamVal, error) {
	var val = &TakeResponseBodyAsStreamVal{}
	return val, c.Call("Fetch.takeResponseBodyAsStream", args, val)
}
