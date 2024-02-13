package scraper

import (
	"log"
	"os"
	"strconv"

	"scraper/ekstraklasa/models"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func Scrape() []models.Standing {
	collector := colly.NewCollector()
	var err error
	var standings []models.Standing
	headerSkipped := false

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		standing := models.Standing{}

		if headerSkipped {
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				switch i {
				case 1:
					standing.Position, err = strconv.Atoi(td.Text)
				case 3:
					standing.TeamName = td.ChildText("a.hidden")
				case 4:
					// TODO Handle these nice svgs
				case 5:
					standing.GamesPlayed, err = strconv.Atoi(td.Text)
				case 6:
					standing.Wins, err = strconv.Atoi(td.Text)
				case 7:
					standing.Draws, err = strconv.Atoi(td.Text)
				case 8:
					standing.Losses, err = strconv.Atoi(td.Text)
				case 9:
					standing.GoalsFor, err = strconv.Atoi(td.Text)
				case 10:
					standing.GoalsAgainst, err = strconv.Atoi(td.Text)
				case 11:
					standing.GoalsDifference, err = strconv.Atoi(td.Text)
				case 12:
					standing.TeamPoints, err = strconv.Atoi(td.Text)
				}
			})
			if err != nil {
				log.Fatal("Error while parsing standing: ", err)
			} else {
				standings = append(standings, standing)
			}
		} else {
			headerSkipped = true
		}
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scrapeUrl := os.Getenv("EKSTRAKLASA_URL")

	collector.Visit(scrapeUrl)

	return standings
}
