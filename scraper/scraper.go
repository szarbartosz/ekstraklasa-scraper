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
			switch i {
			case 1:
				fmt.Println("Team position:", td.Text)
			case 3:
				fmt.Println("Team name:", td.ChildText("a.hidden"))
			case 4:
				// TODO Handle these nice svgs
			case 5:
				fmt.Println("Games played:", td.Text)
			case 6:
				fmt.Println("Wins:", td.Text)
			case 7:
				fmt.Println("Draws:", td.Text)
			case 8:
				fmt.Println("Losses:", td.Text)
			case 9:
				fmt.Println("Goals for:", td.Text)
			case 10:
				fmt.Println("Goals against:", td.Text)
			case 11:
				fmt.Println("Goals difference:", td.Text)
			case 12:
				fmt.Println("Team points:", td.Text)
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
