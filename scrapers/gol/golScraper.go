package gol

import (
	"fmt"
	"scraper/ekstraklasa/models"

	"github.com/gocolly/colly"
)

func ScrapeTable(scrapeUrl string) []models.Standing {
	collector := colly.NewCollector()

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		switch tr.Attr("class") {
		case "kolejkaData":
			fmt.Println(tr.ChildText(".rozgrywka"))
		}
	})

	collector.Visit(scrapeUrl)

	return []models.Standing{}
}
