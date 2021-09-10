package webaudio

/*
	An unique ID for a graph object (AudioContext, AudioNode, AudioParam) in Web Audio API
*/
type GraphObjectId string

/*
	Enum of BaseAudioContext types
*/
type ContextType string

/*
	Enum of AudioContextState from the spec
*/
type ContextState string

/*
	Enum of AudioNode types
*/
type NodeType string

/*
	Enum of AudioNode::ChannelCountMode from the spec
*/
type ChannelCountMode string

/*
	Enum of AudioNode::ChannelInterpretation from the spec
*/
type ChannelInterpretation string

/*
	Enum of AudioParam types
*/
type ParamType string

/*
	Enum of AudioParam::AutomationRate from the spec
*/
type AutomationRate string

/*
	Fields in AudioContext that change in real-time.
*/
type ContextRealtimeData struct {
	CurrentTime              float64 `json:"currentTime"`
	RenderCapacity           float64 `json:"renderCapacity"`
	CallbackIntervalMean     float64 `json:"callbackIntervalMean"`
	CallbackIntervalVariance float64 `json:"callbackIntervalVariance"`
}

/*
	Protocol object for BaseAudioContext
*/
type BaseAudioContext struct {
	ContextId             GraphObjectId        `json:"contextId"`
	ContextType           ContextType          `json:"contextType"`
	ContextState          ContextState         `json:"contextState"`
	RealtimeData          *ContextRealtimeData `json:"realtimeData,omitempty"`
	CallbackBufferSize    float64              `json:"callbackBufferSize"`
	MaxOutputChannelCount float64              `json:"maxOutputChannelCount"`
	SampleRate            float64              `json:"sampleRate"`
}

/*
	Protocol object for AudioListener
*/
type AudioListener struct {
	ListenerId GraphObjectId `json:"listenerId"`
	ContextId  GraphObjectId `json:"contextId"`
}

/*
	Protocol object for AudioNode
*/
type AudioNode struct {
	NodeId                GraphObjectId         `json:"nodeId"`
	ContextId             GraphObjectId         `json:"contextId"`
	NodeType              NodeType              `json:"nodeType"`
	NumberOfInputs        float64               `json:"numberOfInputs"`
	NumberOfOutputs       float64               `json:"numberOfOutputs"`
	ChannelCount          float64               `json:"channelCount"`
	ChannelCountMode      ChannelCountMode      `json:"channelCountMode"`
	ChannelInterpretation ChannelInterpretation `json:"channelInterpretation"`
}

/*
	Protocol object for AudioParam
*/
type AudioParam struct {
	ParamId      GraphObjectId  `json:"paramId"`
	NodeId       GraphObjectId  `json:"nodeId"`
	ContextId    GraphObjectId  `json:"contextId"`
	ParamType    ParamType      `json:"paramType"`
	Rate         AutomationRate `json:"rate"`
	DefaultValue float64        `json:"defaultValue"`
	MinValue     float64        `json:"minValue"`
	MaxValue     float64        `json:"maxValue"`
}

type GetRealtimeDataArgs struct {
	ContextId GraphObjectId `json:"contextId"`
}

type GetRealtimeDataVal struct {
	RealtimeData *ContextRealtimeData `json:"realtimeData"`
}
