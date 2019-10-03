package main

import (
	"log"
	"time"

	"github.com/ecwid/witness"
	"github.com/ecwid/witness/pkg/chrome"
)

func main() {
	chrome, _ := chrome.New()
	defer chrome.Close()
	p, err := chrome.CDP.DefaultPage()
	if err != nil {
		panic(err)
	}

	// Implicitly affected only `Expect` function
	chrome.CDP.Timeouts.Implicitly = time.Second * 5
	// set logging level and hook
	chrome.CDP.Logging.Level = witness.LevelProtocolErrors
	chrome.CDP.Logging.SetHook(func(line string) {
		log.Printf(line)
	})

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
		log.Printf("title = %s, price = %s", title, price)
	}

}
