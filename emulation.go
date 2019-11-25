package witness

import (
	"github.com/ecwid/witness/pkg/devtool"
	"github.com/ecwid/witness/pkg/mobile"
)

func (session *CDPSession) setDeviceMetricsOverride(metrics *devtool.DeviceMetrics) error {
	_, err := session.blockingSend("Emulation.setDeviceMetricsOverride", metrics)
	return err
}

// SetUserAgent set user agent
func (session *CDPSession) SetUserAgent(userAgent string) error {
	return session.setUserAgent(userAgent, nil, nil)
}

func (session *CDPSession) setUserAgent(userAgent string, acceptLanguage, platform *string) error {
	p := Map{
		"userAgent":      userAgent,
		"acceptLanguage": acceptLanguage,
		"platform":       platform,
	}
	p.omitempty()
	_, err := session.blockingSend("Emulation.setUserAgentOverride", p)
	return err
}

func (session *CDPSession) clearDeviceMetricsOverride() error {
	_, err := session.blockingSend("Emulation.clearDeviceMetricsOverride", Map{})
	return err
}

func (session *CDPSession) setScrollbarsHidden(hidden bool) error {
	_, err := session.blockingSend("Emulation.setScrollbarsHidden", Map{"hidden": hidden})
	return err
}

// SetCPUThrottlingRate https://chromedevtools.github.io/devtools-protocol/tot/Emulation#method-setCPUThrottlingRate
func (session *CDPSession) SetCPUThrottlingRate(rate int) error {
	_, err := session.blockingSend("Emulation.setCPUThrottlingRate", Map{"rate": rate})
	return err
}

// Emulate emulate predefined device
func (session *CDPSession) Emulate(device *mobile.Device) error {
	f := true
	device.Metrics.DontSetVisibleSize = &f
	if err := session.setDeviceMetricsOverride(device.Metrics); err != nil {
		return err
	}
	return session.setUserAgent(device.UserAgent, nil, nil)
}
