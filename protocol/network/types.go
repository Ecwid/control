package network

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/debugger"
	"github.com/ecwid/control/protocol/io"
	"github.com/ecwid/control/protocol/runtime"
	"github.com/ecwid/control/protocol/security"
)

/*
Resource type as it was perceived by the rendering engine.
*/
type ResourceType string

/*
Unique loader identifier.
*/
type LoaderId string

/*
	Unique network request identifier.

Note that this does not identify individual HTTP requests that are part of
a network request.
*/
type RequestId string

/*
Unique intercepted request identifier.
*/
type InterceptionId string

/*
Network level fetch failure reason.
*/
type ErrorReason string

/*
UTC time in seconds, counted from January 1, 1970.
*/
type TimeSinceEpoch float64

/*
Monotonically increasing time in seconds since an arbitrary point in the past.
*/
type MonotonicTime float64

/*
Request / response headers as keys / values of JSON object.
*/
type Headers any

/*
The underlying connection technology that the browser is supposedly using.
*/
type ConnectionType string

/*
	Represents the cookie's 'SameSite' status:

https://tools.ietf.org/html/draft-west-first-party-cookies
*/
type CookieSameSite string

/*
	Represents the cookie's 'Priority' status:

https://tools.ietf.org/html/draft-west-cookie-priority-00
*/
type CookiePriority string

/*
	Represents the source scheme of the origin that originally set the cookie.

A value of "Unset" allows protocol clients to emulate legacy cookie scope for the scheme.
This is a temporary ability and it will be removed in the future.
*/
type CookieSourceScheme string

/*
Timing information for the request.
*/
type ResourceTiming struct {
	RequestTime                 float64 `json:"requestTime"`
	ProxyStart                  float64 `json:"proxyStart"`
	ProxyEnd                    float64 `json:"proxyEnd"`
	DnsStart                    float64 `json:"dnsStart"`
	DnsEnd                      float64 `json:"dnsEnd"`
	ConnectStart                float64 `json:"connectStart"`
	ConnectEnd                  float64 `json:"connectEnd"`
	SslStart                    float64 `json:"sslStart"`
	SslEnd                      float64 `json:"sslEnd"`
	WorkerStart                 float64 `json:"workerStart"`
	WorkerReady                 float64 `json:"workerReady"`
	WorkerFetchStart            float64 `json:"workerFetchStart"`
	WorkerRespondWithSettled    float64 `json:"workerRespondWithSettled"`
	WorkerRouterEvaluationStart float64 `json:"workerRouterEvaluationStart,omitempty"`
	WorkerCacheLookupStart      float64 `json:"workerCacheLookupStart,omitempty"`
	SendStart                   float64 `json:"sendStart"`
	SendEnd                     float64 `json:"sendEnd"`
	PushStart                   float64 `json:"pushStart"`
	PushEnd                     float64 `json:"pushEnd"`
	ReceiveHeadersStart         float64 `json:"receiveHeadersStart"`
	ReceiveHeadersEnd           float64 `json:"receiveHeadersEnd"`
}

/*
Loading priority of a resource request.
*/
type ResourcePriority string

/*
The render-blocking behavior of a resource request.
*/
type RenderBlockingBehavior string

/*
Post data entry for HTTP request
*/
type PostDataEntry struct {
	Bytes []byte `json:"bytes,omitempty"`
}

/*
HTTP request data.
*/
type Request struct {
	Url              string                    `json:"url"`
	UrlFragment      string                    `json:"urlFragment,omitempty"`
	Method           string                    `json:"method"`
	Headers          *Headers                  `json:"headers"`
	HasPostData      bool                      `json:"hasPostData,omitempty"`
	PostDataEntries  []*PostDataEntry          `json:"postDataEntries,omitempty"`
	MixedContentType security.MixedContentType `json:"mixedContentType,omitempty"`
	InitialPriority  ResourcePriority          `json:"initialPriority"`
	ReferrerPolicy   string                    `json:"referrerPolicy"`
	IsLinkPreload    bool                      `json:"isLinkPreload,omitempty"`
	TrustTokenParams *TrustTokenParams         `json:"trustTokenParams,omitempty"`
	IsSameSite       bool                      `json:"isSameSite,omitempty"`
	IsAdRelated      bool                      `json:"isAdRelated,omitempty"`
}

/*
Details of a signed certificate timestamp (SCT).
*/
type SignedCertificateTimestamp struct {
	Status             string  `json:"status"`
	Origin             string  `json:"origin"`
	LogDescription     string  `json:"logDescription"`
	LogId              string  `json:"logId"`
	Timestamp          float64 `json:"timestamp"`
	HashAlgorithm      string  `json:"hashAlgorithm"`
	SignatureAlgorithm string  `json:"signatureAlgorithm"`
	SignatureData      string  `json:"signatureData"`
}

/*
Security details about a request.
*/
type SecurityDetails struct {
	Protocol                          string                            `json:"protocol"`
	KeyExchange                       string                            `json:"keyExchange"`
	KeyExchangeGroup                  string                            `json:"keyExchangeGroup,omitempty"`
	Cipher                            string                            `json:"cipher"`
	Mac                               string                            `json:"mac,omitempty"`
	CertificateId                     security.CertificateId            `json:"certificateId"`
	SubjectName                       string                            `json:"subjectName"`
	SanList                           []string                          `json:"sanList"`
	Issuer                            string                            `json:"issuer"`
	ValidFrom                         common.TimeSinceEpoch             `json:"validFrom"`
	ValidTo                           common.TimeSinceEpoch             `json:"validTo"`
	SignedCertificateTimestampList    []*SignedCertificateTimestamp     `json:"signedCertificateTimestampList"`
	CertificateTransparencyCompliance CertificateTransparencyCompliance `json:"certificateTransparencyCompliance"`
	ServerSignatureAlgorithm          int                               `json:"serverSignatureAlgorithm,omitempty"`
	EncryptedClientHello              bool                              `json:"encryptedClientHello"`
}

/*
Whether the request complied with Certificate Transparency policy.
*/
type CertificateTransparencyCompliance string

/*
The reason why request was blocked.
*/
type BlockedReason string

/*
The reason why request was blocked.
*/
type CorsError string

/*
 */
type CorsErrorStatus struct {
	CorsError       CorsError `json:"corsError"`
	FailedParameter string    `json:"failedParameter"`
}

/*
Source of serviceworker response.
*/
type ServiceWorkerResponseSource string

/*
	Determines what type of Trust Token operation is executed and

depending on the type, some additional parameters. The values
are specified in third_party/blink/renderer/core/fetch/trust_token.idl.
*/
type TrustTokenParams struct {
	Operation     TrustTokenOperationType `json:"operation"`
	RefreshPolicy string                  `json:"refreshPolicy"`
	Issuers       []string                `json:"issuers,omitempty"`
}

/*
 */
type TrustTokenOperationType string

/*
The reason why Chrome uses a specific transport protocol for HTTP semantics.
*/
type AlternateProtocolUsage string

/*
Source of service worker router.
*/
type ServiceWorkerRouterSource string

/*
 */
type ServiceWorkerRouterInfo struct {
	RuleIdMatched     int                       `json:"ruleIdMatched,omitempty"`
	MatchedSourceType ServiceWorkerRouterSource `json:"matchedSourceType,omitempty"`
	ActualSourceType  ServiceWorkerRouterSource `json:"actualSourceType,omitempty"`
}

/*
HTTP response data.
*/
type Response struct {
	Url                         string                      `json:"url"`
	Status                      int                         `json:"status"`
	StatusText                  string                      `json:"statusText"`
	Headers                     *Headers                    `json:"headers"`
	MimeType                    string                      `json:"mimeType"`
	Charset                     string                      `json:"charset"`
	RequestHeaders              *Headers                    `json:"requestHeaders,omitempty"`
	ConnectionReused            bool                        `json:"connectionReused"`
	ConnectionId                float64                     `json:"connectionId"`
	RemoteIPAddress             string                      `json:"remoteIPAddress,omitempty"`
	RemotePort                  int                         `json:"remotePort,omitempty"`
	FromDiskCache               bool                        `json:"fromDiskCache,omitempty"`
	FromServiceWorker           bool                        `json:"fromServiceWorker,omitempty"`
	FromPrefetchCache           bool                        `json:"fromPrefetchCache,omitempty"`
	FromEarlyHints              bool                        `json:"fromEarlyHints,omitempty"`
	ServiceWorkerRouterInfo     *ServiceWorkerRouterInfo    `json:"serviceWorkerRouterInfo,omitempty"`
	EncodedDataLength           float64                     `json:"encodedDataLength"`
	Timing                      *ResourceTiming             `json:"timing,omitempty"`
	ServiceWorkerResponseSource ServiceWorkerResponseSource `json:"serviceWorkerResponseSource,omitempty"`
	ResponseTime                common.TimeSinceEpoch       `json:"responseTime,omitempty"`
	CacheStorageCacheName       string                      `json:"cacheStorageCacheName,omitempty"`
	Protocol                    string                      `json:"protocol,omitempty"`
	AlternateProtocolUsage      AlternateProtocolUsage      `json:"alternateProtocolUsage,omitempty"`
	SecurityState               security.SecurityState      `json:"securityState"`
	SecurityDetails             *SecurityDetails            `json:"securityDetails,omitempty"`
}

/*
WebSocket request data.
*/
type WebSocketRequest struct {
	Headers *Headers `json:"headers"`
}

/*
WebSocket response data.
*/
type WebSocketResponse struct {
	Status             int      `json:"status"`
	StatusText         string   `json:"statusText"`
	Headers            *Headers `json:"headers"`
	HeadersText        string   `json:"headersText,omitempty"`
	RequestHeaders     *Headers `json:"requestHeaders,omitempty"`
	RequestHeadersText string   `json:"requestHeadersText,omitempty"`
}

/*
WebSocket message data. This represents an entire WebSocket message, not just a fragmented frame as the name suggests.
*/
type WebSocketFrame struct {
	Opcode      float64 `json:"opcode"`
	Mask        bool    `json:"mask"`
	PayloadData string  `json:"payloadData"`
}

/*
Information about the cached resource.
*/
type CachedResource struct {
	Url      string       `json:"url"`
	Type     ResourceType `json:"type"`
	Response *Response    `json:"response,omitempty"`
	BodySize float64      `json:"bodySize"`
}

/*
Information about the request initiator.
*/
type Initiator struct {
	Type         string              `json:"type"`
	Stack        *runtime.StackTrace `json:"stack,omitempty"`
	Url          string              `json:"url,omitempty"`
	LineNumber   float64             `json:"lineNumber,omitempty"`
	ColumnNumber float64             `json:"columnNumber,omitempty"`
	RequestId    RequestId           `json:"requestId,omitempty"`
}

/*
	cookiePartitionKey object

The representation of the components of the key that are created by the cookiePartitionKey class contained in net/cookies/cookie_partition_key.h.
*/
type CookiePartitionKey struct {
	TopLevelSite         string `json:"topLevelSite"`
	HasCrossSiteAncestor bool   `json:"hasCrossSiteAncestor"`
}

/*
Cookie object
*/
type Cookie struct {
	Name               string              `json:"name"`
	Value              string              `json:"value"`
	Domain             string              `json:"domain"`
	Path               string              `json:"path"`
	Expires            float64             `json:"expires"`
	Size               int                 `json:"size"`
	HttpOnly           bool                `json:"httpOnly"`
	Secure             bool                `json:"secure"`
	Session            bool                `json:"session"`
	SameSite           CookieSameSite      `json:"sameSite,omitempty"`
	Priority           CookiePriority      `json:"priority"`
	SourceScheme       CookieSourceScheme  `json:"sourceScheme"`
	SourcePort         int                 `json:"sourcePort"`
	PartitionKey       *CookiePartitionKey `json:"partitionKey,omitempty"`
	PartitionKeyOpaque bool                `json:"partitionKeyOpaque,omitempty"`
}

/*
Types of reasons why a cookie may not be stored from a response.
*/
type SetCookieBlockedReason string

/*
Types of reasons why a cookie may not be sent with a request.
*/
type CookieBlockedReason string

/*
Types of reasons why a cookie should have been blocked by 3PCD but is exempted for the request.
*/
type CookieExemptionReason string

/*
A cookie which was not stored from a response with the corresponding reason.
*/
type BlockedSetCookieWithReason struct {
	BlockedReasons []SetCookieBlockedReason `json:"blockedReasons"`
	CookieLine     string                   `json:"cookieLine"`
	Cookie         *Cookie                  `json:"cookie,omitempty"`
}

/*
	A cookie should have been blocked by 3PCD but is exempted and stored from a response with the

corresponding reason. A cookie could only have at most one exemption reason.
*/
type ExemptedSetCookieWithReason struct {
	ExemptionReason CookieExemptionReason `json:"exemptionReason"`
	CookieLine      string                `json:"cookieLine"`
	Cookie          *Cookie               `json:"cookie"`
}

/*
	A cookie associated with the request which may or may not be sent with it.

Includes the cookies itself and reasons for blocking or exemption.
*/
type AssociatedCookie struct {
	Cookie          *Cookie               `json:"cookie"`
	BlockedReasons  []CookieBlockedReason `json:"blockedReasons"`
	ExemptionReason CookieExemptionReason `json:"exemptionReason,omitempty"`
}

/*
Cookie parameter object
*/
type CookieParam struct {
	Name         string                `json:"name"`
	Value        string                `json:"value"`
	Url          string                `json:"url,omitempty"`
	Domain       string                `json:"domain,omitempty"`
	Path         string                `json:"path,omitempty"`
	Secure       bool                  `json:"secure,omitempty"`
	HttpOnly     bool                  `json:"httpOnly,omitempty"`
	SameSite     CookieSameSite        `json:"sameSite,omitempty"`
	Expires      common.TimeSinceEpoch `json:"expires,omitempty"`
	Priority     CookiePriority        `json:"priority,omitempty"`
	SourceScheme CookieSourceScheme    `json:"sourceScheme,omitempty"`
	SourcePort   int                   `json:"sourcePort,omitempty"`
	PartitionKey *CookiePartitionKey   `json:"partitionKey,omitempty"`
}

/*
Authorization challenge for HTTP status code 401 or 407.
*/
type AuthChallenge struct {
	Source string `json:"source,omitempty"`
	Origin string `json:"origin"`
	Scheme string `json:"scheme"`
	Realm  string `json:"realm"`
}

/*
Response to an AuthChallenge.
*/
type AuthChallengeResponse struct {
	Response string `json:"response"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

/*
	Stages of the interception to begin intercepting. Request will intercept before the request is

sent. Response will intercept after the response is received.
*/
type InterceptionStage string

/*
Request pattern for interception.
*/
type RequestPattern struct {
	UrlPattern        string            `json:"urlPattern,omitempty"`
	ResourceType      ResourceType      `json:"resourceType,omitempty"`
	InterceptionStage InterceptionStage `json:"interceptionStage,omitempty"`
}

/*
	Information about a signed exchange signature.

https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
*/
type SignedExchangeSignature struct {
	Label        string   `json:"label"`
	Signature    string   `json:"signature"`
	Integrity    string   `json:"integrity"`
	CertUrl      string   `json:"certUrl,omitempty"`
	CertSha256   string   `json:"certSha256,omitempty"`
	ValidityUrl  string   `json:"validityUrl"`
	Date         int      `json:"date"`
	Expires      int      `json:"expires"`
	Certificates []string `json:"certificates,omitempty"`
}

/*
	Information about a signed exchange header.

https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
*/
type SignedExchangeHeader struct {
	RequestUrl      string                     `json:"requestUrl"`
	ResponseCode    int                        `json:"responseCode"`
	ResponseHeaders *Headers                   `json:"responseHeaders"`
	Signatures      []*SignedExchangeSignature `json:"signatures"`
	HeaderIntegrity string                     `json:"headerIntegrity"`
}

/*
Field type for a signed exchange related error.
*/
type SignedExchangeErrorField string

/*
Information about a signed exchange response.
*/
type SignedExchangeError struct {
	Message        string                   `json:"message"`
	SignatureIndex int                      `json:"signatureIndex,omitempty"`
	ErrorField     SignedExchangeErrorField `json:"errorField,omitempty"`
}

/*
Information about a signed exchange response.
*/
type SignedExchangeInfo struct {
	OuterResponse   *Response              `json:"outerResponse"`
	HasExtraInfo    bool                   `json:"hasExtraInfo"`
	Header          *SignedExchangeHeader  `json:"header,omitempty"`
	SecurityDetails *SecurityDetails       `json:"securityDetails,omitempty"`
	Errors          []*SignedExchangeError `json:"errors,omitempty"`
}

/*
List of content encodings supported by the backend.
*/
type ContentEncoding string

/*
 */
type NetworkConditions struct {
	UrlPattern         string         `json:"urlPattern"`
	Latency            float64        `json:"latency"`
	DownloadThroughput float64        `json:"downloadThroughput"`
	UploadThroughput   float64        `json:"uploadThroughput"`
	ConnectionType     ConnectionType `json:"connectionType,omitempty"`
	PacketLoss         float64        `json:"packetLoss,omitempty"`
	PacketQueueLength  int            `json:"packetQueueLength,omitempty"`
	PacketReordering   bool           `json:"packetReordering,omitempty"`
	Offline            bool           `json:"offline,omitempty"`
}

/*
 */
type BlockPattern struct {
	UrlPattern string `json:"urlPattern"`
	Block      bool   `json:"block"`
}

/*
 */
type DirectSocketDnsQueryType string

/*
 */
type DirectTCPSocketOptions struct {
	NoDelay           bool                     `json:"noDelay"`
	KeepAliveDelay    float64                  `json:"keepAliveDelay,omitempty"`
	SendBufferSize    float64                  `json:"sendBufferSize,omitempty"`
	ReceiveBufferSize float64                  `json:"receiveBufferSize,omitempty"`
	DnsQueryType      DirectSocketDnsQueryType `json:"dnsQueryType,omitempty"`
}

/*
 */
type DirectUDPSocketOptions struct {
	RemoteAddr                   string                   `json:"remoteAddr,omitempty"`
	RemotePort                   int                      `json:"remotePort,omitempty"`
	LocalAddr                    string                   `json:"localAddr,omitempty"`
	LocalPort                    int                      `json:"localPort,omitempty"`
	DnsQueryType                 DirectSocketDnsQueryType `json:"dnsQueryType,omitempty"`
	SendBufferSize               float64                  `json:"sendBufferSize,omitempty"`
	ReceiveBufferSize            float64                  `json:"receiveBufferSize,omitempty"`
	MulticastLoopback            bool                     `json:"multicastLoopback,omitempty"`
	MulticastTimeToLive          int                      `json:"multicastTimeToLive,omitempty"`
	MulticastAllowAddressSharing bool                     `json:"multicastAllowAddressSharing,omitempty"`
}

/*
 */
type DirectUDPMessage struct {
	Data       []byte `json:"data"`
	RemoteAddr string `json:"remoteAddr,omitempty"`
	RemotePort int    `json:"remotePort,omitempty"`
}

/*
 */
type LocalNetworkAccessRequestPolicy string

/*
 */
type IPAddressSpace string

/*
 */
type ConnectTiming struct {
	RequestTime float64 `json:"requestTime"`
}

/*
 */
type ClientSecurityState struct {
	InitiatorIsSecureContext        bool                            `json:"initiatorIsSecureContext"`
	InitiatorIPAddressSpace         IPAddressSpace                  `json:"initiatorIPAddressSpace"`
	LocalNetworkAccessRequestPolicy LocalNetworkAccessRequestPolicy `json:"localNetworkAccessRequestPolicy"`
}

/*
	Identifies the script on the stack that caused a resource or element to be

labeled as an ad. For resources, this indicates the context that triggered
the fetch. For elements, this indicates the context that caused the element
to be appended to the DOM.
*/
type AdScriptIdentifier struct {
	ScriptId   runtime.ScriptId         `json:"scriptId"`
	DebuggerId runtime.UniqueDebuggerId `json:"debuggerId"`
	Name       string                   `json:"name"`
}

/*
	Encapsulates the script ancestry and the root script filter list rule that

caused the resource or element to be labeled as an ad.
*/
type AdAncestry struct {
	AncestryChain            []*AdScriptIdentifier `json:"ancestryChain"`
	RootScriptFilterlistRule string                `json:"rootScriptFilterlistRule,omitempty"`
}

/*
	Represents the provenance of an ad resource or element. Only one of

`filterlistRule` or `adScriptAncestry` can be set. If `filterlistRule`
is provided, the resource URL directly matches a filter list rule. If
`adScriptAncestry` is provided, an ad script initiated the resource fetch or
appended the element to the DOM. If neither is provided, the entity is
known to be an ad, but provenance tracking information is unavailable.
*/
type AdProvenance struct {
	FilterlistRule   string      `json:"filterlistRule,omitempty"`
	AdScriptAncestry *AdAncestry `json:"adScriptAncestry,omitempty"`
}

/*
 */
type CrossOriginOpenerPolicyValue string

/*
 */
type CrossOriginOpenerPolicyStatus struct {
	Value                       CrossOriginOpenerPolicyValue `json:"value"`
	ReportOnlyValue             CrossOriginOpenerPolicyValue `json:"reportOnlyValue"`
	ReportingEndpoint           string                       `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string                       `json:"reportOnlyReportingEndpoint,omitempty"`
}

/*
 */
type CrossOriginEmbedderPolicyValue string

/*
 */
type CrossOriginEmbedderPolicyStatus struct {
	Value                       CrossOriginEmbedderPolicyValue `json:"value"`
	ReportOnlyValue             CrossOriginEmbedderPolicyValue `json:"reportOnlyValue"`
	ReportingEndpoint           string                         `json:"reportingEndpoint,omitempty"`
	ReportOnlyReportingEndpoint string                         `json:"reportOnlyReportingEndpoint,omitempty"`
}

/*
 */
type ContentSecurityPolicySource string

/*
 */
type ContentSecurityPolicyStatus struct {
	EffectiveDirectives string                      `json:"effectiveDirectives"`
	IsEnforced          bool                        `json:"isEnforced"`
	Source              ContentSecurityPolicySource `json:"source"`
}

/*
 */
type SecurityIsolationStatus struct {
	Coop *CrossOriginOpenerPolicyStatus   `json:"coop,omitempty"`
	Coep *CrossOriginEmbedderPolicyStatus `json:"coep,omitempty"`
	Csp  []*ContentSecurityPolicyStatus   `json:"csp,omitempty"`
}

/*
The status of a Reporting API report.
*/
type ReportStatus string

/*
 */
type ReportId string

/*
An object representing a report generated by the Reporting API.
*/
type ReportingApiReport struct {
	Id                ReportId              `json:"id"`
	InitiatorUrl      string                `json:"initiatorUrl"`
	Destination       string                `json:"destination"`
	Type              string                `json:"type"`
	Timestamp         common.TimeSinceEpoch `json:"timestamp"`
	Depth             int                   `json:"depth"`
	CompletedAttempts int                   `json:"completedAttempts"`
	Body              any                   `json:"body"`
	Status            ReportStatus          `json:"status"`
}

/*
 */
type ReportingApiEndpoint struct {
	Url       string `json:"url"`
	GroupName string `json:"groupName"`
}

/*
Unique identifier for a device bound session.
*/
type DeviceBoundSessionKey struct {
	Site string `json:"site"`
	Id   string `json:"id"`
}

/*
How a device bound session was used during a request.
*/
type DeviceBoundSessionWithUsage struct {
	SessionKey *DeviceBoundSessionKey `json:"sessionKey"`
	Usage      string                 `json:"usage"`
}

/*
A device bound session's cookie craving.
*/
type DeviceBoundSessionCookieCraving struct {
	Name     string         `json:"name"`
	Domain   string         `json:"domain"`
	Path     string         `json:"path"`
	Secure   bool           `json:"secure"`
	HttpOnly bool           `json:"httpOnly"`
	SameSite CookieSameSite `json:"sameSite,omitempty"`
}

/*
A device bound session's inclusion URL rule.
*/
type DeviceBoundSessionUrlRule struct {
	RuleType    string `json:"ruleType"`
	HostPattern string `json:"hostPattern"`
	PathPrefix  string `json:"pathPrefix"`
}

/*
A device bound session's inclusion rules.
*/
type DeviceBoundSessionInclusionRules struct {
	Origin      string                       `json:"origin"`
	IncludeSite bool                         `json:"includeSite"`
	UrlRules    []*DeviceBoundSessionUrlRule `json:"urlRules"`
}

/*
A device bound session.
*/
type DeviceBoundSession struct {
	Key                      *DeviceBoundSessionKey             `json:"key"`
	RefreshUrl               string                             `json:"refreshUrl"`
	InclusionRules           *DeviceBoundSessionInclusionRules  `json:"inclusionRules"`
	CookieCravings           []*DeviceBoundSessionCookieCraving `json:"cookieCravings"`
	ExpiryDate               common.TimeSinceEpoch              `json:"expiryDate"`
	CachedChallenge          string                             `json:"cachedChallenge,omitempty"`
	AllowedRefreshInitiators []string                           `json:"allowedRefreshInitiators"`
}

/*
A unique identifier for a device bound session event.
*/
type DeviceBoundSessionEventId string

/*
A fetch result for a device bound session creation or refresh.
*/
type DeviceBoundSessionFetchResult string

/*
Details about a failed device bound session network request.
*/
type DeviceBoundSessionFailedRequest struct {
	RequestUrl        string `json:"requestUrl"`
	NetError          string `json:"netError,omitempty"`
	ResponseError     int    `json:"responseError,omitempty"`
	ResponseErrorBody string `json:"responseErrorBody,omitempty"`
}

/*
Session event details specific to creation.
*/
type CreationEventDetails struct {
	FetchResult   DeviceBoundSessionFetchResult    `json:"fetchResult"`
	NewSession    *DeviceBoundSession              `json:"newSession,omitempty"`
	FailedRequest *DeviceBoundSessionFailedRequest `json:"failedRequest,omitempty"`
}

/*
Session event details specific to refresh.
*/
type RefreshEventDetails struct {
	RefreshResult            string                           `json:"refreshResult"`
	FetchResult              DeviceBoundSessionFetchResult    `json:"fetchResult,omitempty"`
	NewSession               *DeviceBoundSession              `json:"newSession,omitempty"`
	WasFullyProactiveRefresh bool                             `json:"wasFullyProactiveRefresh"`
	FailedRequest            *DeviceBoundSessionFailedRequest `json:"failedRequest,omitempty"`
}

/*
Session event details specific to termination.
*/
type TerminationEventDetails struct {
	DeletionReason string `json:"deletionReason"`
}

/*
Session event details specific to challenges.
*/
type ChallengeEventDetails struct {
	ChallengeResult string `json:"challengeResult"`
	Challenge       string `json:"challenge"`
}

/*
An object providing the result of a network resource load.
*/
type LoadNetworkResourcePageResult struct {
	Success        bool            `json:"success"`
	NetError       float64         `json:"netError,omitempty"`
	NetErrorName   string          `json:"netErrorName,omitempty"`
	HttpStatusCode float64         `json:"httpStatusCode,omitempty"`
	Stream         io.StreamHandle `json:"stream,omitempty"`
	Headers        *Headers        `json:"headers,omitempty"`
}

/*
	An options object that may be extended later to better support CORS,

CORB and streaming.
*/
type LoadNetworkResourceOptions struct {
	DisableCache       bool `json:"disableCache"`
	IncludeCredentials bool `json:"includeCredentials"`
}

type SetAcceptedEncodingsArgs struct {
	Encodings []ContentEncoding `json:"encodings"`
}

type DeleteCookiesArgs struct {
	Name         string              `json:"name"`
	Url          string              `json:"url,omitempty"`
	Domain       string              `json:"domain,omitempty"`
	Path         string              `json:"path,omitempty"`
	PartitionKey *CookiePartitionKey `json:"partitionKey,omitempty"`
}

type EmulateNetworkConditionsByRuleArgs struct {
	EmulateOfflineServiceWorker bool                 `json:"emulateOfflineServiceWorker,omitempty"`
	MatchedNetworkConditions    []*NetworkConditions `json:"matchedNetworkConditions"`
}

type EmulateNetworkConditionsByRuleVal struct {
	RuleIds []string `json:"ruleIds"`
}

type OverrideNetworkStateArgs struct {
	Offline            bool           `json:"offline"`
	Latency            float64        `json:"latency"`
	DownloadThroughput float64        `json:"downloadThroughput"`
	UploadThroughput   float64        `json:"uploadThroughput"`
	ConnectionType     ConnectionType `json:"connectionType,omitempty"`
}

type EnableArgs struct {
	MaxTotalBufferSize        int  `json:"maxTotalBufferSize,omitempty"`
	MaxResourceBufferSize     int  `json:"maxResourceBufferSize,omitempty"`
	MaxPostDataSize           int  `json:"maxPostDataSize,omitempty"`
	ReportDirectSocketTraffic bool `json:"reportDirectSocketTraffic,omitempty"`
	EnableDurableMessages     bool `json:"enableDurableMessages,omitempty"`
}

type ConfigureDurableMessagesArgs struct {
	MaxTotalBufferSize    int `json:"maxTotalBufferSize,omitempty"`
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
}

type GetCertificateArgs struct {
	Origin string `json:"origin"`
}

type GetCertificateVal struct {
	TableNames []string `json:"tableNames"`
}

type GetCookiesArgs struct {
	Urls []string `json:"urls,omitempty"`
}

type GetCookiesVal struct {
	Cookies []*Cookie `json:"cookies"`
}

type GetResponseBodyArgs struct {
	RequestId RequestId `json:"requestId"`
}

type GetResponseBodyVal struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type GetRequestPostDataArgs struct {
	RequestId RequestId `json:"requestId"`
}

type GetRequestPostDataVal struct {
	PostData      string `json:"postData"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type GetResponseBodyForInterceptionArgs struct {
	InterceptionId InterceptionId `json:"interceptionId"`
}

type GetResponseBodyForInterceptionVal struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type TakeResponseBodyForInterceptionAsStreamArgs struct {
	InterceptionId InterceptionId `json:"interceptionId"`
}

type TakeResponseBodyForInterceptionAsStreamVal struct {
	Stream io.StreamHandle `json:"stream"`
}

type ReplayXHRArgs struct {
	RequestId RequestId `json:"requestId"`
}

type SearchInResponseBodyArgs struct {
	RequestId     RequestId `json:"requestId"`
	Query         string    `json:"query"`
	CaseSensitive bool      `json:"caseSensitive,omitempty"`
	IsRegex       bool      `json:"isRegex,omitempty"`
}

type SearchInResponseBodyVal struct {
	Result []*debugger.SearchMatch `json:"result"`
}

type SetBlockedURLsArgs struct {
	UrlPatterns []*BlockPattern `json:"urlPatterns,omitempty"`
}

type SetBypassServiceWorkerArgs struct {
	Bypass bool `json:"bypass"`
}

type SetCacheDisabledArgs struct {
	CacheDisabled bool `json:"cacheDisabled"`
}

type SetCookieArgs struct {
	Name         string                `json:"name"`
	Value        string                `json:"value"`
	Url          string                `json:"url,omitempty"`
	Domain       string                `json:"domain,omitempty"`
	Path         string                `json:"path,omitempty"`
	Secure       bool                  `json:"secure,omitempty"`
	HttpOnly     bool                  `json:"httpOnly,omitempty"`
	SameSite     CookieSameSite        `json:"sameSite,omitempty"`
	Expires      common.TimeSinceEpoch `json:"expires,omitempty"`
	Priority     CookiePriority        `json:"priority,omitempty"`
	SourceScheme CookieSourceScheme    `json:"sourceScheme,omitempty"`
	SourcePort   int                   `json:"sourcePort,omitempty"`
	PartitionKey *CookiePartitionKey   `json:"partitionKey,omitempty"`
}

type SetCookiesArgs struct {
	Cookies []*CookieParam `json:"cookies"`
}

type SetExtraHTTPHeadersArgs struct {
	Headers *Headers `json:"headers"`
}

type SetAttachDebugStackArgs struct {
	Enabled bool `json:"enabled"`
}

type SetUserAgentOverrideArgs struct {
	UserAgent         string                    `json:"userAgent"`
	AcceptLanguage    string                    `json:"acceptLanguage,omitempty"`
	Platform          string                    `json:"platform,omitempty"`
	UserAgentMetadata *common.UserAgentMetadata `json:"userAgentMetadata,omitempty"`
}

type StreamResourceContentArgs struct {
	RequestId RequestId `json:"requestId"`
}

type StreamResourceContentVal struct {
	BufferedData []byte `json:"bufferedData"`
}

type GetSecurityIsolationStatusArgs struct {
	FrameId common.FrameId `json:"frameId,omitempty"`
}

type GetSecurityIsolationStatusVal struct {
	Status *SecurityIsolationStatus `json:"status"`
}

type EnableReportingApiArgs struct {
	Enable bool `json:"enable"`
}

type EnableDeviceBoundSessionsArgs struct {
	Enable bool `json:"enable"`
}

type DeleteDeviceBoundSessionArgs struct {
	Key *DeviceBoundSessionKey `json:"key"`
}

type FetchSchemefulSiteArgs struct {
	Origin string `json:"origin"`
}

type FetchSchemefulSiteVal struct {
	SchemefulSite string `json:"schemefulSite"`
}

type LoadNetworkResourceArgs struct {
	FrameId common.FrameId              `json:"frameId,omitempty"`
	Url     string                      `json:"url"`
	Options *LoadNetworkResourceOptions `json:"options"`
}

type LoadNetworkResourceVal struct {
	Resource *LoadNetworkResourcePageResult `json:"resource"`
}

type SetCookieControlsArgs struct {
	EnableThirdPartyCookieRestriction bool `json:"enableThirdPartyCookieRestriction"`
}
