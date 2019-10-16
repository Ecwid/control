package witness

import "github.com/ecwid/witness/pkg/devtool"

// ClearBrowserCookies ...
func (session *Session) ClearBrowserCookies() error {
	_, err := session.blockingSend("Network.clearBrowserCookies", Map{})
	return err
}

// SetCookies ...
func (session *Session) SetCookies(cookies ...*devtool.Cookie) error {
	_, err := session.blockingSend("Network.setCookies", Map{"cookies": cookies})
	return err
}

// SetExtraHTTPHeaders Specifies whether to always send extra HTTP headers with the requests from this page.
func (session *Session) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := session.blockingSend("Network.setExtraHTTPHeaders", Map{"headers": headers})
	return err
}

// fetchEnable https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-enable
func (session *Session) fetchEnable(patterns []*devtool.RequestPattern, handleAuthRequests bool) error {
	_, err := session.blockingSend("Fetch.enable", Map{
		"patterns":           patterns,
		"handleAuthRequests": handleAuthRequests,
	})
	return err
}

// fetchDisable https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-disable
func (session *Session) fetchDisable() error {
	_, err := session.blockingSend("Fetch.enable", Map{})
	return err
}

// failRequest https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-failRequest
func (session *Session) failRequest(requestID string, reason devtool.ErrorReason) error {
	_, err := session.blockingSend("Fetch.failRequest", Map{
		"requestId":   requestID,
		"errorReason": string(reason),
	})
	return err
}

// fulfillRequest https://chromedevtools.github.io/devtools-protocol/tot/Fetch#method-fulfillRequest
func (session *Session) fulfillRequest(
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
func (session *Session) continueRequest(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error {
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

// Proceed continue paused request
type Proceed struct {
	Fail     func(requestID string, reason devtool.ErrorReason) error
	Fulfill  func(requestID string, responseCode int, responseHeaders []*devtool.HeaderEntry, body *string, responsePhrase *string) error
	Continue func(requestID string, url *string, method *string, postData *string, headers []*devtool.HeaderEntry) error
}

// Fetch ...
func (session *Session) Fetch(patterns []*devtool.RequestPattern, fn func(*devtool.RequestPaused, *Proceed)) func() {
	proceed := &Proceed{
		Fail:     session.failRequest,
		Fulfill:  session.fulfillRequest,
		Continue: session.continueRequest,
	}
	unsubscribe := session.subscribe("Fetch.requestPaused", func(msg []byte) {
		request := new(devtool.RequestPaused)
		bytes(msg).Unmarshal(request)
		fn(request, proceed)
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
