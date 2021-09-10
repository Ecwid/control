package main

import (
	"context"
	"encoding/json"
	"log"
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
	//trans := transport.MustConnect(context.TODO(), "ws://127.0.0.1:61958/devtools/browser/2ca11e9a-317d-4c9b-84dd-133598128f61")
	//trans.(*transport.Client).Out = os.Stdout
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
	err = p.MustElement("input[name='email']").InsertText("zoid+bot@ecwid.com")
	if err != nil {
		panic(err)
	}
	err = p.MustElement("input[name='password']").InsertText("123456")
	if err != nil {
		panic(err)
	}
	err = p.NewLifecycleEventCondition(control.LifecycleIdleNetwork).Do(func() error {
		return p.MustElement("[id='SIF.sIB']").Click()
	})
	if err != nil {
		panic(err)
	}
	text, err := p.MustElement("span.store-id span:nth-child(2)").GetText()
	if err != nil {
		panic(err)
	}
	log.Print(text)

}
