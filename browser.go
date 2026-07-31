package control

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/protocol/browser"
	"github.com/ecwid/control/protocol/target"
	"github.com/ecwid/control/transport"
)

const defaultTimeout = 10 * time.Second

type Browser struct {
	caller CdpCaller
	chrome chrome.Chrome
	closer sync.Once
}

func (b *Browser) Context() context.Context {
	return b.caller.Context()
}

func (b *Browser) Close() error {
	var browserErr, chromeErr error
	b.closer.Do(func() {
		browserErr = browser.Close(b.caller)
		b.caller.transport.Close()
		chromeErr = b.chrome.Close(browserErr != nil)
	})
	return errors.Join(browserErr, chromeErr)
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

func (b *Browser) attachToTarget(targetID target.TargetID) (target.SessionID, error) {
	val, err := target.AttachToTarget(b.caller, target.AttachToTargetArgs{
		TargetId: targetID,
		Flatten:  true,
	})
	if err != nil {
		return "", err
	}
	return val.SessionId, nil
}

type Options struct {
	CdpTimeout time.Duration
	Logger     *slog.Logger
	ChromeArgs []string
}

func Launch(ctx context.Context, opts Options) (*Browser, error) {
	chromeBrowser, err := chrome.Launch(ctx, opts.ChromeArgs...)
	if err != nil {
		return nil, errors.Join(err, errors.New("chrome launch failed"))
	}
	cdp, err := transport.DefaultDial(ctx, chromeBrowser.WebSocketUrl, opts.Logger)
	if err != nil {
		_ = chromeBrowser.Close(true)
		return nil, errors.Join(err, errors.New("websocket connection failed"))
	}
	if opts.CdpTimeout <= 0 {
		opts.CdpTimeout = defaultTimeout
	}
	browserCtx, browserCancel := context.WithCancelCause(cdp.Context())
	caller := CdpCaller{
		ctx:       browserCtx,
		cancel:    browserCancel,
		transport: cdp,
		timeout:   opts.CdpTimeout,
	}
	return &Browser{caller: caller, chrome: chromeBrowser}, nil
}
