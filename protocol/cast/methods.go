package cast

import (
	"github.com/ecwid/control/protocol"
)

/*
	Starts observing for sinks that can be used for tab mirroring, and if set,

sinks compatible with |presentationUrl| as well. When sinks are found, a
|sinksUpdated| event is fired.
Also starts observing for issue messages. When an issue is added or removed,
an |issueUpdated| event is fired.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("Cast.enable", args, nil)
}

/*
Stops observing for sinks and issues.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Cast.disable", nil, nil)
}

/*
	Sets a sink to be used when the web page requests the browser to choose a

sink via Presentation API, Remote Playback API, or Cast SDK.
*/
func SetSinkToUse(c protocol.Caller, args SetSinkToUseArgs) error {
	return c.Call("Cast.setSinkToUse", args, nil)
}

/*
Starts mirroring the desktop to the sink.
*/
func StartDesktopMirroring(c protocol.Caller, args StartDesktopMirroringArgs) error {
	return c.Call("Cast.startDesktopMirroring", args, nil)
}

/*
Starts mirroring the tab to the sink.
*/
func StartTabMirroring(c protocol.Caller, args StartTabMirroringArgs) error {
	return c.Call("Cast.startTabMirroring", args, nil)
}

/*
Stops the active Cast session on the sink.
*/
func StopCasting(c protocol.Caller, args StopCastingArgs) error {
	return c.Call("Cast.stopCasting", args, nil)
}
