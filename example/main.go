package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
	"github.com/ecwid/control/transport"
)

// Pretty преобразует struct в читаемый вид (форматированный json)
func Pretty(p interface{}) string {
	s, _ := json.MarshalIndent(p, "", "\t")
	return string(s)
}

func main() {
	b, err := chrome.Launch(context.TODO(), "--disable-popup-blocking") //, "--no-startup-window")
	if err != nil {
		panic(err)
	}
	defer b.Close()
	b.GetTransport().(*transport.WS).Stdout = os.Stdout
	sess := control.New(b.GetTransport())
	err = sess.CreateTarget("")
	if err != nil {
		panic(err)
	}

	var p = sess.Page()
	err = p.Navigate("https://surfparadise.ecwid.com/", control.LifecycleIdleNetwork)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
	app, err := p.QuerySelectorAll(".grid-product__title-inner")
	if err != nil {
		panic(err)
	}
	for _, i := range app {
		_, err = i.GetText()
		if err != nil {
			panic(err)
		}
	}
}
