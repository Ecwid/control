package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

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

type Response struct {
	ID        uint64          `json:"id,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method,omitempty"`
	Params    json.RawMessage `json:"params,omitempty"`
	Result    json.RawMessage `json:"result,omitempty"`
	Error     *Error          `json:"error,omitempty"`
}

type Request struct {
	ID        uint64      `json:"id"`
	SessionID string      `json:"sessionId,omitempty"`
	Method    string      `json:"method"`           // The name of the service and method to call.
	Args      interface{} `json:"params,omitempty"` // The argument to the function (*struct).
	response  chan Response
}

func (request *Request) received(r Response) error {
	select {
	case request.response <- r:
		return nil
	default:
		return errors.New("the response received twice")
	}
}

type DeadlineExceededError struct {
	Request *Request
	Timeout time.Duration
}

func (r DeadlineExceededError) Error() string {
	return fmt.Sprintf("the reply to the request [sessionID: %s, Method: %s, Args: %v] not received in %s",
		r.Request.SessionID, r.Request.Method, r.Request.Args, r.Timeout)
}
