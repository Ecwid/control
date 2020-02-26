package test

import (
	"testing"
	"time"

	"github.com/ecwid/witness/pkg/chrome"
)

func TestNew5TabOpen(t *testing.T) {
	t.Parallel()

	chrome, err := chrome.New("--disable-popup-blocking")
	if err != nil {
		t.Fatal(err)
	}
	defer chrome.Close()
	sess, err := chrome.CDP.DefaultSession()
	if err != nil {
		t.Fatal(err)
	}

	err = sess.Page.Navigate(getFilepath("new_tab.html"))
	if err != nil {
		t.Fatal(err)
	}
	chrome.CDP.Timeouts.Navigation = time.Millisecond * 1000
	c := sess.Tabs.OnNewTabOpen()
	sess.Page.C("#newtabs", true).Click()
	targetID := <-c
	if targetID == "" {
		t.Fatalf("no targetID returned")
	}
	s2, err := sess.Tabs.SwitchToTab(targetID)
	if err != nil {
		t.Fatal(err)
	}
	err = s2.Page.Activate()
	if err != nil {
		t.Fatal(err)
	}
	tabs, err := sess.Tabs.GetTabs()
	if err != nil {
		t.Fatal(err)
	}
	if len(tabs) != 5+1 {
		t.Fatalf("not 6 tabs but %d", len(tabs))
	}
}
