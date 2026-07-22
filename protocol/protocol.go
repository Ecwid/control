package protocol

type Caller interface {
	Call(method string, send, recv any) error
}
