package emulation

import (
	"github.com/ecwid/control/protocol"
)

/*
Tells whether emulation is supported.
*/
func CanEmulate(c protocol.Caller) (*CanEmulateVal, error) {
	var val = &CanEmulateVal{}
	return val, c.Call("Emulation.canEmulate", nil, val)
}

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
	Overrides the values of device screen dimensions (window.screen.width, window.screen.height,

window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media
query results).
*/
func SetDeviceMetricsOverride(c protocol.Caller, args SetDeviceMetricsOverrideArgs) error {
	return c.Call("Emulation.setDeviceMetricsOverride", args, nil)
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
	Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position

unavailable.
*/
func SetGeolocationOverride(c protocol.Caller, args SetGeolocationOverrideArgs) error {
	return c.Call("Emulation.setGeolocationOverride", args, nil)
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
 */
func SetHardwareConcurrencyOverride(c protocol.Caller, args SetHardwareConcurrencyOverrideArgs) error {
	return c.Call("Emulation.setHardwareConcurrencyOverride", args, nil)
}

/*
Allows overriding user agent with the given string.
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
