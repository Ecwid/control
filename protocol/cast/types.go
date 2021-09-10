package cast

/*

 */
type Sink struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Session string `json:"session,omitempty"`
}

type EnableArgs struct {
	PresentationUrl string `json:"presentationUrl,omitempty"`
}

type SetSinkToUseArgs struct {
	SinkName string `json:"sinkName"`
}

type StartTabMirroringArgs struct {
	SinkName string `json:"sinkName"`
}

type StopCastingArgs struct {
	SinkName string `json:"sinkName"`
}
