package webaudio

import (
	"github.com/ecwid/control/protocol"
)

/*
	Enables the WebAudio domain and starts sending context lifetime events.
*/
func Enable(c protocol.Caller) error {
	return c.Call("WebAudio.enable", nil, nil)
}

/*
	Disables the WebAudio domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("WebAudio.disable", nil, nil)
}

/*
	Fetch the realtime data from the registered contexts.
*/
func GetRealtimeData(c protocol.Caller, args GetRealtimeDataArgs) (*GetRealtimeDataVal, error) {
	var val = &GetRealtimeDataVal{}
	return val, c.Call("WebAudio.getRealtimeData", args, val)
}
