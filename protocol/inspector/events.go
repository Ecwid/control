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
type TargetCrashed any

/*
Fired when debugging target has reloaded after crash
*/
type TargetReloadedAfterCrash any

/*
Fired on worker targets when main worker script and any imported scripts have been evaluated.
*/
type WorkerScriptLoaded any
