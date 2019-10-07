package test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func htmlPath(name string) string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)
	return fmt.Sprintf("file://%s/testdata/%s", dir, name)
}

func TestClickHit(t *testing.T) {
	var expectedRate int64 = 60 // 60% click hit

	chrome, err := chrome.New()
	if err != nil {
		t.Fatal(err)
	}
	defer chrome.Close()
	page, err := chrome.CDP.DefaultPage()
	if err != nil {
		t.Fatal(err)
	}
	page.Navigate(htmlPath("click_playground.html"))

	target := page.Doc().Expect("#target", false)

	var pass int64
	var miss int64
	for i := 0; i < 50; i++ {
		err := target.Click()
		switch err {
		case nil:
			pass++
		case witness.ErrElementMissClick:
			miss++
		default:
			t.Fatal(err)
		}
		time.Sleep(time.Millisecond * 300)
	}

	clickedText, err := target.GetText()
	if err != nil {
		t.Fatal(err)
	}
	clicked, err := strconv.ParseInt(clickedText, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	rate := (100 * pass) / (miss + pass)
	t.Logf("pass = %d, miss = %d, rate = %d", pass, miss, rate)
	if rate <= expectedRate {
		t.Fatalf("miss click degradation - expected at least %d%% success click, but was %d", expectedRate, rate)
	}
	if clicked != pass {
		t.Fatalf("%d flaky clicks", pass-clicked)
	}

}
