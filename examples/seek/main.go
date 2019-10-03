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

	// Implicitly affected only Expect function
	chrome.CDP.Logging.Level = witness.LevelProtocolMessage
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	page.Navigate("https://my.ecwid.com")
	doc := page.Doc()

	doc.Expect("[name='email']", true).Type("test@example.com")
	doc.Expect("[name='password']", true).Type("xxxxxx")
	doc.Expect("button.btn-primary", true).Click()
}
