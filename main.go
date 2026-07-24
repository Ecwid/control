package control

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

type Browser struct {
	ctx       context.Context
	transport *transport.Transport
	chrome    chrome.Chrome
}

func (b Browser) NewSession(t chrome.Target) (*Session, error) {
	return NewSession(b.transport, target.TargetID(t.ID))
}

func (b Browser) Close() error {
	return b.chrome.Close(b.ctx, b.transport)
}

func (b Browser) NewTab() (chrome.Target, error) {
	return b.chrome.NewTab(http.DefaultClient, "")
}

func Launch(ctx context.Context, logger *slog.Logger, args ...string) (Browser, error) {
	browser, err := chrome.Launch(ctx, args...)
	if err != nil {
		return Browser{}, errors.Join(err, errors.New("chrome launch failed"))
	}
	cdp, err := transport.DefaultDial(ctx, browser.WebSocketUrl, logger)
	if err != nil {
		browser.Close(ctx, nil)
		return Browser{}, errors.Join(err, errors.New("websocket dial failed"))
	}
	return Browser{ctx: ctx, transport: cdp, chrome: browser}, nil
}
