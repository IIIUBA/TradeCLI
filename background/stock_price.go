package background

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// func isNumber(s string) bool {
// 	_, err := strconv.Atoi(s)
// 	return err == nil
// }

func GetStockPrice(stockName string) float64 {
	var price float64
	counter := 0

	if !strings.HasSuffix(stockName, "F") {
		stockName += "F"
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.tbank.ru"),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Ошибка при запросе: %v\n", err)
	})

	c.OnHTML(".Money-module__money_UZBbh", func(e *colly.HTMLElement) {
		text := e.Text
		counter++
		if counter == 6 {
			cleanedText := strings.ReplaceAll(text, "\u00a0", "")    // Удаление неразрывного пробела
			cleanedText = strings.ReplaceAll(cleanedText, " ", "")   // Удаление обычных пробелов
			cleanedText = strings.ReplaceAll(cleanedText, "₽", "")   // Удаление символа рубля
			cleanedText = strings.ReplaceAll(cleanedText, ",", ".")  // Замена запятой на точку
			cleanedText = strings.ReplaceAll(cleanedText, "пт.", "") // Удаление "пт." для пунктов

			value, err := strconv.ParseFloat(cleanedText, 64)
			if err != nil {
				log.Printf("Ошибка преобразования строки '%s' в число: %v\n", cleanedText, err)
				return
			}
			price = value
		}
	})

	err := c.Visit(fmt.Sprintf("https://www.tbank.ru/invest/futures/%s", stockName))
	if err != nil {
		log.Printf("Ошибка посещения URL: %v\n", err)
	}

	return price
}
