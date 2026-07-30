package chrome

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	startTimeout = 10 * time.Second
)

type Chrome struct {
	ctx          context.Context
	WebSocketUrl string
	StartArgs    string
	cmd          *exec.Cmd
	userDataDir  string
}

type Target struct {
	Description          string `json:"description,omitempty"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl,omitempty"`
	ID                   string `json:"id,omitempty"`
	Title                string `json:"title,omitempty"`
	Type                 string `json:"type,omitempty"`
	Url                  string `json:"url,omitempty"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl,omitempty"`
}

type shutdowner interface {
	Shutdown(context.Context) error
}

func hasUserDataDir(args []string) bool {
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if strings.HasPrefix(arg, "--profile-directory=") || strings.HasPrefix(arg, "--user-data-dir=") {
			return true
		}
	}
	return false
}

func (c Chrome) Wait() error {
	if c.cmd == nil {
		return errors.New("chrome process not started")
	}
	return c.cmd.Wait()
}

func (c Chrome) Close(kill bool) error {
	var errs []error

	if kill {
		errs = append(errs, c.cmd.Process.Kill())
	}

	if err := c.Wait(); err != nil {
		errs = append(errs, errors.Join(err, errors.New("cannot close browser gracefully")))
	}

	if c.userDataDir != "" {
		if err := os.RemoveAll(c.userDataDir); err != nil {
			errs = append(errs, errors.Join(err, errors.New("cannot clear user data directory")))
		}
	}

	return errors.Join(errs...)
}

func bin() (string, error) {
	for _, path := range []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/usr/bin/google-chrome",
		"headless-shell",
		"browser",
		"chromium",
		"chromium-browser",
		"google-chrome",
		"google-chrome-stable",
		"google-chrome-beta",
		"google-chrome-unstable",
	} {
		if _, err := exec.LookPath(path); err == nil {
			return path, nil
		}
	}
	return "", errors.New("chrome binary not found")
}

func Launch(ctx context.Context, userFlags ...string) (value Chrome, err error) {
	// https://github.com/GoogleChrome/chrome-launcher/blob/master/docs/chrome-flags-for-tools.md
	// https://docs.google.com/spreadsheets/d/1n-vw_PCPS45jX3Jt9jQaAhFqBY6Ge1vWF_Pa0k7dCk4/edit#gid=1265672696
	var flags = []string{"--remote-debugging-port=0"}
	if os.Getuid() == 0 {
		flags = append(flags, "--no-sandbox", "--disable-setuid-sandbox")
	}
	if len(userFlags) > 0 {
		flags = append(flags, userFlags...)
	}

	if !hasUserDataDir(flags) {
		value.userDataDir, err = os.MkdirTemp("", "chrome-control-*")
		if err != nil {
			return value, errors.Join(err, errors.New("cannot create temporary user data directory"))
		}
		flags = append(flags, "--user-data-dir="+value.userDataDir)
		defer func() {
			if err != nil {
				_ = os.RemoveAll(value.userDataDir)
				value.userDataDir = ""
			}
		}()
	}

	binary, err := bin()
	if err != nil {
		return value, err
	}

	value.ctx = ctx
	value.StartArgs = fmt.Sprintf("%s %s", binary, strings.Join(flags, " "))
	value.cmd = exec.CommandContext(ctx, binary, flags...)

	stderr, err := value.cmd.StderrPipe()
	if err != nil {
		return value, err
	}

	addr := make(chan string, 1)
	readDone := make(chan error, 1)

	go func() {
		const prefix = "DevTools listening on"
		var scanner = bufio.NewScanner(stderr)
		addrReported := false
		for scanner.Scan() {
			if addrReported {
				continue
			}
			line := scanner.Text()
			if s, ok := strings.CutPrefix(line, prefix); ok {
				select {
				case addr <- strings.TrimSpace(s):
					addrReported = true
				default:
				}
			}
		}
		readDone <- scanner.Err()
	}()

	if err = value.cmd.Start(); err != nil {
		return value, err
	}

	launchCtx, cancel := context.WithTimeout(ctx, startTimeout)
	defer cancel()

	select {

	case value.WebSocketUrl = <-addr:
		return value, nil

	case scanErr := <-readDone:
		waitErr := value.cmd.Wait()
		if scanErr != nil {
			return value, fmt.Errorf("chrome stderr read failed: %w", scanErr)
		}
		return value, fmt.Errorf("chrome stopped before reporting DevTools endpoint: %w", waitErr)

	case <-launchCtx.Done():
		if value.cmd.Process != nil {
			_ = value.cmd.Process.Kill()
		}
		_ = stderr.Close()
		_ = value.cmd.Wait()
		return value, fmt.Errorf("chrome launch timeout after %s: %w", startTimeout, context.Cause(launchCtx))
	}
}
