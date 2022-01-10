# ecwid-control
**control** is a automation tool written from scratch using golang on top of Chrome DevTools

_Warning_ This is an experimental project, backward compatibility is not guaranteed!

## Installation
`go get -u github.com/ecwid/control`

## How to use

Here is an example of using:

```go
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
	
	session, err := ctrl.CreatePageTarget("")
	if err != nil {
		panic(err)
	}

	var page = session.Page() // main frame 
	err = page.Navigate("https://surfparadise.ecwid.com/", control.LifecycleIdleNetwork, time.Second*60)
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
```

You can call any CDP method implemented in protocol package using a session
```go
err = security.SetIgnoreCertificateErrors(sess, security.SetIgnoreCertificateErrorsArgs{
    Ignore: true,
})
```

or even call a custom method
```go
err = sess.Call("Security.setIgnoreCertificateErrors", sendStruct, receiveStruct)
```

Subscribe on domain event
```go
cancel := sess.Subscribe("Overlay.screenshotRequested", func(e observe.Value) {
    v := overlay.ScreenshotRequested{}
    _= json.Unmarshal(e.Params, &v)
    doSomething(v.Viewport.Height)
})
defer cancel()

// Subscribe on all incoming events
sess.Subscribe("*", func(e observe.Value) {
    switch e.Method {
        case "Overlay.screenshotRequested":
    }
})

```
