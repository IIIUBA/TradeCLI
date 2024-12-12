package background

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetStockPrice(stockName string) float64 {
	var price float64
	const targetElementIndex = 6
	counter := 0

	re := regexp.MustCompile(`[^\d,.]`)

	c := colly.NewCollector(
		colly.AllowedDomains("www.tbank.ru"),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Ошибка при запросе: %v\n", err)
	})

	c.OnHTML(".Money-module__money_UZBbh", func(e *colly.HTMLElement) {
		counter++
		if counter == targetElementIndex {

			cleanedText := re.ReplaceAllString(e.Text, "")
			cleanedText = strings.ReplaceAll(cleanedText, ",", ".") 

			value, err := strconv.ParseFloat(cleanedText, 64)
			if err != nil {
				log.Printf("Ошибка преобразования текста '%s': %v\n", e.Text, err)
				return
			}

			price = value
		}
	})

	url := fmt.Sprintf("https://www.tbank.ru/invest/futures/%s", stockName)
	if err := c.Visit(url); err != nil {
		log.Printf("Ошибка посещения URL: %v\n", err)
	}

	return price
}
