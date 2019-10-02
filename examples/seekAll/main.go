package main

import (
	"time"

	"github.com/ecwid/witness/pkg/chrome"
	"github.com/ecwid/witness/pkg/log"
)

func main() {
	log.Logging = log.LevelFatal

	chrome, _ := chrome.New()
	defer chrome.Close()
	p, err := chrome.CDP.DefaultPage()
	if err != nil {
		panic(err)
	}

	// Implicitly affected only `Expect` function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5

	p.Navigate("https://mdemo.ecwid.com/")
	doc := p.Doc()

	// expected element with visibility = true must be found else panic
	exp := doc.Expect(".ec-static-container .grid-product", true)
	// exp never nil
	exp.Release()

	for _, card := range doc.SeekAll(".ec-static-container .grid-product") {
		titleElement, err := card.Seek(".grid-product__title-inner")
		if err != nil {
			panic("title is not exist")
		}
		title, err := titleElement.GetText()
		if err != nil {
			panic("can't read title")
		}
		priceElement, err := card.Seek(".grid-product__price-amount")
		if err != nil {
			panic("price is not exist")
		}
		price, err := priceElement.GetText()
		if err != nil {
			panic("can't read price")
		}
		log.Printf(log.LevelFatal, "title = %s, price = %s", title, price)
	}

}
