package chrome

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/log"
)

// Chrome браузер
type Chrome struct {
	CDP    *witness.CDP
	cancel func()
}

// Close закрывает хром
func (c *Chrome) Close() {
	c.CDP.Close()
	c.cancel()
}

// Connect ...
func Connect(webSocketURL string) (*Chrome, error) {
	var err error
	chrome := &Chrome{}
	chrome.cancel = func() {}
	chrome.CDP, err = witness.New(context.Background(), webSocketURL)
	if err != nil {
		return nil, err
	}
	return chrome, nil
}

// New запускает новый хром
func New(userFlags ...string) (*Chrome, error) {

	chrome := &Chrome{}
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

	flags := []string{
		"--no-first-run",
		"--no-default-browser-check",
		"--remote-debugging-port=0",
		"--hide-scrollbars",
		"--password-store=basic",
		"--use-mock-keychain",
		"--enable-automation",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--disable-default-apps",
		"--disable-extensions",
		"--disable-browser-side-navigation",
		"--disable-features=site-per-process,TranslateUI,BlinkGenPropertyTrees",
		"--disable-background-networking",
		"--disable-backgrounding-occluded-windows",
		"--disable-renderer-backgrounding",
		"--disable-hang-monitor",
		"--enable-features=NetworkService,NetworkServiceInProcess",
		"--user-data-dir=" + userDataDir,
	}

	for _, f := range userFlags {
		flags = append(flags, f)
	}

	if os.Getuid() == 0 {
		flags = append(flags, "--no-sandbox")
	}

	cmd := exec.CommandContext(context.Background(), path, flags...)
	chrome.cancel = func() {
		exited := make(chan int)
		go func() {
			state, _ := cmd.Process.Wait()
			exited <- state.ExitCode()
		}()
		select {
		case <-exited:
			return
		case <-time.After(time.Millisecond * 2000):
			if err := cmd.Process.Kill(); err != nil {
				log.Print(log.LevelFatal, err)
			}
			return
		}
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	webSocketURL, err := addrFromStderr(stderr)
	if err != nil {
		return nil, err
	}
	chrome.CDP, err = witness.New(context.Background(), webSocketURL)
	if err != nil {
		return nil, err
	}
	return chrome, nil
}

func addrFromStderr(rc io.ReadCloser) (string, error) {
	defer rc.Close()
	url := ""
	scanner := bufio.NewScanner(rc)
	prefix := "DevTools listening on"

	var lines []string
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
