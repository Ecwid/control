package control

import (
	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/page"
)

// CaptureScreenshot get screen of current page
func (s Session) CaptureScreenshot(format string, quality int, clip *page.Viewport, fromSurface, captureBeyondViewport bool) ([]byte, error) {
	val, err := page.CaptureScreenshot(s, page.CaptureScreenshotArgs{
		Format:                format,
		Quality:               quality,
		Clip:                  clip,
		FromSurface:           fromSurface,
		CaptureBeyondViewport: captureBeyondViewport,
	})
	if err != nil {
		return nil, err
	}
	return val.Data, nil
}

// AddScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-addScriptToEvaluateOnNewDocument
func (s Session) AddScriptToEvaluateOnNewDocument(source string) (page.ScriptIdentifier, error) {
	val, err := page.AddScriptToEvaluateOnNewDocument(s, page.AddScriptToEvaluateOnNewDocumentArgs{
		Source: source,
	})
	if err != nil {
		return "", err
	}
	return val.Identifier, nil
}

// RemoveScriptToEvaluateOnNewDocument https://chromedevtools.github.io/devtools-protocol/tot/Page#method-removeScriptToEvaluateOnNewDocument
func (s Session) RemoveScriptToEvaluateOnNewDocument(identifier page.ScriptIdentifier) error {
	return page.RemoveScriptToEvaluateOnNewDocument(s, page.RemoveScriptToEvaluateOnNewDocumentArgs{
		Identifier: identifier,
	})
}

// SetDownloadBehavior https://chromedevtools.github.io/devtools-protocol/tot/Page#method-setDownloadBehavior
func (s Session) SetDownloadBehavior(behavior string, downloadPath string, eventsEnabled bool) error {
	return browser.SetDownloadBehavior(s, browser.SetDownloadBehaviorArgs{
		Behavior:      behavior,
		DownloadPath:  downloadPath,
		EventsEnabled: eventsEnabled, // default false
	})
}

// HandleJavaScriptDialog ...
func (s Session) HandleJavaScriptDialog(accept bool, promptText string) error {
	return page.HandleJavaScriptDialog(s, page.HandleJavaScriptDialogArgs{
		Accept:     accept,
		PromptText: promptText,
	})
}

func (s Session) GetLayoutMetrics() (*page.GetLayoutMetricsVal, error) {
	view, err := page.GetLayoutMetrics(s)
	if err != nil {
		return nil, err
	}
	return view, nil
}
