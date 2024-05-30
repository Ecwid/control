# ecwid-control
**control** is a automation tool written from scratch using golang on top of Chrome DevTools

_Warning_ This is an experimental project, backward compatibility is not guaranteed!

## Installation
`go get -u github.com/ecwid/control`

## How to use

Here is an example of using:

```go
func main() {
	session, cancel, err := control.Take("--no-startup-window")
	if err != nil {
		panic(err)
	}
	defer cancel()

	retrier := retry.Static{
		Timeout: 10 * time.Second,
		Delay:   500 * time.Millisecond, // delay between attempts
	}

	session.Frame.MustNavigate("https://zoid.ecwid.com")

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

    // "must way" throws panic on an error

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

Subscribe on the domain event
```go
future := control.Subscribe(session, "Target.targetCreated", func(t target.TargetCreated) bool {
    return t.TargetInfo.Type == "page"
})
defer future.Cancel()

// do something here ...

ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
result /* target.TargetCreated */, err := future.Get(ctx)
```