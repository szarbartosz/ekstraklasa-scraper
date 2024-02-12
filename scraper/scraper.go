package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func Scrape() {
	c := colly.NewCollector()

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			testID := el.Attr("data-testid")
			fmt.Println(testID)
			if testID == "atom-table-cell-last-results" {
				el.ForEach("svg", func(i int, el *colly.HTMLElement) {
					fmt.Println(el.Attr("class"))
				})
			} else if testID == "table-cell-team" {
				el.ForEach("a", func(i int, el *colly.HTMLElement) {
					fmt.Println(el.Text)
				})
			} else if testID == "table-cell-icon" {
				fmt.Println(el)
				el.ForEach("source", func(i int, el *colly.HTMLElement) {
					fmt.Println(el.Attr("srcset"))
				})
			}
		})

		fmt.Println(e.Attr("table-cell-team"))
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scrapeUrl := os.Getenv("EKSTRAKLASA_URL")

	c.Visit(scrapeUrl)
}
