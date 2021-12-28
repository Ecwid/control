package chrome

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ecwid/control/transport"
)

// Browser ...
type Browser struct {
	context      context.Context
	webSocketURL string
	cmd          *exec.Cmd
	client       *transport.Client
	UserDataDir  string
}

func (c Browser) GetClient() *transport.Client {
	return c.client
}

// Close close browser
func (c Browser) Close() error {
	// Close close browser and websocket connection
	exited := make(chan int, 1)
	go func() {
		state, _ := c.cmd.Process.Wait()
		exited <- state.ExitCode()
	}()
	_ = c.client.Call(c.context, "", "Browser.close", nil, nil)
	select {
	case <-exited:
		return nil
	case <-time.After(time.Second * 10):
		if err := c.cmd.Process.Kill(); err != nil {
			return err
		}
		return errors.New("browser is not closing gracefully, process was killed")
	}
}

// Launch launch a new browser process
func Launch(ctx context.Context, userFlags ...string) (*Browser, error) {
	browser := &Browser{context: ctx}
	var (
		path string
		err  error
	)
	bin := []string{
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/usr/bin/google-chrome",
		"headless-shell",
		"chromium",
		"chromium-browser",
		"google-chrome",
		"google-chrome-stable",
		"google-chrome-beta",
		"google-chrome-unstable",
	}
	for _, c := range bin {
		if _, err = exec.LookPath(c); err == nil {
			path = c
			break
		}
	}

	if browser.UserDataDir, err = os.MkdirTemp("", "chrome-control"); err != nil {
		return nil, err
	}

	// https: //github.com/GoogleChrome/chrome-launcher/blob/master/docs/chrome-flags-for-tools.md
	flags := []string{
		"about:blank", // open url
		"--no-first-run",
		"--no-default-browser-check",
		"--remote-debugging-port=0",
		"--hide-scrollbars",
		"--mute-audio",
		"--password-store=basic",
		"--use-mock-keychain",
		"--enable-automation",
		"--disable-gpu",
		"--disable-sync",
		"--disable-background-networking",
		"--disable-default-apps",
		"--disable-extensions",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-renderer-backgrounding",
		"--disable-hang-monitor",
		"--disable-breakpad",
		"--disable-client-side-phishing-detection",
		"--disable-component-extensions-with-background-pages",
		"--disable-ipc-flooding-protection",
		"--disable-prompt-on-repost",
		"--metrics-recording-only",
		"--disable-features=site-per-process,Translate,BlinkGenPropertyTrees",
		"--enable-features=NetworkService,NetworkServiceInProcess",
		"--user-data-dir=" + browser.UserDataDir,
	}

	flags = append(flags, userFlags...)
	if os.Getuid() == 0 {
		flags = append(flags, "--no-sandbox")
	}

	browser.cmd = exec.CommandContext(ctx, path, flags...)
	stderr, err := browser.cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	if err = browser.cmd.Start(); err != nil {
		return nil, err
	}
	browser.webSocketURL, err = addrFromStderr(stderr)
	if err != nil {
		return nil, err
	}
	browser.client, err = transport.Connect(ctx, browser.webSocketURL)
	return browser, err
}

func addrFromStderr(rc io.ReadCloser) (string, error) {
	const prefix = "DevTools listening on"
	var (
		url     = ""
		scanner = bufio.NewScanner(rc)
		lines   []string
	)
	for scanner.Scan() {
		line := scanner.Text()
		if s := strings.TrimPrefix(line, prefix); s != line {
			url = strings.TrimSpace(s)
			break
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if url == "" {
		return "", fmt.Errorf("chrome stopped too early; stderr:\n%s", strings.Join(lines, "\n"))
	}
	return url, nil
}
