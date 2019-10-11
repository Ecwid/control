package test

import (
	"testing"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func TestInFrameRefresh(t *testing.T) {
	t.Parallel()

	chrome, err := chrome.New()
	check(t, err)
	defer chrome.Close()
	page, err := chrome.CDP.DefaultPage()
	check(t, err)

	get := func(sel string) witness.Element {
		t.Helper()
		el, err := page.Doc().Seek(sel)
		check(t, err)
		return el
	}

	check(t, page.Navigate(getFilepath("frame_playground.html")))
	fid, err := get("#my_frame").GetFrameID()
	check(t, err)
	check(t, page.SwitchToFrame(fid))
	check(t, get("#frameInput1").Type("123456"))
	check(t, get("#refresh").Click())
	time.Sleep(time.Second * 2)
	check(t, get("#frameInput1").Type("654321"))
}
