package control

import (
	"math"

	"github.com/ecwid/control/mobile"
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/emulation"
)

type Emulation struct {
	s *Session
}

// SetDeviceMetricsOverride ...
func (e Emulation) SetDeviceMetricsOverride(metrics emulation.SetDeviceMetricsOverrideArgs) error {
	return emulation.SetDeviceMetricsOverride(e.s, metrics)
}

// SetUserAgentOverride ...
func (e Emulation) SetUserAgentOverride(userAgent, acceptLanguage, platform string, userAgentMetadata *common.UserAgentMetadata) error {
	return emulation.SetUserAgentOverride(e.s, emulation.SetUserAgentOverrideArgs{
		UserAgent:         userAgent,
		AcceptLanguage:    acceptLanguage,
		Platform:          platform,
		UserAgentMetadata: userAgentMetadata,
	})
}

// ClearDeviceMetricsOverride ...
func (e Emulation) ClearDeviceMetricsOverride() error {
	return emulation.ClearDeviceMetricsOverride(e.s)
}

// SetScrollbarsHidden ...
func (e Emulation) SetScrollbarsHidden(hidden bool) error {
	return emulation.SetScrollbarsHidden(e.s, emulation.SetScrollbarsHiddenArgs{
		Hidden: hidden,
	})
}

// SetCPUThrottlingRate https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setCPUThrottlingRate
func (e Emulation) SetCPUThrottlingRate(rate float64) error {
	return emulation.SetCPUThrottlingRate(e.s, emulation.SetCPUThrottlingRateArgs{
		Rate: rate,
	})
}

// SetDocumentCookieDisabled https://chromedevtools.github.io/devtools-protocol/tot/Emulation/#method-setDocumentCookieDisabled
func (e Emulation) SetDocumentCookieDisabled(disabled bool) error {
	return emulation.SetDocumentCookieDisabled(e.s, emulation.SetDocumentCookieDisabledArgs{
		Disabled: disabled,
	})
}

// Emulate emulate predefined device
func (e Emulation) Emulate(device *mobile.Device) error {
	device.Metrics.DontSetVisibleSize = true
	if err := e.SetDeviceMetricsOverride(device.Metrics); err != nil {
		return err
	}
	return e.SetUserAgentOverride(device.UserAgent, "", "", nil)
}

func (e Emulation) FitZoomToWindow() error {
	view, err := e.s.GetLayoutMetrics()
	if err != nil {
		return err
	}

	return e.SetDeviceMetricsOverride(emulation.SetDeviceMetricsOverrideArgs{
		Width:             view.CssLayoutViewport.ClientWidth,
		Height:            int(math.Ceil(view.CssContentSize.Height)),
		DeviceScaleFactor: 1,
		Mobile:            false,
	})
}
