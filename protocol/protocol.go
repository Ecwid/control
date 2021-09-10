package protocol

type Caller interface {
	Call(method string, send, recv interface{}) error
}
