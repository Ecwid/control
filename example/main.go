package main

import (
	"context"
	"log"
	"time"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
)

func main() {

	chromium, err := chrome.Launch(context.TODO(), "--disable-popup-blocking") // you can specify more startup parameters for chrome
	if err != nil {
		panic(err)
	}
	defer chromium.Close()
	ctrl := control.New(chromium.GetClient())
	//ctrl.Client.Logger = os.Stdout
	ctrl.Client.Timeout = time.Second * 60

	go func() {
		s1, err := ctrl.CreatePageTarget("")
		if err != nil {
			panic(err)
		}
		if err := s1.Page().Navigate("https://google.com/", control.LifecycleIdleNetwork, time.Second*60); err != nil {
			panic(err)
		}
	}()

	session, err := ctrl.CreatePageTarget("")
	if err != nil {
		panic(err)
	}

	var page = session.Page() // main frame
	err = page.Navigate("https://surfparadise.ecwid.com/", control.LifecycleIdleNetwork, time.Second*60)
	if err != nil {
		panic(err)
	}

	_ = session.Activate()

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
