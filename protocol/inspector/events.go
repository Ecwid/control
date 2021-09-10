package inspector

/*
	Fired when remote debugging connection is about to be terminated. Contains detach reason.
*/
type Detached struct {
	Reason string `json:"reason"`
}

/*
	Fired when debugging target has crashed
*/
type TargetCrashed interface{}

/*
	Fired when debugging target has reloaded after crash
*/
type TargetReloadedAfterCrash interface{}
