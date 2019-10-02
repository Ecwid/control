package main

import (
	"time"

	"github.com/ecwid/witness/pkg/chrome"
	"github.com/ecwid/witness/pkg/log"
)

func main() {
	log.Logging = log.LevelFatal | log.LevelProtocolMessage

	chrome, _ := chrome.New()
	defer chrome.Close()
	page, err := chrome.CDP.DefaultPage()
	if err != nil {
		panic(err)
	}

	// Implicitly affected only Expect function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	page.Navigate("https://my.ecwid.com")
	doc := page.Doc()

	doc.Expect("[name='email']", true).Type("test@example.com")
	doc.Expect("[name='password']", true).Type("xxxxxx")
	doc.Expect("button.btn-primary", true).Click()
}
