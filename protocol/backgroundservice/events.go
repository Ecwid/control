package backgroundservice

/*
	Called when the recording state for the service has been updated.
*/
type RecordingStateChanged struct {
	IsRecording bool        `json:"isRecording"`
	Service     ServiceName `json:"service"`
}

/*
	Called with all existing backgroundServiceEvents when enabled, and all new
events afterwards if enabled and recording.
*/
type BackgroundServiceEventReceived struct {
	BackgroundServiceEvent *BackgroundServiceEvent `json:"backgroundServiceEvent"`
}
