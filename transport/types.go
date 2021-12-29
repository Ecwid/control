package transport

import (
	"encoding/json"
	"fmt"
	"time"
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

type ReceiveTimeoutError struct {
	Value   []byte
	Timeout time.Duration
}

func (r ReceiveTimeoutError) Error() string {
	return fmt.Sprintf("the reply to the request %s not received in %s", string(r.Value), r.Timeout)
}

type ansiColor = string

const (
	black   ansiColor = "\u001b[30;1m"
	red     ansiColor = "\u001b[31;1m"
	green   ansiColor = "\u001b[32;1m"
	yellow  ansiColor = "\u001b[33;1m"
	blue    ansiColor = "\u001b[34;1m"
	magenta ansiColor = "\u001b[35;1m"
	cyan    ansiColor = "\u001b[36;1m"
	white   ansiColor = "\u001b[37;1m"
	reset   ansiColor = "\u001b[0m"
)
