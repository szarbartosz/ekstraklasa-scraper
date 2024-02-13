package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func Scrape() {
	collector := colly.NewCollector()

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		tr.ForEach("td", func(i int, td *colly.HTMLElement) {
			if td.Attr("data-testid") == "table-cell-value" && i == 1 {
				fmt.Println("Team position", td.Text)
			} else if td.Attr("data-testid") == "table-cell-team" && i == 3 {
				fmt.Println("Team name", td.ChildText("a.hidden"))
			} else if td.Attr("data-testid") == "atom-table-cell-last-results" && i == 4 {
				// TODO Handle these nice svgs
			} else if td.Attr("data-testid") == "table-cell-value" && i == 5 {
				fmt.Println("Games played", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 6 {
				fmt.Println("Wins", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 7 {
				fmt.Println("Draws", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 8 {
				fmt.Println("Losses", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 9 {
				fmt.Println("Goals for", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 10 {
				fmt.Println("Goals against", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 11 {
				fmt.Println("Goals difference", td.Text)
			} else if td.Attr("data-testid") == "table-cell-value" && i == 12 {
				fmt.Println("Team points", td.Text)
			}
		})
		fmt.Println("\n")
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scrapeUrl := os.Getenv("EKSTRAKLASA_URL")

	collector.Visit(scrapeUrl)
}
