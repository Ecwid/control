package chrome

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

var MaxTimeToStart = 10 * time.Second

type Chrome struct {
	WebSocketUrl string
	UserDataDir  string
	StartArgs    string
	cmd          *exec.Cmd
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

func (c Chrome) NewTab(cli *http.Client, address string) (target Target, err error) {
	u, err := url.Parse(c.WebSocketUrl)
	if err != nil {
		return target, err
	}
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf(`http://`+u.Host+`/json/new?`+address), nil)
	if err != nil {
		return target, err
	}
	r, err := cli.Do(request)
	if err != nil {
		return target, err
	}
	var b []byte
	b, err = io.ReadAll(r.Body)
	if err != nil {
		return target, err
	}
	if err = r.Body.Close(); err != nil {
		return
	}
	if err = json.Unmarshal(b, &target); err != nil {
		return
	}
	return
}

func (c Chrome) WaitCloseGracefully() error {
	defer func() {
		err := os.RemoveAll(c.UserDataDir)
		if err != nil {
			log.Println(err)
		}
	}()
	return c.cmd.Wait()
}

func bin() string {
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
			return path
		}
	}
	panic("chrome binary not found")
}

func Launch(ctx context.Context, userFlags ...string) (value Chrome, err error) {
	if value.UserDataDir, err = os.MkdirTemp("", "chrome-control-*"); err != nil {
		return value, err
	}
	// https://github.com/GoogleChrome/chrome-launcher/blob/master/docs/chrome-flags-for-tools.md
	// https://docs.google.com/spreadsheets/d/1n-vw_PCPS45jX3Jt9jQaAhFqBY6Ge1vWF_Pa0k7dCk4/edit#gid=1265672696
	flags := []string{
		"--remote-debugging-port=0",
		"--user-data-dir=" + value.UserDataDir,
	}
	if len(userFlags) > 0 {
		flags = append(flags, userFlags...)
	}
	if os.Getuid() == 0 {
		flags = append(flags, "--no-sandbox", "--disable-setuid-sandbox")
	}
	binary := bin()
	value.StartArgs = fmt.Sprint(binary, strings.Join(flags, " "))
	value.cmd = exec.CommandContext(ctx, binary, flags...)

	stderr, err := value.cmd.StderrPipe()
	if err != nil {
		return value, err
	}

	addr := make(chan string)
	var std []string
	go func() {
		const prefix = "DevTools listening on"
		var scanner = bufio.NewScanner(stderr)
		for scanner.Scan() {
			line := scanner.Text()
			std = append(std, line)
			if s := strings.TrimPrefix(line, prefix); s != line {
				addr <- strings.TrimSpace(s)
				return
			}
		}
	}()

	if err = value.cmd.Start(); err != nil {
		return value, err
	}

	select {
	case value.WebSocketUrl = <-addr:
		return value, nil
	case <-time.After(MaxTimeToStart):
		return value, fmt.Errorf("chrome stopped too early %s", strings.Join(std, "\n"))
	}
}
