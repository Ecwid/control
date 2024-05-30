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
	session, dfr, err := control.TakeWithContext(context.TODO(), sl, "--no-startup-window")
	if err != nil {
		panic(err)
	}
	defer dfr()

	err = session.Frame.Navigate("https://zoid.ecwid.com")
	if err != nil {
		panic(err)
	}

	retrier := retry.DefaultTiming

	var values []string
	err = retry.Func(retrier, func() error {
		values = []string{}
		return session.Frame.QueryAll(".grid-product__title-inner").Then(func(nl control.NodeList) error {
			return nl.Foreach(func(n *control.Node) error {
				return n.GetText().Then(func(s string) error {
					values = append(values, s)
					return nil
				})
			})
		})
	})

	log.Println(values, err)

	err = retry.FuncPanic(retrier, func() {
		node := session.Frame.MustQuery(`.pager__count-pages`)
		node.MustGetBoundingClientRect()
		node.MustClick()
	})
	log.Println(err)

	p := session.Frame.Evaluate(`new Promise((a,b) => a('ok'))`, false).MustGetValue().(control.RemoteObject)
	a, b := session.Frame.AwaitPromise(p)
	log.Println(a, b)
}
