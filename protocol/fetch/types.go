package fetch

import (
	"github.com/ecwid/control/protocol/io"
	"github.com/ecwid/control/protocol/network"
)

/*
Unique request identifier.
*/
type RequestId string

/*
	Stages of the request to handle. Request will intercept before the request is

sent. Response will intercept after the response is received (but before response
body is received).
*/
type RequestStage string

/*
 */
type RequestPattern struct {
	UrlPattern   string               `json:"urlPattern,omitempty"`
	ResourceType network.ResourceType `json:"resourceType,omitempty"`
	RequestStage RequestStage         `json:"requestStage,omitempty"`
}

/*
Response HTTP header entry
*/
type HeaderEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

type EnableArgs struct {
	Patterns           []*RequestPattern `json:"patterns,omitempty"`
	HandleAuthRequests bool              `json:"handleAuthRequests,omitempty"`
}

type FailRequestArgs struct {
	RequestId   RequestId           `json:"requestId"`
	ErrorReason network.ErrorReason `json:"errorReason"`
}

type FulfillRequestArgs struct {
	RequestId             RequestId      `json:"requestId"`
	ResponseCode          int            `json:"responseCode"`
	ResponseHeaders       []*HeaderEntry `json:"responseHeaders,omitempty"`
	BinaryResponseHeaders []byte         `json:"binaryResponseHeaders,omitempty"`
	Body                  []byte         `json:"body,omitempty"`
	ResponsePhrase        string         `json:"responsePhrase,omitempty"`
}

type ContinueRequestArgs struct {
	RequestId         RequestId      `json:"requestId"`
	Url               string         `json:"url,omitempty"`
	Method            string         `json:"method,omitempty"`
	PostData          []byte         `json:"postData,omitempty"`
	Headers           []*HeaderEntry `json:"headers,omitempty"`
	InterceptResponse bool           `json:"interceptResponse,omitempty"`
}

type ContinueWithAuthArgs struct {
	RequestId             RequestId              `json:"requestId"`
	AuthChallengeResponse *AuthChallengeResponse `json:"authChallengeResponse"`
}

type ContinueResponseArgs struct {
	RequestId             RequestId      `json:"requestId"`
	ResponseCode          int            `json:"responseCode,omitempty"`
	ResponsePhrase        string         `json:"responsePhrase,omitempty"`
	ResponseHeaders       []*HeaderEntry `json:"responseHeaders,omitempty"`
	BinaryResponseHeaders []byte         `json:"binaryResponseHeaders,omitempty"`
}

type GetResponseBodyArgs struct {
	RequestId RequestId `json:"requestId"`
}

type GetResponseBodyVal struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type TakeResponseBodyAsStreamArgs struct {
	RequestId RequestId `json:"requestId"`
}

type TakeResponseBodyAsStreamVal struct {
	Stream io.StreamHandle `json:"stream"`
}
