package har

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"time"
)

/*
HAR object represents the root of the exported data. This object MUST be
present and its name MUST be "log". The object contains the following
name/value pairs:
*/
type HAR struct {

	// Required. Version number of the format.
	Version string `json:"version"`

	// Required. An object of type creator that contains the name and version
	// information of the log creator application.
	Creator *Creator `json:"creator"`

	// Optional. An object of type browser that contains the name and version
	// information of the user agent.
	Browser *Browser `json:"browser,omitempty"`

	// Optional. An array of objects of type page, each representing one exported
	// (tracked) page. Leave out this field if the application does not support
	// grouping by pages.
	Pages []*Page `json:"pages,omitempty"`

	// Required. An array of objects of type entry, each representing one
	// exported (tracked) HTTP request.
	Entries []*Entry `json:"entries"`

	// Optional. A comment provided by the user or the application. Sorting
	// entries by startedDateTime (starting from the oldest) is preferred way how
	// to export data since it can make importing faster. However the reader
	// application should always make sure the array is sorted (if required for
	// the import).
	Comment string `json:"comment,omitempty"`
}

/*
Creator object contains information about the log creator application and
contains the following name/value pairs:
*/
type Creator struct {
	// Required. The name of the application that created the log.
	Name string `json:"name"`
	// Required. The version number of the application that created the log.
	Version string `json:"version"`
	// Optional. A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

/*
Browser object contains information about the browser that created the log
and contains the following name/value pairs:
*/
type Browser struct {
	// Required. The name of the browser that created the log.
	Name string `json:"name"`
	// Required. The version number of the browser that created the log.
	Version string `json:"version"`
	// Optional. A comment provided by the user or the browser.
	Comment string `json:"comment,omitempty"`
}

/*
Page is one <page> object for every exported web page and one <entry>
object for every HTTP request. In case when an HTTP trace tool isn't able to
group requests by a page, the <pages> object is empty and individual
requests doesn't have a parent page.
*/
type Page struct {
	// Date and time stamp for the beginning of the page load
	// (ISO 8601 YYYY-MM-DDThh:mm:ss.sTZD, e.g. 2009-07-24T19:20:30.45+01:00).
	StartedDateTime time.Time `json:"startedDateTime"`
	// Unique identifier of a page within the . Entries use it to refer the parent page.
	ID string `json:"id"`
	// Page title.
	Title string `json:"title"`
	// Detailed timing info about page load.
	PageTiming *PageTiming `json:"pageTimings"`
	// (new in 1.2) A comment provided by the user or the application.
	Comment   string  `json:"comment,omitempty"`
	Timestamp float64 `json:"-"`
}

/*
PageTiming object describes timings for various events (states) fired during the
page load. All times are specified in milliseconds. If a time info is not
available appropriate field is set to -1.
*/
type PageTiming struct {
	// Content of the page loaded. Number of milliseconds since page load started
	// (page.startedDateTime). Use -1 if the timing does not apply to the current
	// request.
	// Depeding on the browser, onContentLoad property represents DOMContentLoad
	// event or document.readyState == interactive.
	OnContentLoad int64 `json:"onContentLoad"`
	// Page is loaded (onLoad event fired). Number of milliseconds since page
	// load started (page.startedDateTime). Use -1 if the timing does not apply
	// to the current request.
	OnLoad int64 `json:"onLoad"`
	// (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// Entry ...
type Entry struct {
	// Unique, optional Reference to the parent page. Leave out this field if
	// the application does not support grouping by pages.
	Pageref string `json:"pageref"`
	// Date and time stamp of the request start
	// (ISO 8601 YYYY-MM-DDThh:mm:ss.sTZD).
	StartedDateTime time.Time `json:"startedDateTime"`
	// Total elapsed time of the request in milliseconds. This is the sum of all
	// timings available in the timings object (i.e. not including -1 values) .
	Time float64 `json:"time"`
	// Detailed info about the request.
	Request *Request `json:"request"`
	// Detailed info about the response.
	Response *Response `json:"response"`
	// Info about cache usage.
	Cache *Cache `json:"cache"`
	// Detailed timing info about request/response round trip.
	PageTimings *PageTimings `json:"timings"`
	// optional (new in 1.2) IP address of the server that was connected
	// (result of DNS resolution).
	ServerIPAddress string `json:"serverIPAddress,omitempty"`
	// optional (new in 1.2) Unique ID of the parent TCP/IP connection, can be
	// the client port number. Note that a port number doesn't have to be unique
	// identifier in cases where the port is shared for more connections. If the
	// port isn't available for the application, any other unique connection ID
	// can be used instead (e.g. connection index). Leave out this field if the
	// application doesn't support this info.
	Connection string `json:"connection,omitempty"`
	// (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// Request object contains detailed info about performed request.
type Request struct {
	ID string `json:"-"`
	// Request method (GET, POST, ...).
	Method string `json:"method"`
	// Absolute URL of the request (fragments are not included).
	URL string `json:"url"`
	// Request HTTP Version.
	HTTPVersion string `json:"httpVersion"`
	// List of cookie objects.
	Cookies []*Cookie `json:"cookies"`
	// List of header objects.
	Headers []*NVP `json:"headers"`
	// List of query parameter objects.
	QueryString []*NVP `json:"queryString"`
	// Posted data.
	PostData *PostData `json:"postData"`
	// Total number of bytes from the start of the HTTP request message until
	// (and including) the double CRLF before the body. Set to -1 if the info
	// is not available.
	HeaderSize int `json:"headerSize"`
	// Size of the request body (POST data payload) in bytes. Set to -1 if the
	// info is not available.
	BodySize int `json:"bodySize"`
	// (new in 1.2) A comment provided by the user or the application.
	Comment   string  `json:"comment,omitempty"`
	Timestamp float64 `json:"-"`
}

// SetHeadersSize calculate request's headers size
func (req *Request) SetHeadersSize() error {
	var b bytes.Buffer
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		return err
	}
	b.Write([]byte(fmt.Sprintf("%s %s %s\r\n", req.Method, reqURL.RequestURI(), req.HTTPVersion)))
	for _, h := range req.Headers {
		b.Write([]byte(fmt.Sprintf("%s: %s\r\n", h.Name, h.Value)))
	}
	b.Write([]byte("\r\n"))
	req.HeaderSize = b.Len()
	return nil
}

// SetCookies set cookies from headers
func (req *Request) SetCookies() {
	req.Cookies = make([]*Cookie, 0)
	for _, h := range req.Headers {
		if h.Name == "Cookie" {
			cookies := strings.Split(h.Value, ";")
			for _, c := range cookies {
				cookie := strings.TrimSpace(c)
				parts := strings.SplitN(cookie, "=", 2)
				req.Cookies = append(req.Cookies, &Cookie{Name: parts[0], Value: parts[1]})
			}
		}
	}
}

// Response object contains detailed info about the response.
type Response struct {
	// Response status.
	// Status string
	Status int `json:"status"`
	// Response status description.
	StatusText string `json:"statusText"`
	// Response HTTP Version.
	HTTPVersion string `json:"httpVersion"`
	// List of cookie objects.
	Cookies []*Cookie `json:"cookies"`
	// List of header objects.
	Headers []*NVP `json:"headers"`
	// Details about the response body.
	Content *Content `json:"content"`
	// Redirection target URL from the Location response header.
	RedirectURL string `json:"redirectURL"`
	// Total number of bytes from the start of the HTTP response message until
	// (and including) the double CRLF before the body. Set to -1 if the info is
	// not available.
	// The size of received response-headers is computed only from headers that
	// are really received from the server. Additional headers appended by the
	// browser are not included in this number, but they appear in the list of
	// header objects.
	HeadersSize int `json:"headersSize"`
	// Size of the received response body in bytes. Set to zero in case of
	// responses coming from the cache (304). Set to -1 if the info is not
	// available.
	BodySize int64 `json:"bodySize"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment   string  `json:"comment,omitempty"`
	Timestamp float64 `json:"-"`
}

// SetHeadersSize calculate response's headers size
func (resp *Response) SetHeadersSize() {
	var b bytes.Buffer
	b.Write([]byte(fmt.Sprintf("%s %d %s\r\n", resp.HTTPVersion, resp.Status, resp.StatusText)))
	for _, h := range resp.Headers {
		b.Write([]byte(fmt.Sprintf("%s: %s\r\n", h.Name, h.Value)))
	}
	b.Write([]byte("\r\n"))
	resp.HeadersSize = b.Len()
}

// Cookie object contains list of all cookies (used in <request> and <response> objects).
type Cookie struct {
	// The name of the cookie.
	Name string `json:"name"`
	// The cookie value.
	Value string `json:"value"`
	// optional The path pertaining to the cookie.
	Path string `json:"path,omitempty"`
	// optional The host of the cookie.
	Domain string `json:"domain,omitempty"`
	// optional Cookie expiration time.
	// (ISO 8601 YYYY-MM-DDThh:mm:ss.sTZD, e.g. 2009-07-24T19:20:30.123+02:00).
	Expires string `json:"expires,omitempty"`
	// optional Set to true if the cookie is HTTP only, false otherwise.
	HTTPOnly string `json:"httpOnly,omitempty"`
	// optional (new in 1.2) True if the cookie was transmitted over ssl, false
	// otherwise.
	Secure bool `json:"secure,omitempty"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment bool `json:"comment,omitempty"`
}

// NVP is simply a name/value pair with a comment
type NVP struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

// PostData object describes posted data, if any (embedded in <request> object).
type PostData struct {
	//  Mime type of posted data.
	MimeType string `json:"mimeType"`
	//  List of posted parameters (in case of URL encoded parameters).
	Params []*PostParam `json:"params"`
	//  Plain text posted data
	Text string `json:"text"`
	// optional (new in 1.2) A comment provided by the user or the
	// application.
	Comment string `json:"comment,omitempty"`
}

// PostParam List of posted parameters, if any (embedded in <postData> object).
type PostParam struct {
	// name of a posted parameter.
	Name string `json:"name"`
	// optional value of a posted parameter or content of a posted file.
	Value string `json:"value,omitempty"`
	// optional name of a posted file.
	FileName string `json:"fileName,omitempty"`
	// optional content type of a posted file.
	ContentType string `json:"contentType,omitempty"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// Content object describes details about response content (embedded in <response> object).
type Content struct {
	// Length of the returned content in bytes. Should be equal to
	// response.bodySize if there is no compression and bigger when the content
	// has been compressed.
	Size int64 `json:"size"`
	// optional Number of bytes saved. Leave out this field if the information
	// is not available.
	Compression int64 `json:"compression,omitempty"`
	// MIME type of the response text (value of the Content-Type response
	// header). The charset attribute of the MIME type is included (if
	// available).
	MimeType string `json:"mimeType"`
	// optional Response body sent from the server or loaded from the browser
	// cache. This field is populated with textual content only. The text field
	// is either HTTP decoded text or a encoded (e.g. "base64") representation of
	// the response body. Leave out this field if the information is not
	// available.
	Text string `json:"text,omitempty"`
	// optional (new in 1.2) Encoding used for response text field e.g
	// "base64". Leave out this field if the text field is HTTP decoded
	// (decompressed & unchunked), than trans-coded from its original character
	// set into UTF-8.
	Encoding string `json:"encoding,omitempty"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// Cache objects contains info about a request coming from browser cache.
type Cache struct {
	// optional State of a cache entry before the request. Leave out this field
	// if the information is not available.
	BeforeRequest *CacheObject `json:"beforeRequest,omitempty"`
	// optional State of a cache entry after the request. Leave out this field if
	// the information is not available.
	AfterRequest *CacheObject `json:"afterRequest,omitempty"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// CacheObject Both beforeRequest and afterRequest object share the following structure.
type CacheObject struct {
	// optional - Expiration time of the cache entry.
	Expires string `json:"expires,omitempty"`
	// The last time the cache entry was opened.
	LastAccess string `json:"lastAccess"`
	// Etag
	ETag string `json:"etag"`
	// The number of times the cache entry has been opened.
	HitCount int `json:"hitCount"`
	// optional (new in 1.2) A comment provided by the user or the application.
	Comment string `json:"comment,omitempty"`
}

// PageTimings object describes various phases within request-response round trip. All times are specified in milliseconds.
type PageTimings struct {
	Blocked int `json:"blocked,omitempty"`
	// optional - Time spent in a queue waiting for a network connection. Use -1
	// if the timing does not apply to the current request.
	DNS int `json:"dns,omitempty"`
	// optional - DNS resolution time. The time required to resolve a host name.
	// Use -1 if the timing does not apply to the current request.
	Connect int `json:"connect,omitempty"`
	// optional - Time required to create TCP connection. Use -1 if the timing
	// does not apply to the current request.
	Send int `json:"send,omitempty"`
	// Time required to send HTTP request to the server.
	Wait int `json:"wait,omitempty"`
	// Waiting for a response from the server.
	Receive int `json:"receive,omitempty"`
	// Time required to read entire response from the server (or cache).
	SSL int `json:"ssl,omitempty"`
	// optional (new in 1.2) - Time required for SSL/TLS negotiation. If this
	// field is defined then the time is also included in the connect field (to
	// ensure backward compatibility with HAR 1.1). Use -1 if the timing does not
	// apply to the current request.
	Comment string `json:"comment,omitempty"`
	// optional (new in 1.2) - A comment provided by the user or the application.
}
