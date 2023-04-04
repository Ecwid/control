package deviceorientation

import (
	"github.com/ecwid/control/protocol"
)

/*
Clears the overridden Device Orientation.
*/
func ClearDeviceOrientationOverride(c protocol.Caller) error {
	return c.Call("DeviceOrientation.clearDeviceOrientationOverride", nil, nil)
}

/*
Overrides the Device Orientation.
*/
func SetDeviceOrientationOverride(c protocol.Caller, args SetDeviceOrientationOverrideArgs) error {
	return c.Call("DeviceOrientation.setDeviceOrientationOverride", args, nil)
}
