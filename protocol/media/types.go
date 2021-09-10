package media

/*
	Players will get an ID that is unique within the agent context.
*/
type PlayerId string

/*

 */
type Timestamp float64

/*
	Have one type per entry in MediaLogRecord::Type
Corresponds to kMessage
*/
type PlayerMessage struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

/*
	Corresponds to kMediaPropertyChange
*/
type PlayerProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
	Corresponds to kMediaEventTriggered
*/
type PlayerEvent struct {
	Timestamp Timestamp `json:"timestamp"`
	Value     string    `json:"value"`
}

/*
	Corresponds to kMediaError
*/
type PlayerError struct {
	Type      string `json:"type"`
	ErrorCode string `json:"errorCode"`
}
