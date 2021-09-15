package control

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/transport/observe"
)

func (s *Session) NewNetworkCondition(predicate func(request *network.Request) bool) *Condition {
	var requestID network.RequestId
	return s.NewCondition(func(value observe.Value) (bool, error) {

		switch value.Method {

		case "Network.requestWillBeSent":
			var sent = new(network.RequestWillBeSent)
			if err := json.Unmarshal(value.Params, sent); err != nil {
				return false, err
			}
			if predicate(sent.Request) {
				requestID = sent.RequestId
			}

		case "Network.responseReceived":
			var recv = new(network.ResponseReceived)
			if err := json.Unmarshal(value.Params, recv); err != nil {
				return false, err
			}
			if recv.RequestId == requestID {
				return true, nil
			}

		case "Network.loadingFailed":
			var fail = new(network.LoadingFailed)
			if err := json.Unmarshal(value.Params, fail); err != nil {
				return false, err
			}
			if fail.RequestId == requestID {
				return false, errors.New(fail.ErrorText)
			}
		}

		return false, nil
	})
}

type Network struct {
	s *Session
}

// ClearBrowserCookies ...
func (n Network) ClearBrowserCookies() error {
	return network.ClearBrowserCookies(n.s)
}

// SetCookies ...
func (n Network) SetCookies(cookies ...*network.CookieParam) error {
	return network.SetCookies(n.s, network.SetCookiesArgs{
		Cookies: cookies,
	})
}

// GetCookies returns all browser cookies for the current URL
func (n Network) GetCookies(urls ...string) ([]*network.Cookie, error) {
	val, err := network.GetCookies(n.s, network.GetCookiesArgs{
		Urls: urls,
	})
	if err != nil {
		return nil, err
	}
	return val.Cookies, nil
}

// SetExtraHTTPHeaders Specifies whether to always send extra HTTP headers with the requests from this page.
func (n Network) SetExtraHTTPHeaders(v map[string]string) error {
	val := network.Headers(v)
	return network.SetExtraHTTPHeaders(n.s, network.SetExtraHTTPHeadersArgs{
		Headers: &val,
	})
}

// SetOffline set offline/online mode
// SetOffline(false) - reset all network conditions to default
func (n Network) SetOffline(e bool) error {
	return n.EmulateNetworkConditions(e, 0, -1, -1, ConnectionTypeNone)
}

const (
	ConnectionTypeNone       network.ConnectionType = "none"
	ConnectionTypeCellular2g network.ConnectionType = "cellular2g"
	ConnectionTypeCellular3g network.ConnectionType = "cellular3g"
	ConnectionTypeCellular4g network.ConnectionType = "cellular4g"
	ConnectionTypeBluetooth  network.ConnectionType = "bluetooth"
	ConnectionTypeEthernet   network.ConnectionType = "ethernet"
	ConnectionTypeWIFI       network.ConnectionType = "wifi"
	ConnectionTypeWIMAX      network.ConnectionType = "wimax"
	ConnectionTypeOther      network.ConnectionType = "other"
)

func (n Network) EmulateNetworkConditions(offline bool, latency, downloadThroughput, uploadThroughput float64, connectionType network.ConnectionType) error {
	return network.EmulateNetworkConditions(n.s, network.EmulateNetworkConditionsArgs{
		Offline:            offline,
		Latency:            latency,
		DownloadThroughput: downloadThroughput,
		UploadThroughput:   uploadThroughput,
		ConnectionType:     connectionType,
	})
}

// SetBlockedURLs ...
func (n Network) SetBlockedURLs(urls []string) error {
	return network.SetBlockedURLs(n.s, network.SetBlockedURLsArgs{
		Urls: urls,
	})
}

// GetRequestPostData https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getRequestPostData
func (n Network) GetRequestPostData(requestID network.RequestId) (string, error) {
	val, err := network.GetRequestPostData(n.s, network.GetRequestPostDataArgs{
		RequestId: requestID,
	})
	if err != nil {
		return "", err
	}
	return val.PostData, nil
}

// GetResponseBody https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
func (n Network) GetResponseBody(requestID network.RequestId) (string, error) {
	val, err := network.GetResponseBody(n.s, network.GetResponseBodyArgs{
		RequestId: requestID,
	})
	if err != nil {
		return "", err
	}
	if val.Base64Encoded {
		b, err1 := base64.StdEncoding.DecodeString(val.Body)
		return string(b), err1
	}
	return val.Body, nil
}
