package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Quotes struct {
	QUOTE  string
	AUTHOR string
	TAGS   string
}

func main() {

	citas := make([]Quotes, 0)

	c := colly.NewCollector(colly.AllowedDomains("quotes.toscrape.com", "www.quotes.toscrape.com/"))

	c.OnHTML("div.quote", func(e *colly.HTMLElement) {
		cita := Quotes{
			QUOTE:  e.ChildText("span.text"),
			AUTHOR: e.ChildText("small.author"),
			TAGS:   e.ChildText("div.tags > a"),
		}

		citas = append(citas, cita)
		jsonData, err := json.Marshal(citas)
		if err != nil {
			return
		}
		ioutil.WriteFile("Quotes.json", jsonData, 0644)

	})
	c.OnHTML(".next a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" {
			fmt.Println("Not fund link")
		}
		fmt.Printf("link fund : -> %s\n", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.Visit("https://quotes.toscrape.com/")
}
