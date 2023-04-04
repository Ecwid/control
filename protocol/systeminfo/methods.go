package systeminfo

import (
	"github.com/ecwid/control/protocol"
)

/*
Returns information about the system.
*/
func GetInfo(c protocol.Caller) (*GetInfoVal, error) {
	var val = &GetInfoVal{}
	return val, c.Call("SystemInfo.getInfo", nil, val)
}

/*
Returns information about the feature state.
*/
func GetFeatureState(c protocol.Caller, args GetFeatureStateArgs) (*GetFeatureStateVal, error) {
	var val = &GetFeatureStateVal{}
	return val, c.Call("SystemInfo.getFeatureState", args, val)
}

/*
Returns information about all running processes.
*/
func GetProcessInfo(c protocol.Caller) (*GetProcessInfoVal, error) {
	var val = &GetProcessInfoVal{}
	return val, c.Call("SystemInfo.getProcessInfo", nil, val)
}
