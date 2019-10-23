package witness

import "github.com/ecwid/witness/pkg/devtool"

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

// SetExtraHTTPHeaders Specifies whether to always send extra HTTP headers with the requests from this page.
func (session *CDPSession) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := session.blockingSend("Network.setExtraHTTPHeaders", Map{"headers": headers})
	return err
}

// SetOffline set offline/online mode
func (session *CDPSession) SetOffline(e bool) error {
	return session.emulateNetworkConditions(e, 0, -1, -1)
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
	_, err := session.blockingSend("Fetch.enable", Map{})
	return err
}

// failRequest https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-failRequest
func (session *CDPSession) failRequest(requestID string, reason devtool.ErrorReason) error {
	_, err := session.blockingSend("Fetch.failRequest", Map{
		"requestId":   requestID,
		"errorReason": string(reason),
	})
	return err
}

// fulfillRequest https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-fulfillRequest
func (session *CDPSession) fulfillRequest(
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

// continueRequest https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-continueRequest
func (session *CDPSession) continueRequest(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error {
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

// Intercepted continue paused request
type Intercepted struct {
	Fail     func(requestID string, reason devtool.ErrorReason) error
	Fulfill  func(requestID string, responseCode int, responseHeaders []*devtool.HeaderEntry, body *string, responsePhrase *string) error
	Continue func(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error
}

// Intercept ...
func (session *CDPSession) Intercept(patterns []*devtool.RequestPattern, fn func(*devtool.RequestPaused, *Intercepted)) func() {
	intercepted := &Intercepted{
		Fail:     session.failRequest,
		Fulfill:  session.fulfillRequest,
		Continue: session.continueRequest,
	}
	unsubscribe := session.subscribe("Fetch.requestPaused", func(msg []byte) {
		request := new(devtool.RequestPaused)
		if err := bytes(msg).Unmarshal(request); err != nil {
			panic(err)
		}
		fn(request, intercepted)
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
