# ecwid-control
**control** is a test automation tool written from scratch using golang on top of Chrome DevTools

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
	"os"

	"github.com/ecwid/control"
	"github.com/ecwid/control/chrome"
)

func main() {
	browser, err := chrome.Launch(context.TODO(), "--disable-popup-blocking") // you can specify more startup parameters for chrome
	if err != nil {
		panic(err)
	}
	defer browser.Close()
	browser.GetTransport().Stderr = os.Stderr // enabled by default
	// browser.GetTransport().Stdout = os.Stdout // uncomment to get CDP logs
	session := control.New(browser.GetTransport())
	err = session.CreateTarget("") // create a new browser tab with a blank page
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
```
