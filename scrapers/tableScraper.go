package scrapers

import (
	"log"
	"os"

	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/utils"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func ScrapeTable(c *gin.Context) {
	scrapeUrl := os.Getenv("SCRAPE_URL")

	if scrapeUrl == "" {
		log.Panic("No SCRAPE_URL env variable found!")
	}

	collector := colly.NewCollector()
	headerSkipped := false

	var standings []models.Standing

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		standing := models.Standing{}

		if headerSkipped {
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				switch i {
				case 1:
					standing.Position = utils.ParseToInt(td.Text)
				case 3:
					standing.TeamName = td.ChildText("a.hidden")
				case 4:
					// TODO Handle these nice svgs
				case 5:
					standing.GamesPlayed = utils.ParseToInt(td.Text)
				case 6:
					standing.Wins = utils.ParseToInt(td.Text)
				case 7:
					standing.Draws = utils.ParseToInt(td.Text)
				case 8:
					standing.Losses = utils.ParseToInt(td.Text)
				case 9:
					standing.GoalsFor = utils.ParseToInt(td.Text)
				case 10:
					standing.GoalsAgainst = utils.ParseToInt(td.Text)
				case 11:
					standing.GoalsDifference = utils.ParseToInt(td.Text)
				case 12:
					standing.TeamPoints = utils.ParseToInt(td.Text)
				}
			})
			standings = append(standings, standing)
		} else {
			headerSkipped = true
		}
	})

	collector.Visit(scrapeUrl)

	c.JSON(200, gin.H{
		"standings": standings,
	})
}
