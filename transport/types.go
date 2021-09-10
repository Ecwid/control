package transport

import (
	"encoding/json"
	"fmt"
)

var (
	ErrConnectionClosed = ProtoError{Message: "Connection closed"}
)

type ProtoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
	Request []byte `json:"-"`
}

func (e ProtoError) Error() string {
	return fmt.Sprintf("%s\n%s", e.Message, string(e.Request))
}

type Request struct {
	ID        uint64      `json:"id"`
	SessionID string      `json:"sessionId,omitempty"`
	Method    string      `json:"method"`
	Params    interface{} `json:"params,omitempty"`
}

type Response struct {
	ID        uint64          `json:"id,omitempty"`
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method,omitempty"`
	Params    json.RawMessage `json:"params,omitempty"`
	Result    json.RawMessage `json:"result,omitempty"`
	Error     *ProtoError     `json:"error,omitempty"`
}

func (r Response) isBroadcast() bool {
	return r.ID == 0 && r.Method != ""
}

func (r Response) isError() bool {
	return r.Error != nil && r.Error.Code != 0
}
