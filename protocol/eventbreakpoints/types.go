package eventbreakpoints

type SetInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"`
}

type RemoveInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"`
}
