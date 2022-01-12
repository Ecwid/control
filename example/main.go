package main

import (
	"context"
	"log"
	"time"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/transport"
)

func main() {

	chromium, err := chrome.Launch(context.TODO(), "--disable-popup-blocking") // you can specify more startup parameters for chrome
	if err != nil {
		panic(err)
	}
	defer chromium.Close()
	ctrl := control.New(chromium.GetClient())
	ctrl.Client.Timeout = time.Second * 60

	go func() {
		s1, err := ctrl.CreatePageTarget("")
		if err != nil {
			panic(err)
		}
		cancel := s1.Subscribe("Page.domContentEventFired", func(e transport.Event) {
			v, err1 := s1.Page().GetNavigationEntry()
			log.Println(v)
			log.Println(err1)
		})
		defer cancel()
		if err = s1.Page().Navigate("https://google.com/", control.LifecycleIdleNetwork, time.Second*60); err != nil {
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
