package webaudio

/*
Notifies that a new BaseAudioContext has been created.
*/
type ContextCreated struct {
	Context *BaseAudioContext `json:"context"`
}

/*
Notifies that an existing BaseAudioContext will be destroyed.
*/
type ContextWillBeDestroyed struct {
	ContextId GraphObjectId `json:"contextId"`
}

/*
Notifies that existing BaseAudioContext has changed some properties (id stays the same)..
*/
type ContextChanged struct {
	Context *BaseAudioContext `json:"context"`
}

/*
Notifies that the construction of an AudioListener has finished.
*/
type AudioListenerCreated struct {
	Listener *AudioListener `json:"listener"`
}

/*
Notifies that a new AudioListener has been created.
*/
type AudioListenerWillBeDestroyed struct {
	ContextId  GraphObjectId `json:"contextId"`
	ListenerId GraphObjectId `json:"listenerId"`
}

/*
Notifies that a new AudioNode has been created.
*/
type AudioNodeCreated struct {
	Node *AudioNode `json:"node"`
}

/*
Notifies that an existing AudioNode has been destroyed.
*/
type AudioNodeWillBeDestroyed struct {
	ContextId GraphObjectId `json:"contextId"`
	NodeId    GraphObjectId `json:"nodeId"`
}

/*
Notifies that a new AudioParam has been created.
*/
type AudioParamCreated struct {
	Param *AudioParam `json:"param"`
}

/*
Notifies that an existing AudioParam has been destroyed.
*/
type AudioParamWillBeDestroyed struct {
	ContextId GraphObjectId `json:"contextId"`
	NodeId    GraphObjectId `json:"nodeId"`
	ParamId   GraphObjectId `json:"paramId"`
}

/*
Notifies that two AudioNodes are connected.
*/
type NodesConnected struct {
	ContextId             GraphObjectId `json:"contextId"`
	SourceId              GraphObjectId `json:"sourceId"`
	DestinationId         GraphObjectId `json:"destinationId"`
	SourceOutputIndex     float64       `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64       `json:"destinationInputIndex,omitempty"`
}

/*
Notifies that AudioNodes are disconnected. The destination can be null, and it means all the outgoing connections from the source are disconnected.
*/
type NodesDisconnected struct {
	ContextId             GraphObjectId `json:"contextId"`
	SourceId              GraphObjectId `json:"sourceId"`
	DestinationId         GraphObjectId `json:"destinationId"`
	SourceOutputIndex     float64       `json:"sourceOutputIndex,omitempty"`
	DestinationInputIndex float64       `json:"destinationInputIndex,omitempty"`
}

/*
Notifies that an AudioNode is connected to an AudioParam.
*/
type NodeParamConnected struct {
	ContextId         GraphObjectId `json:"contextId"`
	SourceId          GraphObjectId `json:"sourceId"`
	DestinationId     GraphObjectId `json:"destinationId"`
	SourceOutputIndex float64       `json:"sourceOutputIndex,omitempty"`
}

/*
Notifies that an AudioNode is disconnected to an AudioParam.
*/
type NodeParamDisconnected struct {
	ContextId         GraphObjectId `json:"contextId"`
	SourceId          GraphObjectId `json:"sourceId"`
	DestinationId     GraphObjectId `json:"destinationId"`
	SourceOutputIndex float64       `json:"sourceOutputIndex,omitempty"`
}
