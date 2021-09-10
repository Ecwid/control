package io

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
	This is either obtained from another method or specifed as `blob:&lt;uuid&gt;` where
`&lt;uuid&gt` is an UUID of a Blob.
*/
type StreamHandle string

type CloseArgs struct {
	Handle StreamHandle `json:"handle"`
}

type ReadArgs struct {
	Handle StreamHandle `json:"handle"`
	Offset int          `json:"offset,omitempty"`
	Size   int          `json:"size,omitempty"`
}

type ReadVal struct {
	Base64Encoded bool   `json:"base64Encoded,omitempty"`
	Data          string `json:"data"`
	Eof           bool   `json:"eof"`
}

type ResolveBlobArgs struct {
	ObjectId runtime.RemoteObjectId `json:"objectId"`
}

type ResolveBlobVal struct {
	Uuid string `json:"uuid"`
}
