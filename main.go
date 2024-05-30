package control

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ecwid/control/cdp"
	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/protocol/target"
)

func Take(args ...string) (session *Session, cancel func(), err error) {
	return TakeWithContext(context.TODO(), nil, args...)
}

func TakeWithContext(ctx context.Context, logger *slog.Logger, chromeArgs ...string) (session *Session, cancel func(), err error) {
	browser, err := chrome.Launch(ctx, chromeArgs...)
	if err != nil {
		return nil, nil, errors.Join(err, errors.New("chrome launch failed"))
	}
	tab, err := browser.NewTab(http.DefaultClient, "")
	if err != nil {
		return nil, nil, errors.Join(err, errors.New("failed to open a new tab"))
	}
	transport, err := cdp.DefaultDial(ctx, browser.WebSocketUrl, logger)
	if err != nil {
		return nil, nil, errors.Join(err, errors.New("websocket dial failed"))
	}
	session, err = NewSession(transport, target.TargetID(tab.ID))
	if err != nil {
		return nil, nil, errors.Join(err, errors.New("failed to create a new session"))
	}
	teardown := func() {
		if err := transport.Close(); err != nil {
			transport.Log(slog.LevelError, "can't close transport", "err", err.Error())
		}
		if err = browser.WaitCloseGracefully(); err != nil {
			transport.Log(slog.LevelError, "can't close browser gracefully", "err", err.Error())
		}
	}
	return session, teardown, nil
}

func Subscribe[T any](s *Session, method string, filter func(T) bool) cdp.Future[T] {
	var (
		channel, cancel = s.Subscribe()
	)
	callback := func(resolve func(T), reject func(error)) {
		for value := range channel {
			if value.Method == method {
				var result T
				if err := json.Unmarshal(value.Params, &result); err != nil {
					reject(err)
					return
				}
				if filter(result) {
					resolve(result)
					return
				}
			}
		}
	}
	return cdp.NewPromise(callback, cancel)
}
