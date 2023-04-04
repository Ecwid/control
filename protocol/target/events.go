package target

/*
Issued when attached to target because of auto-attach or `attachToTarget` command.
*/
type AttachedToTarget struct {
	SessionId          SessionID   `json:"sessionId"`
	TargetInfo         *TargetInfo `json:"targetInfo"`
	WaitingForDebugger bool        `json:"waitingForDebugger"`
}

/*
	Issued when detached from target for any reason (including `detachFromTarget` command). Can be

issued multiple times per target if multiple sessions have been attached to it.
*/
type DetachedFromTarget struct {
	SessionId SessionID `json:"sessionId"`
}

/*
	Notifies about a new protocol message received from the session (as reported in

`attachedToTarget` event).
*/
type ReceivedMessageFromTarget struct {
	SessionId SessionID `json:"sessionId"`
	Message   string    `json:"message"`
}

/*
Issued when a possible inspection target is created.
*/
type TargetCreated struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}

/*
Issued when a target is destroyed.
*/
type TargetDestroyed struct {
	TargetId TargetID `json:"targetId"`
}

/*
Issued when a target has crashed.
*/
type TargetCrashed struct {
	TargetId  TargetID `json:"targetId"`
	Status    string   `json:"status"`
	ErrorCode int      `json:"errorCode"`
}

/*
	Issued when some information about a target has changed. This only happens between

`targetCreated` and `targetDestroyed`.
*/
type TargetInfoChanged struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}
