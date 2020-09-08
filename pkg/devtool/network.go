package devtool

// ResourceTiming https://chromedevtools.github.io/devtools-protocol/tot/Network#type-ResourceTiming
type ResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`
	ProxyStart        float64 `json:"proxyStart"`
	ProxyEnd          float64 `json:"proxyEnd"`
	DNSStart          float64 `json:"dnsStart"`
	DNSEnd            float64 `json:"dnsEnd"`
	ConnectStart      float64 `json:"connectStart"`
	ConnectEnd        float64 `json:"connectEnd"`
	SSLStart          float64 `json:"sslStart"`
	SSLEnd            float64 `json:"sslEnd"`
	WorkerStart       float64 `json:"workerStart"`
	WorkerReady       float64 `json:"workerReady"`
	SendStart         float64 `json:"sendStart"`
	SendEnd           float64 `json:"sendEnd"`
	PushStart         float64 `json:"pushStart"`
	PushEnd           float64 `json:"pushEnd"`
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
}

// Response https://chromedevtools.github.io/devtools-protocol/tot/Network#type-Response
type Response struct {
	URL                string                 `json:"url"`
	Status             int                    `json:"status"`
	StatusText         string                 `json:"statusText"`
	Headers            map[string]interface{} `json:"headers"`
	HeadersText        string                 `json:"headersText"`
	MimeType           string                 `json:"mimeType"`
	RequestHeaders     map[string]interface{} `json:"requestHeaders"`
	RequestHeadersText string                 `json:"requestHeadersText"`
	ConnectionReused   bool                   `json:"connectionReused"`
	ConnectionID       int64                  `json:"connectionId"`
	RemoteIPAddress    string                 `json:"remoteIPAddress"`
	RemotePort         int64                  `json:"remotePort"`
	FromDiskCache      bool                   `json:"fromDiskCache"`
	FromServiceWorker  bool                   `json:"fromServiceWorker"`
	FromPrefetchCache  bool                   `json:"fromPrefetchCache"`
	EncodedDataLength  int64                  `json:"encodedDataLength"`
	Timing             *ResourceTiming        `json:"timing"`
	Protocol           string                 `json:"devtool"`
	SecurityState      string                 `json:"securityState"`
}

// LoadingFailed https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFailed
type LoadingFailed struct {
	RequestID     string  `json:"requestId"`
	Timestamp     float64 `json:"timestamp"`
	Type          string  `json:"type"`
	ErrorText     string  `json:"errorText"`
	Canceled      bool    `json:"canceled"`
	BlockedReason string  `json:"blockedReason"`
}

// Cookie https://chromedevtools.github.io/devtools-protocol/tot/Network#type-CookieParam
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	URL      string `json:"url"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Expires  int64  `json:"expires"`
	Size     int64  `json:"size"`
	HTTPOnly bool   `json:"httpOnly"`
	Secure   bool   `json:"secure"`
}

// GetCookies https://chromedevtools.github.io/devtools-protocol/tot/Network#method-getCookies
type GetCookies struct {
	Cookies []*Cookie `json:"cookies"`
}

// Request https://chromedevtools.github.io/devtools-protocol/tot/Network#type-Request
type Request struct {
	URL              string                 `json:"url"`
	URLFragment      string                 `json:"urlFragment"`
	Method           string                 `json:"method"`
	Headers          map[string]interface{} `json:"headers"`
	PostData         string                 `json:"postData"`
	HasPostData      bool                   `json:"hasPostData"`
	MixedContentType string                 `json:"mixedContentType"`
	InitialPriority  string                 `json:"initialPriority"`
	ReferrerPolicy   string                 `json:"referrerPolicy"`
	IsLinkPreload    bool                   `json:"isLinkPreload"`
}

// RequestWillBeSent https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestWillBeSent
type RequestWillBeSent struct {
	RequestID        string    `json:"requestId"`
	LoaderID         string    `json:"loaderId"`
	DocumentURL      string    `json:"documentURL"`
	Request          *Request  `json:"request"`
	Timestamp        float64   `json:"timestamp"`
	WallTime         float64   `json:"wallTime"`
	RedirectResponse *Response `json:"redirectResponse"`
	Type             string    `json:"type"`
	FrameID          string    `json:"frameId"`
	HasUserGesture   bool      `json:"hasUserGesture"`
}

// ResponseReceived https://chromedevtools.github.io/devtools-protocol/tot/Network#event-responseReceived
type ResponseReceived struct {
	RequestID string    `json:"requestId"`
	LoaderID  string    `json:"loaderId"`
	Timestamp float64   `json:"timestamp"`
	Type      string    `json:"type"`
	Response  *Response `json:"response"`
	FrameID   string    `json:"frameId"`
}

// DataReceived https://chromedevtools.github.io/devtools-protocol/tot/Network#event-dataReceived
type DataReceived struct {
	RequestID         string  `json:"requestId"`
	Timestamp         float64 `json:"timestamp"`
	DataLength        int64   `json:"dataLength"`
	EncodedDataLength int64   `json:"encodedDataLength"`
}

// LoadingFinished https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFinished
type LoadingFinished struct {
	RequestID                string  `json:"requestId"`
	Timestamp                float64 `json:"timestamp"`
	EncodedDataLength        float64 `json:"encodedDataLength"`
	ShouldReportCorbBlocking bool    `json:"shouldReportCorbBlocking"`
}

// ServedFromCache https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestServedFromCache
type ServedFromCache struct {
	RequestID string `json:"requestId"`
}

// ErrorReason https://chromedevtools.github.io/devtools-protocol/tot/Network#type-ErrorReason
type ErrorReason string

// error reasons
const (
	Failed               ErrorReason = "Failed"
	Aborted              ErrorReason = "Aborted"
	TimedOut             ErrorReason = "TimedOut"
	AccessDenied         ErrorReason = "AccessDenied"
	ConnectionClosed     ErrorReason = "ConnectionClosed"
	ConnectionReset      ErrorReason = "ConnectionReset"
	ConnectionRefused    ErrorReason = "ConnectionRefused"
	ConnectionAborted    ErrorReason = "ConnectionAborted"
	ConnectionFailed     ErrorReason = "ConnectionFailed"
	NameNotResolved      ErrorReason = "NameNotResolved"
	InternetDisconnected ErrorReason = "InternetDisconnected"
	AddressUnreachable   ErrorReason = "AddressUnreachable"
	BlockedByClient      ErrorReason = "BlockedByClient"
	BlockedByResponse    ErrorReason = "BlockedByResponse"
)

// HeaderEntry https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-HeaderEntry
type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RequestPattern https://chromedevtools.github.io/devtools-protocol/tot/Fetch#type-RequestPattern
type RequestPattern struct {
	URLPattern   string `json:"urlPattern"`
	ResourceType string `json:"resourceType"`
	RequestStage string `json:"requestStage"`
}

// RequestPaused RequestPaused
type RequestPaused struct {
	RequestID           string         `json:"requestId"`
	Request             *Request       `json:"request"`
	FrameID             string         `json:"frameId"`
	ResponseErrorReason *ErrorReason   `json:"responseErrorReason,omitempty"`
	ResponseStatusCode  int            `json:"responseStatusCode,omitempty"`
	ResponseHeaders     []*HeaderEntry `json:"responseHeaders,omitempty"`
	NetworkID           string         `json:"networkId,omitempty"`
}

// RequestPostData https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getRequestPostData
type RequestPostData struct {
	PostData string `json:"postData"`
}

// ResponseBody https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
type ResponseBody struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}
