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
		"log/slog"
		"time"

		"github.com/ecwid/control"
		"github.com/ecwid/control/retry"
	)

	func main() {
		logger := slog.Default()
		browser, err := control.Launch(context.Background(), logger, "--no-startup-window")
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := browser.Close(); err != nil {
				log.Println("browser close error:", err)
			}
		}()

		tab, err := browser.NewTab()
		if err != nil {
			panic(err)
		}

		session, err := browser.NewSession(tab)
		if err != nil {
			panic(err)
		}

		session.Frame.MustNavigate("https://zoid.ecwid.com")

		retrier := retry.Static{
			Timeout: 10 * time.Second,
			Delay:   500 * time.Millisecond,
		}

		var products []string
		err = retry.Func(retrier, func() error {
			products = []string{}
			return session.Frame.QueryAll(".grid-product__title-inner").Then(func(nl control.NodeList) error {
				return nl.Foreach(func(n *control.Node) error {
					return n.GetText().Then(func(s string) error {
						products = append(products, s)
						return nil
					})
				})
			})
		})
		if err != nil {
			panic(err)
		}

		for _, node := range session.Frame.MustQueryAll(".grid-product__title-inner") {
			log.Println(node.MustGetText())
		}
	}
```
You can call any CDP method implemented in protocol package using a session
```go
	err = security.SetIgnoreCertificateErrors(session, security.SetIgnoreCertificateErrorsArgs{
		Ignore: true,
	})
```
or call a custom unimplemented method
```go
	err = session.Call("Security.setIgnoreCertificateErrors", sendStruct, receiveStruct)
```