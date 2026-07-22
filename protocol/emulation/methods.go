package emulation

import (
	"github.com/ecwid/control/protocol"
)

/*
Clears the overridden device metrics.
*/
func ClearDeviceMetricsOverride(c protocol.Caller) error {
	return c.Call("Emulation.clearDeviceMetricsOverride", nil, nil)
}

/*
Clears the overridden Geolocation Position and Error.
*/
func ClearGeolocationOverride(c protocol.Caller) error {
	return c.Call("Emulation.clearGeolocationOverride", nil, nil)
}

/*
Requests that page scale factor is reset to initial values.
*/
func ResetPageScaleFactor(c protocol.Caller) error {
	return c.Call("Emulation.resetPageScaleFactor", nil, nil)
}

/*
Enables or disables simulating a focused and active page.
*/
func SetFocusEmulationEnabled(c protocol.Caller, args SetFocusEmulationEnabledArgs) error {
	return c.Call("Emulation.setFocusEmulationEnabled", args, nil)
}

/*
Automatically render all web contents using a dark theme.
*/
func SetAutoDarkModeOverride(c protocol.Caller, args SetAutoDarkModeOverrideArgs) error {
	return c.Call("Emulation.setAutoDarkModeOverride", args, nil)
}

/*
Enables CPU throttling to emulate slow CPUs.
*/
func SetCPUThrottlingRate(c protocol.Caller, args SetCPUThrottlingRateArgs) error {
	return c.Call("Emulation.setCPUThrottlingRate", args, nil)
}

/*
	Sets or clears an override of the default background color of the frame. This override is used

if the content does not specify one.
*/
func SetDefaultBackgroundColorOverride(c protocol.Caller, args SetDefaultBackgroundColorOverrideArgs) error {
	return c.Call("Emulation.setDefaultBackgroundColorOverride", args, nil)
}

/*
	Overrides the values for env(safe-area-inset-*) and env(safe-area-max-inset-*). Unset values will cause the

respective variables to be undefined, even if previously overridden.
*/
func SetSafeAreaInsetsOverride(c protocol.Caller, args SetSafeAreaInsetsOverrideArgs) error {
	return c.Call("Emulation.setSafeAreaInsetsOverride", args, nil)
}

/*
	Overrides the values of device screen dimensions (window.screen.width, window.screen.height,

window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
query results).
*/
func SetDeviceMetricsOverride(c protocol.Caller, args SetDeviceMetricsOverrideArgs) error {
	return c.Call("Emulation.setDeviceMetricsOverride", args, nil)
}

/*
	Start reporting the given posture value to the Device Posture API.

This override can also be set in setDeviceMetricsOverride().
*/
func SetDevicePostureOverride(c protocol.Caller, args SetDevicePostureOverrideArgs) error {
	return c.Call("Emulation.setDevicePostureOverride", args, nil)
}

/*
	Clears a device posture override set with either setDeviceMetricsOverride()

or setDevicePostureOverride() and starts using posture information from the
platform again.
Does nothing if no override is set.
*/
func ClearDevicePostureOverride(c protocol.Caller) error {
	return c.Call("Emulation.clearDevicePostureOverride", nil, nil)
}

/*
	Start using the given display features to pupulate the Viewport Segments API.

This override can also be set in setDeviceMetricsOverride().
*/
func SetDisplayFeaturesOverride(c protocol.Caller, args SetDisplayFeaturesOverrideArgs) error {
	return c.Call("Emulation.setDisplayFeaturesOverride", args, nil)
}

/*
	Clears the display features override set with either setDeviceMetricsOverride()

or setDisplayFeaturesOverride() and starts using display features from the
platform again.
Does nothing if no override is set.
*/
func ClearDisplayFeaturesOverride(c protocol.Caller) error {
	return c.Call("Emulation.clearDisplayFeaturesOverride", nil, nil)
}

/*
 */
func SetScrollbarsHidden(c protocol.Caller, args SetScrollbarsHiddenArgs) error {
	return c.Call("Emulation.setScrollbarsHidden", args, nil)
}

/*
 */
func SetDocumentCookieDisabled(c protocol.Caller, args SetDocumentCookieDisabledArgs) error {
	return c.Call("Emulation.setDocumentCookieDisabled", args, nil)
}

/*
 */
func SetEmitTouchEventsForMouse(c protocol.Caller, args SetEmitTouchEventsForMouseArgs) error {
	return c.Call("Emulation.setEmitTouchEventsForMouse", args, nil)
}

/*
Emulates the given media type or media feature for CSS media queries.
*/
func SetEmulatedMedia(c protocol.Caller, args SetEmulatedMediaArgs) error {
	return c.Call("Emulation.setEmulatedMedia", args, nil)
}

/*
Emulates the given vision deficiency.
*/
func SetEmulatedVisionDeficiency(c protocol.Caller, args SetEmulatedVisionDeficiencyArgs) error {
	return c.Call("Emulation.setEmulatedVisionDeficiency", args, nil)
}

/*
Emulates the given OS text scale.
*/
func SetEmulatedOSTextScale(c protocol.Caller, args SetEmulatedOSTextScaleArgs) error {
	return c.Call("Emulation.setEmulatedOSTextScale", args, nil)
}

/*
	Overrides the Geolocation Position or Error. Omitting latitude, longitude or

accuracy emulates position unavailable.
*/
func SetGeolocationOverride(c protocol.Caller, args SetGeolocationOverrideArgs) error {
	return c.Call("Emulation.setGeolocationOverride", args, nil)
}

/*
 */
func GetOverriddenSensorInformation(c protocol.Caller, args GetOverriddenSensorInformationArgs) (*GetOverriddenSensorInformationVal, error) {
	var val = &GetOverriddenSensorInformationVal{}
	return val, c.Call("Emulation.getOverriddenSensorInformation", args, val)
}

/*
	Overrides a platform sensor of a given type. If |enabled| is true, calls to

Sensor.start() will use a virtual sensor as backend rather than fetching
data from a real hardware sensor. Otherwise, existing virtual
sensor-backend Sensor objects will fire an error event and new calls to
Sensor.start() will attempt to use a real sensor instead.
*/
func SetSensorOverrideEnabled(c protocol.Caller, args SetSensorOverrideEnabledArgs) error {
	return c.Call("Emulation.setSensorOverrideEnabled", args, nil)
}

/*
	Updates the sensor readings reported by a sensor type previously overridden

by setSensorOverrideEnabled.
*/
func SetSensorOverrideReadings(c protocol.Caller, args SetSensorOverrideReadingsArgs) error {
	return c.Call("Emulation.setSensorOverrideReadings", args, nil)
}

/*
	Overrides a pressure source of a given type, as used by the Compute

Pressure API, so that updates to PressureObserver.observe() are provided
via setPressureStateOverride instead of being retrieved from
platform-provided telemetry data.
*/
func SetPressureSourceOverrideEnabled(c protocol.Caller, args SetPressureSourceOverrideEnabledArgs) error {
	return c.Call("Emulation.setPressureSourceOverrideEnabled", args, nil)
}

/*
	Provides a given pressure state that will be processed and eventually be

delivered to PressureObserver users. |source| must have been previously
overridden by setPressureSourceOverrideEnabled.
*/
func SetPressureStateOverride(c protocol.Caller, args SetPressureStateOverrideArgs) error {
	return c.Call("Emulation.setPressureStateOverride", args, nil)
}

/*
Overrides the Idle state.
*/
func SetIdleOverride(c protocol.Caller, args SetIdleOverrideArgs) error {
	return c.Call("Emulation.setIdleOverride", args, nil)
}

/*
Clears Idle state overrides.
*/
func ClearIdleOverride(c protocol.Caller) error {
	return c.Call("Emulation.clearIdleOverride", nil, nil)
}

/*
Sets a specified page scale factor.
*/
func SetPageScaleFactor(c protocol.Caller, args SetPageScaleFactorArgs) error {
	return c.Call("Emulation.setPageScaleFactor", args, nil)
}

/*
Switches script execution in the page.
*/
func SetScriptExecutionDisabled(c protocol.Caller, args SetScriptExecutionDisabledArgs) error {
	return c.Call("Emulation.setScriptExecutionDisabled", args, nil)
}

/*
Enables touch on platforms which do not support them.
*/
func SetTouchEmulationEnabled(c protocol.Caller, args SetTouchEmulationEnabledArgs) error {
	return c.Call("Emulation.setTouchEmulationEnabled", args, nil)
}

/*
	Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets

the current virtual time policy.  Note this supersedes any previous time budget.
*/
func SetVirtualTimePolicy(c protocol.Caller, args SetVirtualTimePolicyArgs) (*SetVirtualTimePolicyVal, error) {
	var val = &SetVirtualTimePolicyVal{}
	return val, c.Call("Emulation.setVirtualTimePolicy", args, val)
}

/*
Overrides default host system locale with the specified one.
*/
func SetLocaleOverride(c protocol.Caller, args SetLocaleOverrideArgs) error {
	return c.Call("Emulation.setLocaleOverride", args, nil)
}

/*
Overrides default host system timezone with the specified one.
*/
func SetTimezoneOverride(c protocol.Caller, args SetTimezoneOverrideArgs) error {
	return c.Call("Emulation.setTimezoneOverride", args, nil)
}

/*
 */
func SetDisabledImageTypes(c protocol.Caller, args SetDisabledImageTypesArgs) error {
	return c.Call("Emulation.setDisabledImageTypes", args, nil)
}

/*
Override the value of navigator.connection.saveData
*/
func SetDataSaverOverride(c protocol.Caller, args SetDataSaverOverrideArgs) error {
	return c.Call("Emulation.setDataSaverOverride", args, nil)
}

/*
 */
func SetHardwareConcurrencyOverride(c protocol.Caller, args SetHardwareConcurrencyOverrideArgs) error {
	return c.Call("Emulation.setHardwareConcurrencyOverride", args, nil)
}

/*
	Allows overriding user agent with the given string.

`userAgentMetadata` must be set for Client Hint headers to be sent.
*/
func SetUserAgentOverride(c protocol.Caller, args SetUserAgentOverrideArgs) error {
	return c.Call("Emulation.setUserAgentOverride", args, nil)
}

/*
Allows overriding the automation flag.
*/
func SetAutomationOverride(c protocol.Caller, args SetAutomationOverrideArgs) error {
	return c.Call("Emulation.setAutomationOverride", args, nil)
}

/*
	Allows overriding the difference between the small and large viewport sizes, which determine the

value of the `svh` and `lvh` unit, respectively. Only supported for top-level frames.
*/
func SetSmallViewportHeightDifferenceOverride(c protocol.Caller, args SetSmallViewportHeightDifferenceOverrideArgs) error {
	return c.Call("Emulation.setSmallViewportHeightDifferenceOverride", args, nil)
}

/*
	Returns device's screen configuration. In headful mode, the physical screens configuration is returned,

whereas in headless mode, a virtual headless screen configuration is provided instead.
*/
func GetScreenInfos(c protocol.Caller) (*GetScreenInfosVal, error) {
	var val = &GetScreenInfosVal{}
	return val, c.Call("Emulation.getScreenInfos", nil, val)
}

/*
Add a new screen to the device. Only supported in headless mode.
*/
func AddScreen(c protocol.Caller, args AddScreenArgs) (*AddScreenVal, error) {
	var val = &AddScreenVal{}
	return val, c.Call("Emulation.addScreen", args, val)
}

/*
Updates specified screen parameters. Only supported in headless mode.
*/
func UpdateScreen(c protocol.Caller, args UpdateScreenArgs) (*UpdateScreenVal, error) {
	var val = &UpdateScreenVal{}
	return val, c.Call("Emulation.updateScreen", args, val)
}

/*
Remove screen from the device. Only supported in headless mode.
*/
func RemoveScreen(c protocol.Caller, args RemoveScreenArgs) error {
	return c.Call("Emulation.removeScreen", args, nil)
}

/*
	Set primary screen. Only supported in headless mode.

Note that this changes the coordinate system origin to the top-left
of the new primary screen, updating the bounds and work areas
of all existing screens accordingly.
*/
func SetPrimaryScreen(c protocol.Caller, args SetPrimaryScreenArgs) error {
	return c.Call("Emulation.setPrimaryScreen", args, nil)
}
