package io

import (
	"github.com/ecwid/control/protocol"
)

/*
	Close the stream, discard any temporary backing storage.
*/
func Close(c protocol.Caller, args CloseArgs) error {
	return c.Call("IO.close", args, nil)
}

/*
	Read a chunk of the stream
*/
func Read(c protocol.Caller, args ReadArgs) (*ReadVal, error) {
	var val = &ReadVal{}
	return val, c.Call("IO.read", args, val)
}

/*
	Return UUID of Blob object specified by a remote object id.
*/
func ResolveBlob(c protocol.Caller, args ResolveBlobArgs) (*ResolveBlobVal, error) {
	var val = &ResolveBlobVal{}
	return val, c.Call("IO.resolveBlob", args, val)
}
