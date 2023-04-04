package network

import (
	"github.com/ecwid/control/protocol/common"
)

/*
Fired when data chunk was received over the network.
*/
type DataReceived struct {
	RequestId         RequestId     `json:"requestId"`
	Timestamp         MonotonicTime `json:"timestamp"`
	DataLength        int           `json:"dataLength"`
	EncodedDataLength int           `json:"encodedDataLength"`
}

/*
Fired when EventSource message is received.
*/
type EventSourceMessageReceived struct {
	RequestId RequestId     `json:"requestId"`
	Timestamp MonotonicTime `json:"timestamp"`
	EventName string        `json:"eventName"`
	EventId   string        `json:"eventId"`
	Data      string        `json:"data"`
}

/*
Fired when HTTP request has failed to load.
*/
type LoadingFailed struct {
	RequestId       RequestId        `json:"requestId"`
	Timestamp       MonotonicTime    `json:"timestamp"`
	Type            ResourceType     `json:"type"`
	ErrorText       string           `json:"errorText"`
	Canceled        bool             `json:"canceled,omitempty"`
	BlockedReason   BlockedReason    `json:"blockedReason,omitempty"`
	CorsErrorStatus *CorsErrorStatus `json:"corsErrorStatus,omitempty"`
}

/*
Fired when HTTP request has finished loading.
*/
type LoadingFinished struct {
	RequestId                RequestId     `json:"requestId"`
	Timestamp                MonotonicTime `json:"timestamp"`
	EncodedDataLength        float64       `json:"encodedDataLength"`
	ShouldReportCorbBlocking bool          `json:"shouldReportCorbBlocking,omitempty"`
}

/*
Fired if request ended up loading from cache.
*/
type RequestServedFromCache struct {
	RequestId RequestId `json:"requestId"`
}

/*
Fired when page is about to send HTTP request.
*/
type RequestWillBeSent struct {
	RequestId            RequestId             `json:"requestId"`
	LoaderId             LoaderId              `json:"loaderId"`
	DocumentURL          string                `json:"documentURL"`
	Request              *Request              `json:"request"`
	Timestamp            MonotonicTime         `json:"timestamp"`
	WallTime             common.TimeSinceEpoch `json:"wallTime"`
	Initiator            *Initiator            `json:"initiator"`
	RedirectHasExtraInfo bool                  `json:"redirectHasExtraInfo"`
	RedirectResponse     *Response             `json:"redirectResponse,omitempty"`
	Type                 ResourceType          `json:"type,omitempty"`
	FrameId              common.FrameId        `json:"frameId,omitempty"`
	HasUserGesture       bool                  `json:"hasUserGesture,omitempty"`
}

/*
Fired when resource loading priority is changed
*/
type ResourceChangedPriority struct {
	RequestId   RequestId        `json:"requestId"`
	NewPriority ResourcePriority `json:"newPriority"`
	Timestamp   MonotonicTime    `json:"timestamp"`
}

/*
Fired when a signed exchange was received over the network
*/
type SignedExchangeReceived struct {
	RequestId RequestId           `json:"requestId"`
	Info      *SignedExchangeInfo `json:"info"`
}

/*
Fired when HTTP response is available.
*/
type ResponseReceived struct {
	RequestId    RequestId      `json:"requestId"`
	LoaderId     LoaderId       `json:"loaderId"`
	Timestamp    MonotonicTime  `json:"timestamp"`
	Type         ResourceType   `json:"type"`
	Response     *Response      `json:"response"`
	HasExtraInfo bool           `json:"hasExtraInfo"`
	FrameId      common.FrameId `json:"frameId,omitempty"`
}

/*
Fired when WebSocket is closed.
*/
type WebSocketClosed struct {
	RequestId RequestId     `json:"requestId"`
	Timestamp MonotonicTime `json:"timestamp"`
}

/*
Fired upon WebSocket creation.
*/
type WebSocketCreated struct {
	RequestId RequestId  `json:"requestId"`
	Url       string     `json:"url"`
	Initiator *Initiator `json:"initiator,omitempty"`
}

/*
Fired when WebSocket message error occurs.
*/
type WebSocketFrameError struct {
	RequestId    RequestId     `json:"requestId"`
	Timestamp    MonotonicTime `json:"timestamp"`
	ErrorMessage string        `json:"errorMessage"`
}

/*
Fired when WebSocket message is received.
*/
type WebSocketFrameReceived struct {
	RequestId RequestId       `json:"requestId"`
	Timestamp MonotonicTime   `json:"timestamp"`
	Response  *WebSocketFrame `json:"response"`
}

/*
Fired when WebSocket message is sent.
*/
type WebSocketFrameSent struct {
	RequestId RequestId       `json:"requestId"`
	Timestamp MonotonicTime   `json:"timestamp"`
	Response  *WebSocketFrame `json:"response"`
}

/*
Fired when WebSocket handshake response becomes available.
*/
type WebSocketHandshakeResponseReceived struct {
	RequestId RequestId          `json:"requestId"`
	Timestamp MonotonicTime      `json:"timestamp"`
	Response  *WebSocketResponse `json:"response"`
}

/*
Fired when WebSocket is about to initiate handshake.
*/
type WebSocketWillSendHandshakeRequest struct {
	RequestId RequestId             `json:"requestId"`
	Timestamp MonotonicTime         `json:"timestamp"`
	WallTime  common.TimeSinceEpoch `json:"wallTime"`
	Request   *WebSocketRequest     `json:"request"`
}

/*
Fired upon WebTransport creation.
*/
type WebTransportCreated struct {
	TransportId RequestId     `json:"transportId"`
	Url         string        `json:"url"`
	Timestamp   MonotonicTime `json:"timestamp"`
	Initiator   *Initiator    `json:"initiator,omitempty"`
}

/*
Fired when WebTransport handshake is finished.
*/
type WebTransportConnectionEstablished struct {
	TransportId RequestId     `json:"transportId"`
	Timestamp   MonotonicTime `json:"timestamp"`
}

/*
Fired when WebTransport is disposed.
*/
type WebTransportClosed struct {
	TransportId RequestId     `json:"transportId"`
	Timestamp   MonotonicTime `json:"timestamp"`
}

/*
	Fired when additional information about a requestWillBeSent event is available from the

network stack. Not every requestWillBeSent event will have an additional
requestWillBeSentExtraInfo fired for it, and there is no guarantee whether requestWillBeSent
or requestWillBeSentExtraInfo will be fired first for the same request.
*/
type RequestWillBeSentExtraInfo struct {
	RequestId                     RequestId                  `json:"requestId"`
	AssociatedCookies             []*BlockedCookieWithReason `json:"associatedCookies"`
	Headers                       *Headers                   `json:"headers"`
	ConnectTiming                 *ConnectTiming             `json:"connectTiming"`
	ClientSecurityState           *ClientSecurityState       `json:"clientSecurityState,omitempty"`
	SiteHasCookieInOtherPartition bool                       `json:"siteHasCookieInOtherPartition,omitempty"`
}

/*
	Fired when additional information about a responseReceived event is available from the network

stack. Not every responseReceived event will have an additional responseReceivedExtraInfo for
it, and responseReceivedExtraInfo may be fired before or after responseReceived.
*/
type ResponseReceivedExtraInfo struct {
	RequestId                RequestId                     `json:"requestId"`
	BlockedCookies           []*BlockedSetCookieWithReason `json:"blockedCookies"`
	Headers                  *Headers                      `json:"headers"`
	ResourceIPAddressSpace   IPAddressSpace                `json:"resourceIPAddressSpace"`
	StatusCode               int                           `json:"statusCode"`
	HeadersText              string                        `json:"headersText,omitempty"`
	CookiePartitionKey       string                        `json:"cookiePartitionKey,omitempty"`
	CookiePartitionKeyOpaque bool                          `json:"cookiePartitionKeyOpaque,omitempty"`
}

/*
	Fired exactly once for each Trust Token operation. Depending on

the type of the operation and whether the operation succeeded or
failed, the event is fired before the corresponding request was sent
or after the response was received.
*/
type TrustTokenOperationDone struct {
	Status           string                  `json:"status"`
	Type             TrustTokenOperationType `json:"type"`
	RequestId        RequestId               `json:"requestId"`
	TopLevelOrigin   string                  `json:"topLevelOrigin,omitempty"`
	IssuerOrigin     string                  `json:"issuerOrigin,omitempty"`
	IssuedTokenCount int                     `json:"issuedTokenCount,omitempty"`
}

/*
	Fired once when parsing the .wbn file has succeeded.

The event contains the information about the web bundle contents.
*/
type SubresourceWebBundleMetadataReceived struct {
	RequestId RequestId `json:"requestId"`
	Urls      []string  `json:"urls"`
}

/*
Fired once when parsing the .wbn file has failed.
*/
type SubresourceWebBundleMetadataError struct {
	RequestId    RequestId `json:"requestId"`
	ErrorMessage string    `json:"errorMessage"`
}

/*
	Fired when handling requests for resources within a .wbn file.

Note: this will only be fired for resources that are requested by the webpage.
*/
type SubresourceWebBundleInnerResponseParsed struct {
	InnerRequestId  RequestId `json:"innerRequestId"`
	InnerRequestURL string    `json:"innerRequestURL"`
	BundleRequestId RequestId `json:"bundleRequestId,omitempty"`
}

/*
Fired when request for resources within a .wbn file failed.
*/
type SubresourceWebBundleInnerResponseError struct {
	InnerRequestId  RequestId `json:"innerRequestId"`
	InnerRequestURL string    `json:"innerRequestURL"`
	ErrorMessage    string    `json:"errorMessage"`
	BundleRequestId RequestId `json:"bundleRequestId,omitempty"`
}

/*
	Is sent whenever a new report is added.

And after 'enableReportingApi' for all existing reports.
*/
type ReportingApiReportAdded struct {
	Report *ReportingApiReport `json:"report"`
}

/*
 */
type ReportingApiReportUpdated struct {
	Report *ReportingApiReport `json:"report"`
}

/*
 */
type ReportingApiEndpointsChangedForOrigin struct {
	Origin    string                  `json:"origin"`
	Endpoints []*ReportingApiEndpoint `json:"endpoints"`
}
