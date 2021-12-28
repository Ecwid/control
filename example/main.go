package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
)

// Pretty преобразует struct в читаемый вид (форматированный json)
func Pretty(p interface{}) string {
	s, _ := json.MarshalIndent(p, "", "\t")
	return string(s)
}

func main() {
	browser, err := chrome.Launch(context.TODO(), "--disable-popup-blocking") // you can specify more startup parameters for chrome
	if err != nil {
		panic(err)
	}
	defer browser.Close()
	browser.GetClient().Stderr = os.Stderr // enabled by default
	// browser.GetClient().Stdout = os.Stdout // uncomment to get CDP logs

	cdp := control.New(browser.GetClient())
	go func() {
		s1, err := cdp.CreatePageTarget("")
		if err != nil {
			panic(err)
		}
		if err := s1.Page().Navigate("https://google.com/", control.LifecycleIdleNetwork); err != nil {
			panic(err)
		}
	}()

	session, err := cdp.CreatePageTarget("")
	if err != nil {
		panic(err)
	}

	var page = session.Page() // main frame
	err = page.Navigate("https://surfparadise.ecwid.com/", control.LifecycleIdleNetwork)
	if err != nil {
		panic(err)
	}

	items, err := page.QuerySelectorAll(".grid-product__title-inner")
	if err != nil {
		panic(err)
	}
	for _, i := range items {
		title, err := i.GetText()
		if err != nil {
			panic(err)
		}
		log.Print(title)
	}
}
