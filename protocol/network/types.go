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
	Unique request identifier.
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
type Headers interface{}

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
	RequestTime              float64 `json:"requestTime"`
	ProxyStart               float64 `json:"proxyStart"`
	ProxyEnd                 float64 `json:"proxyEnd"`
	DnsStart                 float64 `json:"dnsStart"`
	DnsEnd                   float64 `json:"dnsEnd"`
	ConnectStart             float64 `json:"connectStart"`
	ConnectEnd               float64 `json:"connectEnd"`
	SslStart                 float64 `json:"sslStart"`
	SslEnd                   float64 `json:"sslEnd"`
	WorkerStart              float64 `json:"workerStart"`
	WorkerReady              float64 `json:"workerReady"`
	WorkerFetchStart         float64 `json:"workerFetchStart"`
	WorkerRespondWithSettled float64 `json:"workerRespondWithSettled"`
	SendStart                float64 `json:"sendStart"`
	SendEnd                  float64 `json:"sendEnd"`
	PushStart                float64 `json:"pushStart"`
	PushEnd                  float64 `json:"pushEnd"`
	ReceiveHeadersEnd        float64 `json:"receiveHeadersEnd"`
}

/*
	Loading priority of a resource request.
*/
type ResourcePriority string

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
	PostData         string                    `json:"postData,omitempty"`
	HasPostData      bool                      `json:"hasPostData,omitempty"`
	PostDataEntries  []*PostDataEntry          `json:"postDataEntries,omitempty"`
	MixedContentType security.MixedContentType `json:"mixedContentType,omitempty"`
	InitialPriority  ResourcePriority          `json:"initialPriority"`
	ReferrerPolicy   string                    `json:"referrerPolicy"`
	IsLinkPreload    bool                      `json:"isLinkPreload,omitempty"`
	TrustTokenParams *TrustTokenParams         `json:"trustTokenParams,omitempty"`
}

/*
	Details of a signed certificate timestamp (SCT).
*/
type SignedCertificateTimestamp struct {
	Status             string                `json:"status"`
	Origin             string                `json:"origin"`
	LogDescription     string                `json:"logDescription"`
	LogId              string                `json:"logId"`
	Timestamp          common.TimeSinceEpoch `json:"timestamp"`
	HashAlgorithm      string                `json:"hashAlgorithm"`
	SignatureAlgorithm string                `json:"signatureAlgorithm"`
	SignatureData      string                `json:"signatureData"`
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
	Type          TrustTokenOperationType `json:"type"`
	RefreshPolicy string                  `json:"refreshPolicy"`
	Issuers       []string                `json:"issuers,omitempty"`
}

/*

 */
type TrustTokenOperationType string

/*
	HTTP response data.
*/
type Response struct {
	Url                         string                      `json:"url"`
	Status                      int                         `json:"status"`
	StatusText                  string                      `json:"statusText"`
	Headers                     *Headers                    `json:"headers"`
	HeadersText                 string                      `json:"headersText,omitempty"`
	MimeType                    string                      `json:"mimeType"`
	RequestHeaders              *Headers                    `json:"requestHeaders,omitempty"`
	RequestHeadersText          string                      `json:"requestHeadersText,omitempty"`
	ConnectionReused            bool                        `json:"connectionReused"`
	ConnectionId                float64                     `json:"connectionId"`
	RemoteIPAddress             string                      `json:"remoteIPAddress,omitempty"`
	RemotePort                  int                         `json:"remotePort,omitempty"`
	FromDiskCache               bool                        `json:"fromDiskCache,omitempty"`
	FromServiceWorker           bool                        `json:"fromServiceWorker,omitempty"`
	FromPrefetchCache           bool                        `json:"fromPrefetchCache,omitempty"`
	EncodedDataLength           float64                     `json:"encodedDataLength"`
	Timing                      *ResourceTiming             `json:"timing,omitempty"`
	ServiceWorkerResponseSource ServiceWorkerResponseSource `json:"serviceWorkerResponseSource,omitempty"`
	ResponseTime                common.TimeSinceEpoch       `json:"responseTime,omitempty"`
	CacheStorageCacheName       string                      `json:"cacheStorageCacheName,omitempty"`
	Protocol                    string                      `json:"protocol,omitempty"`
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
	Cookie object
*/
type Cookie struct {
	Name         string             `json:"name"`
	Value        string             `json:"value"`
	Domain       string             `json:"domain"`
	Path         string             `json:"path"`
	Expires      float64            `json:"expires"`
	Size         int                `json:"size"`
	HttpOnly     bool               `json:"httpOnly"`
	Secure       bool               `json:"secure"`
	Session      bool               `json:"session"`
	SameSite     CookieSameSite     `json:"sameSite,omitempty"`
	Priority     CookiePriority     `json:"priority"`
	SameParty    bool               `json:"sameParty"`
	SourceScheme CookieSourceScheme `json:"sourceScheme"`
	SourcePort   int                `json:"sourcePort"`
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
	A cookie which was not stored from a response with the corresponding reason.
*/
type BlockedSetCookieWithReason struct {
	BlockedReasons []SetCookieBlockedReason `json:"blockedReasons"`
	CookieLine     string                   `json:"cookieLine"`
	Cookie         *Cookie                  `json:"cookie,omitempty"`
}

/*
	A cookie with was not sent with a request with the corresponding reason.
*/
type BlockedCookieWithReason struct {
	BlockedReasons []CookieBlockedReason `json:"blockedReasons"`
	Cookie         *Cookie               `json:"cookie"`
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
	SameParty    bool                  `json:"sameParty,omitempty"`
	SourceScheme CookieSourceScheme    `json:"sourceScheme,omitempty"`
	SourcePort   int                   `json:"sourcePort,omitempty"`
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
type PrivateNetworkRequestPolicy string

/*

 */
type IPAddressSpace string

/*

 */
type ClientSecurityState struct {
	InitiatorIsSecureContext    bool                        `json:"initiatorIsSecureContext"`
	InitiatorIPAddressSpace     IPAddressSpace              `json:"initiatorIPAddressSpace"`
	PrivateNetworkRequestPolicy PrivateNetworkRequestPolicy `json:"privateNetworkRequestPolicy"`
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
type SecurityIsolationStatus struct {
	Coop *CrossOriginOpenerPolicyStatus   `json:"coop,omitempty"`
	Coep *CrossOriginEmbedderPolicyStatus `json:"coep,omitempty"`
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
	Name   string `json:"name"`
	Url    string `json:"url,omitempty"`
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

type EmulateNetworkConditionsArgs struct {
	Offline            bool           `json:"offline"`
	Latency            float64        `json:"latency"`
	DownloadThroughput float64        `json:"downloadThroughput"`
	UploadThroughput   float64        `json:"uploadThroughput"`
	ConnectionType     ConnectionType `json:"connectionType,omitempty"`
}

type EnableArgs struct {
	MaxTotalBufferSize    int `json:"maxTotalBufferSize,omitempty"`
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
	MaxPostDataSize       int `json:"maxPostDataSize,omitempty"`
}

type GetAllCookiesVal struct {
	Cookies []*Cookie `json:"cookies"`
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
	PostData string `json:"postData"`
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
	Urls []string `json:"urls"`
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
	SameParty    bool                  `json:"sameParty,omitempty"`
	SourceScheme CookieSourceScheme    `json:"sourceScheme,omitempty"`
	SourcePort   int                   `json:"sourcePort,omitempty"`
}

type SetCookiesArgs struct {
	Cookies []*CookieParam `json:"cookies"`
}

type SetDataSizeLimitsForTestArgs struct {
	MaxTotalSize    int `json:"maxTotalSize"`
	MaxResourceSize int `json:"maxResourceSize"`
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

type GetSecurityIsolationStatusArgs struct {
	FrameId common.FrameId `json:"frameId,omitempty"`
}

type GetSecurityIsolationStatusVal struct {
	Status *SecurityIsolationStatus `json:"status"`
}

type LoadNetworkResourceArgs struct {
	FrameId common.FrameId              `json:"frameId"`
	Url     string                      `json:"url"`
	Options *LoadNetworkResourceOptions `json:"options"`
}

type LoadNetworkResourceVal struct {
	Resource *LoadNetworkResourcePageResult `json:"resource"`
}
