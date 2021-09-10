package chrome

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ecwid/control/transport"
)

// BrowserTarget ...
type BrowserTarget struct {
	Description          string `json:"description"`
	DevtoolsFrontendURL  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// BrowserVersion ...
type BrowserVersion struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// Browser ...
type Browser struct {
	url      *url.URL
	cmd      *exec.Cmd
	conn     transport.T
	deadline time.Duration
}

// GetTransport ...
func (c Browser) GetTransport() transport.T {
	return c.conn
}

// Crash ...
func (c Browser) Crash() {
	_ = c.conn.Call("", "Browser.crash", nil, nil)
}

func (c Browser) request(path string, response interface{}) error {
	r, err := http.Get("http://" + c.url.Host + path)
	if err != nil {
		return err
	}
	return json.NewDecoder(r.Body).Decode(response)
}

// GetVersion ...
func (c Browser) GetVersion() (BrowserVersion, error) {
	var result = BrowserVersion{}
	err := c.request("/json/version", &result)
	return result, err
}

// GetTargets ...
func (c Browser) GetTargets() ([]BrowserTarget, error) {
	var result []BrowserTarget
	err := c.request("/json", &result)
	return result, err
}

// Close close browser
func (c Browser) Close() error {
	// Close close browser and websocket connection
	exited := make(chan int)
	go func() {
		state, _ := c.cmd.Process.Wait()
		exited <- state.ExitCode()
	}()
	_ = c.conn.Call("", "Browser.close", nil, nil)
	select {
	case <-exited:
		return nil
	case <-time.After(c.deadline):
		if err := c.cmd.Process.Kill(); err != nil {
			return err
		}
		return errors.New("browser is not closing gracefully, process was killed")
	}
}

// Launch launch a new browser process
func Launch(ctx context.Context, userFlags ...string) (*Browser, error) {
	browser := &Browser{deadline: 10 * time.Second}
	var path string
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
		if _, err := exec.LookPath(c); err == nil {
			path = c
			break
		}
	}

	userDataDir, err := ioutil.TempDir("", "tmp")
	if err != nil {
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
		"--disable-browser-side-navigation",
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
		"--disable-features=site-per-process,TranslateUI,BlinkGenPropertyTrees",
		"--enable-features=NetworkService,NetworkServiceInProcess",
		"--user-data-dir=" + userDataDir,
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
	defer stderr.Close()
	if err := browser.cmd.Start(); err != nil {
		return nil, err
	}
	webSocketURL, err := addrFromStderr(stderr)
	if err != nil {
		return nil, err
	}
	browser.url, err = url.Parse(webSocketURL)
	if err != nil {
		return nil, err
	}
	browser.conn, err = transport.Connect(ctx, webSocketURL)
	return browser, err
}

func addrFromStderr(rc io.ReadCloser) (string, error) {
	defer rc.Close()
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
