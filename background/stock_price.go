package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{"www.tbank.ru"}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("------Visiting------")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML(".Money-module__money_UZBbh", func(e *colly.HTMLElement) { fmt.Println(e.Text) })

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("------Finished------")
	})

	c.Visit("https://www.tbank.ru/invest/futures/SBERF/")
}
