package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ecwid/witness/pkg/chrome"
	"github.com/ecwid/witness/pkg/mobile"
)

func main() {
	chrome, _ := chrome.New()
	defer chrome.Close()
	session, err := chrome.CDP.DefaultSession()
	if err != nil {
		panic(err)
	}

	mobile := []*mobile.Device{
		mobile.GalaxyS5,
		mobile.IPad,
		mobile.IPhone8.Rotated(),
		mobile.IPadPro.Rotated(),
	}

	session.Page.Navigate("https://mdemo.ecwid.com/")

	for i, m := range mobile {
		if err := session.Emulation.Emulate(m); err != nil {
			panic(err)
		}
		b, _ := session.Page.CaptureScreenshot("png", 100, true, nil)
		ioutil.WriteFile(fmt.Sprintf("%d.png", i), b, 0644)
	}

}
