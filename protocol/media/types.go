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
	Represents logged source line numbers reported in an error.

NOTE: file and line are from chromium c++ implementation code, not js.
*/
type PlayerErrorSourceLocation struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

/*
Corresponds to kMediaError
*/
type PlayerError struct {
	ErrorType string                       `json:"errorType"`
	Code      int                          `json:"code"`
	Stack     []*PlayerErrorSourceLocation `json:"stack"`
	Cause     []*PlayerError               `json:"cause"`
	Data      interface{}                  `json:"data"`
}
