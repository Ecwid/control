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
	Data              []byte        `json:"data,omitempty"`
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
	RequestId         RequestId     `json:"requestId"`
	Timestamp         MonotonicTime `json:"timestamp"`
	EncodedDataLength float64       `json:"encodedDataLength"`
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
	RequestId              RequestId              `json:"requestId"`
	LoaderId               LoaderId               `json:"loaderId"`
	DocumentURL            string                 `json:"documentURL"`
	Request                *Request               `json:"request"`
	Timestamp              MonotonicTime          `json:"timestamp"`
	WallTime               common.TimeSinceEpoch  `json:"wallTime"`
	Initiator              *Initiator             `json:"initiator"`
	RedirectHasExtraInfo   bool                   `json:"redirectHasExtraInfo"`
	RedirectResponse       *Response              `json:"redirectResponse,omitempty"`
	Type                   ResourceType           `json:"type,omitempty"`
	FrameId                common.FrameId         `json:"frameId,omitempty"`
	HasUserGesture         bool                   `json:"hasUserGesture,omitempty"`
	RenderBlockingBehavior RenderBlockingBehavior `json:"renderBlockingBehavior,omitempty"`
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
Fired upon direct_socket.TCPSocket creation.
*/
type DirectTCPSocketCreated struct {
	Identifier RequestId               `json:"identifier"`
	RemoteAddr string                  `json:"remoteAddr"`
	RemotePort int                     `json:"remotePort"`
	Options    *DirectTCPSocketOptions `json:"options"`
	Timestamp  MonotonicTime           `json:"timestamp"`
	Initiator  *Initiator              `json:"initiator,omitempty"`
}

/*
Fired when direct_socket.TCPSocket connection is opened.
*/
type DirectTCPSocketOpened struct {
	Identifier RequestId     `json:"identifier"`
	RemoteAddr string        `json:"remoteAddr"`
	RemotePort int           `json:"remotePort"`
	Timestamp  MonotonicTime `json:"timestamp"`
	LocalAddr  string        `json:"localAddr,omitempty"`
	LocalPort  int           `json:"localPort,omitempty"`
}

/*
Fired when direct_socket.TCPSocket is aborted.
*/
type DirectTCPSocketAborted struct {
	Identifier   RequestId     `json:"identifier"`
	ErrorMessage string        `json:"errorMessage"`
	Timestamp    MonotonicTime `json:"timestamp"`
}

/*
Fired when direct_socket.TCPSocket is closed.
*/
type DirectTCPSocketClosed struct {
	Identifier RequestId     `json:"identifier"`
	Timestamp  MonotonicTime `json:"timestamp"`
}

/*
Fired when data is sent to tcp direct socket stream.
*/
type DirectTCPSocketChunkSent struct {
	Identifier RequestId     `json:"identifier"`
	Data       []byte        `json:"data"`
	Timestamp  MonotonicTime `json:"timestamp"`
}

/*
Fired when data is received from tcp direct socket stream.
*/
type DirectTCPSocketChunkReceived struct {
	Identifier RequestId     `json:"identifier"`
	Data       []byte        `json:"data"`
	Timestamp  MonotonicTime `json:"timestamp"`
}

/*
 */
type DirectUDPSocketJoinedMulticastGroup struct {
	Identifier RequestId `json:"identifier"`
	IPAddress  string    `json:"IPAddress"`
}

/*
 */
type DirectUDPSocketLeftMulticastGroup struct {
	Identifier RequestId `json:"identifier"`
	IPAddress  string    `json:"IPAddress"`
}

/*
Fired upon direct_socket.UDPSocket creation.
*/
type DirectUDPSocketCreated struct {
	Identifier RequestId               `json:"identifier"`
	Options    *DirectUDPSocketOptions `json:"options"`
	Timestamp  MonotonicTime           `json:"timestamp"`
	Initiator  *Initiator              `json:"initiator,omitempty"`
}

/*
Fired when direct_socket.UDPSocket connection is opened.
*/
type DirectUDPSocketOpened struct {
	Identifier RequestId     `json:"identifier"`
	LocalAddr  string        `json:"localAddr"`
	LocalPort  int           `json:"localPort"`
	Timestamp  MonotonicTime `json:"timestamp"`
	RemoteAddr string        `json:"remoteAddr,omitempty"`
	RemotePort int           `json:"remotePort,omitempty"`
}

/*
Fired when direct_socket.UDPSocket is aborted.
*/
type DirectUDPSocketAborted struct {
	Identifier   RequestId     `json:"identifier"`
	ErrorMessage string        `json:"errorMessage"`
	Timestamp    MonotonicTime `json:"timestamp"`
}

/*
Fired when direct_socket.UDPSocket is closed.
*/
type DirectUDPSocketClosed struct {
	Identifier RequestId     `json:"identifier"`
	Timestamp  MonotonicTime `json:"timestamp"`
}

/*
Fired when message is sent to udp direct socket stream.
*/
type DirectUDPSocketChunkSent struct {
	Identifier RequestId         `json:"identifier"`
	Message    *DirectUDPMessage `json:"message"`
	Timestamp  MonotonicTime     `json:"timestamp"`
}

/*
Fired when message is received from udp direct socket stream.
*/
type DirectUDPSocketChunkReceived struct {
	Identifier RequestId         `json:"identifier"`
	Message    *DirectUDPMessage `json:"message"`
	Timestamp  MonotonicTime     `json:"timestamp"`
}

/*
	Fired when additional information about a requestWillBeSent event is available from the

network stack. Not every requestWillBeSent event will have an additional
requestWillBeSentExtraInfo fired for it, and there is no guarantee whether requestWillBeSent
or requestWillBeSentExtraInfo will be fired first for the same request.
*/
type RequestWillBeSentExtraInfo struct {
	RequestId                     RequestId                      `json:"requestId"`
	AssociatedCookies             []*AssociatedCookie            `json:"associatedCookies"`
	Headers                       *Headers                       `json:"headers"`
	ConnectTiming                 *ConnectTiming                 `json:"connectTiming"`
	DeviceBoundSessionUsages      []*DeviceBoundSessionWithUsage `json:"deviceBoundSessionUsages,omitempty"`
	ClientSecurityState           *ClientSecurityState           `json:"clientSecurityState,omitempty"`
	SiteHasCookieInOtherPartition bool                           `json:"siteHasCookieInOtherPartition,omitempty"`
	AppliedNetworkConditionsId    string                         `json:"appliedNetworkConditionsId,omitempty"`
}

/*
	Fired when additional information about a responseReceived event is available from the network

stack. Not every responseReceived event will have an additional responseReceivedExtraInfo for
it, and responseReceivedExtraInfo may be fired before or after responseReceived.
*/
type ResponseReceivedExtraInfo struct {
	RequestId                RequestId                      `json:"requestId"`
	BlockedCookies           []*BlockedSetCookieWithReason  `json:"blockedCookies"`
	Headers                  *Headers                       `json:"headers"`
	ResourceIPAddressSpace   IPAddressSpace                 `json:"resourceIPAddressSpace"`
	StatusCode               int                            `json:"statusCode"`
	HeadersText              string                         `json:"headersText,omitempty"`
	CookiePartitionKey       *CookiePartitionKey            `json:"cookiePartitionKey,omitempty"`
	CookiePartitionKeyOpaque bool                           `json:"cookiePartitionKeyOpaque,omitempty"`
	ExemptedCookies          []*ExemptedSetCookieWithReason `json:"exemptedCookies,omitempty"`
}

/*
	Fired when 103 Early Hints headers is received in addition to the common response.

Not every responseReceived event will have an responseReceivedEarlyHints fired.
Only one responseReceivedEarlyHints may be fired for eached responseReceived event.
*/
type ResponseReceivedEarlyHints struct {
	RequestId RequestId `json:"requestId"`
	Headers   *Headers  `json:"headers"`
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
Fired once security policy has been updated.
*/
type PolicyUpdated any

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

/*
Triggered when the initial set of device bound sessions is added.
*/
type DeviceBoundSessionsAdded struct {
	Sessions []*DeviceBoundSession `json:"sessions"`
}

/*
Triggered when a device bound session event occurs.
*/
type DeviceBoundSessionEventOccurred struct {
	EventId                 DeviceBoundSessionEventId `json:"eventId"`
	Site                    string                    `json:"site"`
	Succeeded               bool                      `json:"succeeded"`
	SessionId               string                    `json:"sessionId,omitempty"`
	CreationEventDetails    *CreationEventDetails     `json:"creationEventDetails,omitempty"`
	RefreshEventDetails     *RefreshEventDetails      `json:"refreshEventDetails,omitempty"`
	TerminationEventDetails *TerminationEventDetails  `json:"terminationEventDetails,omitempty"`
	ChallengeEventDetails   *ChallengeEventDetails    `json:"challengeEventDetails,omitempty"`
}
