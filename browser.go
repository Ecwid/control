package control

import (
	"github.com/ecwid/control/transport"
)

type BrowserContext struct {
	Client *transport.Client
}

func (b BrowserContext) Call(method string, send, recv interface{}) error {
	return b.Client.Call(b.Client.Ctx, "", method, send, recv)
}
