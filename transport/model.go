package transport

import (
	"encoding/json"
	"errors"
)

type Request struct {
	ID        uint64 `json:"id"`
	SessionID string `json:"sessionId,omitempty"`
	Method    string `json:"method"`
	Params    any    `json:"params,omitempty"`
}

func (r Request) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type Response struct {
	ID     uint64  `json:"id,omitempty"`
	Result Untyped `json:"result,omitempty"`
	Error  *Error  `json:"error,omitempty"`
	*Message
}

func (r Response) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type Message struct {
	SessionID string  `json:"sessionId,omitempty"`
	Method    string  `json:"method,omitempty"`
	Params    Untyped `json:"params,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

type Untyped []byte

func (m Untyped) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *Untyped) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("cdpnext.Untyped: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func Unmarshal[T any](m Message) (T, error) {
	var zero T
	if err := json.Unmarshal(m.Params, &zero); err != nil {
		return zero, err
	}
	return zero, nil
}
