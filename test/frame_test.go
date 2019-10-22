package test

import (
	"testing"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func TestFrameRefresh(t *testing.T) {
	t.Parallel()

	chrome, err := chrome.New()
	check(t, err)
	defer chrome.Close()
	sess, err := chrome.CDP.DefaultSession()
	check(t, err)

	get := func(sel string) witness.Element {
		t.Helper()
		el, err := sess.Page.Query(sel)
		check(t, err)
		return el
	}

	check(t, sess.Page.Navigate(getFilepath("frame_playground.html")))
	fid, err := get("#my_frame").GetFrameID()
	check(t, err)
	check(t, sess.Page.SwitchToFrame(fid))
	finp := get("#frameInput1")
	check(t, finp.Type("123456"))
	check(t, get("#refresh").Click())
	time.Sleep(time.Second * 2)
	check(t, finp.Type("654321"))
}

func TestFrameRenew(t *testing.T) {
	t.Parallel()

	chrome, err := chrome.New()
	check(t, err)
	defer chrome.Close()
	// chrome.CDP.Logging.Level = witness.LevelProtocolVerbose
	sess, err := chrome.CDP.DefaultSession()
	check(t, err)

	get := func(sel string) witness.Element {
		t.Helper()
		el, err := sess.Page.Query(sel)
		check(t, err)
		return el
	}

	url := getFilepath("frame_playground.html")

	check(t, sess.Page.Navigate(url))
	fid, err := get("#my_frame").GetFrameID()
	check(t, err)
	check(t, get("#button1").Click())

	check(t, sess.Page.SwitchToFrame(fid))
	time.Sleep(time.Second * 4)

	if _, err := sess.Page.Query("#frameButton1"); err != witness.ErrFrameDetached {
		t.Fatalf("not expected error: %s", err.Error())
	}
}
