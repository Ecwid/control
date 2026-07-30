package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/ecwid/control"
	"github.com/ecwid/control/lazy"
	"github.com/ecwid/control/retry"
)

type Handler struct {
	h slog.Handler
}

func (Handler) Enabled(c context.Context, l slog.Level) bool {
	return l >= slog.LevelInfo
}

func (h Handler) Handle(c context.Context, r slog.Record) error {
	buf := bytes.Buffer{}
	buf.WriteString(r.Time.Format(time.TimeOnly))
	buf.WriteByte(' ')
	buf.WriteString(r.Level.String())
	buf.WriteByte(' ')
	buf.WriteString(r.Message)
	buf.WriteByte(' ')
	body := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		body[a.Key] = a.Value.Any()
		return true
	})
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", " ")
	err := enc.Encode(body)
	if err != nil {
		return err
	}
	fmt.Print(buf.String())
	return nil
}

func (h Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.h.WithAttrs(attrs)
}

func (h Handler) WithGroup(name string) slog.Handler {
	return h.h.WithGroup(name)
}

func main() {
	sl := slog.New(Handler{h: slog.Default().Handler()})

	browser, err := control.Launch(context.TODO(), control.Options{
		CdpTimeout: 10 * time.Second,
		Logger:     sl,
		ChromeArgs: []string{"--no-startup-window"},
	})
	if err != nil {
		panic(err)
	}
	defer browser.Close()
	session, err := browser.NewTab("")
	if err != nil {
		panic(err)
	}

	r := retry.Static{
		Delay:   time.Second,
		Timeout: 10 * time.Second,
	}

	frame := lazy.New(session.Frame)

	session.NetworkIdle(time.Second*30, time.Millisecond*250, func() {
		err = frame.Navigate("https://mdemo.company.site/").Call(r)
		if err != nil {
			panic(err)
		}
	})

	text := frame.Query(`.cover__title a span`).InnerText().MustCall(r)
	log.Println(text, err)

	req := `function() { return new Promise( _r => requestIdleCallback(_r) ) }`

	b := frame.Query(`.cover__title a span`).CallFunctionOn(req).MustCall(r)
	log.Println(b)

	// p := session.Frame.Evaluate(`new Promise((a,b) => a('ok'))`, false).MustGetValue().(runtime.RemoteObjectId)
	// a, b := session.Frame.AwaitPromise(p)
	// log.Println(a, b)
}
