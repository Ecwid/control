package fetch

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
)

/*
	Issued when the domain is enabled and the request URL matches the
specified filter. The request is paused until the client responds
with one of continueRequest, failRequest or fulfillRequest.
The stage of the request can be determined by presence of responseErrorReason
and responseStatusCode -- the request is at the response stage if either
of these fields is present and in the request stage otherwise.
*/
type RequestPaused struct {
	RequestId           RequestId            `json:"requestId"`
	Request             *network.Request     `json:"request"`
	FrameId             common.FrameId       `json:"frameId"`
	ResourceType        network.ResourceType `json:"resourceType"`
	ResponseErrorReason network.ErrorReason  `json:"responseErrorReason,omitempty"`
	ResponseStatusCode  int                  `json:"responseStatusCode,omitempty"`
	ResponseHeaders     []*HeaderEntry       `json:"responseHeaders,omitempty"`
	NetworkId           RequestId            `json:"networkId,omitempty"`
}

/*
	Issued when the domain is enabled with handleAuthRequests set to true.
The request is paused until client responds with continueWithAuth.
*/
type AuthRequired struct {
	RequestId     RequestId            `json:"requestId"`
	Request       *network.Request     `json:"request"`
	FrameId       common.FrameId       `json:"frameId"`
	ResourceType  network.ResourceType `json:"resourceType"`
	AuthChallenge *AuthChallenge       `json:"authChallenge"`
}
