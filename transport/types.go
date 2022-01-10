package transport

import (
	"encoding/json"
	"fmt"
	"time"
)

var ErrShutdown = Error{Message: "connection is shut down"}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

type Event struct {
	Method string
	Params []byte
}

type Reply struct {
	ID        uint64          `json:"id,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method,omitempty"`
	Params    json.RawMessage `json:"params,omitempty"`
	Result    json.RawMessage `json:"result,omitempty"`
	Error     *Error          `json:"error,omitempty"`
}

type Call struct {
	ID        uint64      `json:"id"`
	SessionID string      `json:"sessionId,omitempty"`
	Method    string      `json:"method"`           // The name of the service and method to call.
	Args      interface{} `json:"params,omitempty"` // The argument to the function (*struct).
	Reply     chan Reply  `json:"-"`
}

type CallTimeoutError struct {
	Call    *Call
	Timeout time.Duration
}

func (r CallTimeoutError) Error() string {
	return fmt.Sprintf("the reply to the request %+v not received in %s", r.Call, r.Timeout)
}
