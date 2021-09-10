package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ecwid/control/transport"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
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

	var p = sess.Target()
	err = p.Navigate("https://my.ecwid.com", control.LifecycleIdleNetwork)
	if err != nil {
		panic(err)
	}
}
