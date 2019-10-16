package main

import (
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func main() {
	chrome, _ := chrome.New("--headless")
	defer chrome.Close()
	page, err := chrome.CDP.DefaultPage()
	if err != nil {
		panic(err)
	}

	chrome.CDP.Logging.Level = witness.LevelProtocolMessage
	// Implicitly affected only C() function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	page.Navigate("https://my.ecwid.com")

	page.C("[name='email']", true).Type("test@example.com")
	page.C("[name='password']", true).Type("xxxxxx")
	page.C("button.btn-primary", true).Click()
}
