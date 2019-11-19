package mobile

import "github.com/ecwid/witness/pkg/devtool"

// Device device description
type Device struct {
	Metrics   *devtool.DeviceMetrics
	UserAgent string
}

var (
	landscape = &devtool.ScreenOrientation{Type: devtool.LandscapePrimary, Angle: 90}
	portrait  = &devtool.ScreenOrientation{Type: devtool.PortraitPrimary, Angle: 0}
)

var (
	iphoneUA = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	ipadUA   = "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1"
)

// Rotated rotate copy of current device and return new device
func (d Device) Rotated() *Device {
	nd := &Device{UserAgent: d.UserAgent, Metrics: &devtool.DeviceMetrics{}}
	*nd.Metrics = *d.Metrics
	switch d.Metrics.ScreenOrientation {
	case portrait:
		nd.Metrics.ScreenOrientation = landscape
	case landscape:
		nd.Metrics.ScreenOrientation = portrait
	}
	nd.Metrics.Width, nd.Metrics.Height = d.Metrics.Height, d.Metrics.Width
	return nd
}

// Predefined devices
var (
	f        = true
	GalaxyS5 = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             360,
			Height:            640,
			DeviceScaleFactor: 3,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36",
	}

	Pixel2 = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             411,
			Height:            731,
			DeviceScaleFactor: 2.625,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36",
	}

	Pixel2XL = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             411,
			Height:            823,
			DeviceScaleFactor: 3.5,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36",
	}

	IPad = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             768,
			Height:            1024,
			DeviceScaleFactor: 2,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: ipadUA,
	}

	IPadMini = IPad

	IPadPro = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             1024,
			Height:            1366,
			DeviceScaleFactor: 2,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: ipadUA,
	}

	IPhone6 = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             375,
			Height:            667,
			DeviceScaleFactor: 2,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: iphoneUA,
	}
	IPhone7 = IPhone6
	IPhone8 = IPhone6

	IPhone6Plus = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             414,
			Height:            736,
			DeviceScaleFactor: 3,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: iphoneUA,
	}
	IPhone7Plus = IPhone6Plus
	IPhone8Plus = IPhone6Plus

	IPhoneX = &Device{
		Metrics: &devtool.DeviceMetrics{
			Width:             375,
			Height:            812,
			DeviceScaleFactor: 3,
			Mobile:            true,
			ScreenOrientation: portrait,
		},
		UserAgent: iphoneUA,
	}
)
