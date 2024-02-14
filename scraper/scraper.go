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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file - ignore this if running on docker tho.")
	}

	scrapeUrl := os.Getenv("EKSTRAKLASA_URL")

	collector := colly.NewCollector()
	headerSkipped := false

	var standings []models.Standing

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		standing := models.Standing{}

		if headerSkipped {
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				switch i {
				case 1:
					standing.Position = ParseToInt(td.Text)
				case 3:
					standing.TeamName = td.ChildText("a.hidden")
				case 4:
					// TODO Handle these nice svgs
				case 5:
					standing.GamesPlayed = ParseToInt(td.Text)
				case 6:
					standing.Wins = ParseToInt(td.Text)
				case 7:
					standing.Draws = ParseToInt(td.Text)
				case 8:
					standing.Losses = ParseToInt(td.Text)
				case 9:
					standing.GoalsFor = ParseToInt(td.Text)
				case 10:
					standing.GoalsAgainst = ParseToInt(td.Text)
				case 11:
					standing.GoalsDifference = ParseToInt(td.Text)
				case 12:
					standing.TeamPoints = ParseToInt(td.Text)
				}
			})
			standings = append(standings, standing)
		} else {
			headerSkipped = true
		}
	})

	collector.Visit(scrapeUrl)

	return standings
}

func ParseToInt(toBeParsed string) int {
	var parsed int
	var err error

	parsed, err = strconv.Atoi(toBeParsed)

	if err != nil {
		log.Fatal("Error while parsing to int: ", err)
	}

	return parsed
}
