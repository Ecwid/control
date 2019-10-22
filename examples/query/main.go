package main

import (
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func main() {
	chrome, _ := chrome.New("--headless")
	defer chrome.Close()
	session, err := chrome.CDP.DefaultSession()
	if err != nil {
		panic(err)
	}

	chrome.CDP.Logging.Level = witness.LevelProtocolMessage
	// Implicitly affected only C() function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	session.Page.Navigate("https://my.ecwid.com")

	session.Page.C("[name='email']", true).Type("test@example.com")
	session.Page.C("[name='password']", true).Type("xxxxxx")
	session.Page.C("button.btn-primary", true).Click()
}
