package control

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/ecwid/control/cdp"
	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/protocol/target"
)

func Take(args ...string) (session *Session, cancel func() error, err error) {
	return TakeWithContext(context.TODO(), nil, args...)
}

func hasProfile(args ...string) bool {
	for _, a := range args {
		a = strings.TrimSpace(a)
		if strings.HasPrefix(a, "--profile-directory") || strings.HasPrefix(a, "--user-data-dir") {
			return true
		}
	}
	return false
}

func TakeWithContext(ctx context.Context, logger *slog.Logger, args ...string) (session *Session, cancel func() error, err error) {
	var userDataDir string

	if !hasProfile(args...) {
		// If there is no profile or user-data-dir in the startup arguments, then user-data-dir must be set
		userDataDir, err = os.MkdirTemp("", "chrome-control-*")
		if err != nil {
			return nil, nil, errors.Join(err, errors.New("can't create temporary user data dir"))
		}
		args = append(args, "--user-data-dir="+userDataDir)
	}

	browser, err := chrome.Launch(ctx, args...)
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

	cleanup := func() error {
		if err := transport.Close(); err != nil {
			return errors.Join(err, errors.New("can't close transport"))
		}
		if err = browser.Wait(); err != nil {
			return errors.Join(err, errors.New("can't close browser gracefully"))
		}
		if userDataDir != "" {
			if err := os.RemoveAll(userDataDir); err != nil {
				return errors.Join(err, errors.New("can't clear user data dir"))
			}
		}
		return nil
	}
	return session, cleanup, nil
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
