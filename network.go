package witness

import (
	"encoding/base64"
	"encoding/json"

	"github.com/ecwid/witness/pkg/devtool"
)

// ClearBrowserCookies ...
func (session *CDPSession) ClearBrowserCookies() error {
	_, err := session.blockingSend("Network.clearBrowserCookies", Map{})
	return err
}

// SetCookies ...
func (session *CDPSession) SetCookies(cookies ...*devtool.Cookie) error {
	_, err := session.blockingSend("Network.setCookies", Map{"cookies": cookies})
	return err
}

// GetCookies returns all browser cookies for the current URL
func (session *CDPSession) GetCookies(urls ...string) ([]*devtool.Cookie, error) {
	cmap := Map{}
	if urls != nil {
		cmap["urls"] = urls
	}
	getCookies, err := session.blockingSend("Network.getCookies", cmap)
	if err != nil {
		return nil, err
	}
	cookies := new(devtool.GetCookies)
	if err := json.Unmarshal(getCookies, cookies); err != nil {
		return nil, err
	}
	return cookies.Cookies, nil
}

// SetExtraHTTPHeaders Specifies whether to always send extra HTTP headers with the requests from this page.
func (session *CDPSession) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := session.blockingSend("Network.setExtraHTTPHeaders", Map{"headers": headers})
	return err
}

// SetOffline set offline/online mode
// SetOffline(false) - reset all network conditions to default
func (session *CDPSession) SetOffline(e bool) error {
	return session.emulateNetworkConditions(e, 0, -1, -1)
}

// SetThrottling set latency in milliseconds, download & upload throttling in bytes per second
func (session *CDPSession) SetThrottling(latencyMs, downloadThroughputBps, uploadThroughputBps int) error {
	return session.emulateNetworkConditions(false, latencyMs, downloadThroughputBps, downloadThroughputBps)
}

// SetBlockedURLs blocking some resources
func (session *CDPSession) SetBlockedURLs(urls []string) error {
	_, err := session.blockingSend("Network.setBlockedURLs", Map{
		"urls": urls,
	})
	return err
}

// GetRequestPostData https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getRequestPostData
func (session *CDPSession) GetRequestPostData(requestID string) (string, error) {
	body, err := session.blockingSend("Network.getRequestPostData", Map{
		"requestId": requestID,
	})
	if err != nil {
		return "", err
	}
	result := new(devtool.RequestPostData)
	err = body.Unmarshal(result)
	return result.PostData, err
}

// GetResponseBody https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-getResponseBody
func (session *CDPSession) GetResponseBody(requestID string) (string, error) {
	body, err := session.blockingSend("Network.getResponseBody", Map{
		"requestId": requestID,
	})
	if err != nil {
		return "", err
	}
	result := new(devtool.ResponseBody)
	if err = body.Unmarshal(result); err != nil {
		return "", err
	}
	if result.Base64Encoded {
		b, err := base64.StdEncoding.DecodeString(result.Body)
		return string(b), err
	}
	return result.Body, nil
}

func (session *CDPSession) emulateNetworkConditions(offline bool, latencyMs, downloadThroughputBps, uploadThroughputBps int) error {
	_, err := session.blockingSend("Network.emulateNetworkConditions", Map{
		"offline":            offline,
		"latency":            latencyMs,
		"downloadThroughput": downloadThroughputBps,
		"uploadThroughput":   uploadThroughputBps,
	})
	return err
}

// fetchEnable https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-enable
func (session *CDPSession) fetchEnable(patterns []*devtool.RequestPattern, handleAuthRequests bool) error {
	_, err := session.blockingSend("Fetch.enable", Map{
		"patterns":           patterns,
		"handleAuthRequests": handleAuthRequests,
	})
	return err
}

// fetchDisable https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-disable
func (session *CDPSession) fetchDisable() error {
	_, err := session.blockingSend("Fetch.disable", Map{})
	return err
}

// Fail https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-failRequest
func (session *CDPSession) Fail(requestID string, reason devtool.ErrorReason) error {
	_, err := session.blockingSend("Fetch.failRequest", Map{
		"requestId":   requestID,
		"errorReason": string(reason),
	})
	return err
}

// Fulfill https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-fulfillRequest
func (session *CDPSession) Fulfill(
	requestID string,
	responseCode int,
	responseHeaders []*devtool.HeaderEntry,
	body *string,
	responsePhrase *string) error {

	p := Map{
		"requestId":       requestID,
		"responseCode":    responseCode,
		"responseHeaders": responseHeaders,
		"body":            body,
		"responsePhrase":  responsePhrase,
	}
	p.omitempty()
	_, err := session.blockingSend("Fetch.fulfillRequest", p)
	return err
}

// Continue https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-continueRequest
func (session *CDPSession) Continue(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error {
	p := Map{
		"requestId": requestID,
		"url":       url,
		"method":    method,
		"postData":  postData,
		"headers":   headers,
	}
	p.omitempty()
	_, err := session.blockingSend("Fetch.continueRequest", p)
	return err
}

// Interceptor continue paused request
type Interceptor interface {
	Fail(requestID string, reason devtool.ErrorReason) error
	Fulfill(requestID string, responseCode int, responseHeaders []*devtool.HeaderEntry, body *string, responsePhrase *string) error
	Continue(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error
}

// Intercept ...
func (session *CDPSession) Intercept(patterns []*devtool.RequestPattern, fn func(*devtool.RequestPaused, Interceptor)) func() {
	unsubscribe := session.subscribe("Fetch.requestPaused", func(e *Event) {
		request := new(devtool.RequestPaused)
		if err := json.Unmarshal(e.Params, request); err != nil {
			session.panic(err)
			return
		}
		go fn(request, session)
	})
	if err := session.fetchEnable(patterns, false); err != nil {
		panic(err)
	}
	return func() {
		unsubscribe()
		if err := session.fetchDisable(); err != nil {
			panic(err)
		}
	}
}
