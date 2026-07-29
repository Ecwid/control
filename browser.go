package control

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

const defaultTimeout = 10 * time.Second

type Browser struct {
	caller CdpCaller
	chrome chrome.Chrome
}

func (b Browser) Close() error {
	return b.chrome.Close(b.caller.transport)
}

func (b *Browser) NewTab(url string) (*Session, error) {
	if url == "" {
		url = Blank // headless chrome crash when url is empty
	}
	created, err := target.CreateTarget(b.caller, target.CreateTargetArgs{Url: url})
	if err != nil {
		return nil, err
	}
	return b.NewSession(created.TargetId)
}

func (b Browser) attachToTarget(targetID target.TargetID) (target.SessionID, error) {
	val, err := target.AttachToTarget(b.caller, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return "", err
	}
	return val.SessionId, nil
}

func Launch(ctx context.Context, logger *slog.Logger, args ...string) (Browser, error) {
	chromeBrowser, err := chrome.Launch(ctx, args...)
	if err != nil {
		return Browser{}, errors.Join(err, errors.New("chrome launch failed"))
	}
	cdp, err := transport.DefaultDial(ctx, chromeBrowser.WebSocketUrl, logger)
	if err != nil {
		_ = chromeBrowser.Close(nil)
		return Browser{}, errors.Join(err, errors.New("websocket connection failed"))
	}
	browserCtx, browserCancel := context.WithCancelCause(cdp.Context())
	caller := CdpCaller{
		ctx:       browserCtx,
		cancel:    browserCancel,
		transport: cdp,
		timeout:   defaultTimeout,
	}
	return Browser{caller: caller, chrome: chromeBrowser}, nil
}
