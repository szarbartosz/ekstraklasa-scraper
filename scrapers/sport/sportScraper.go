package sport

import (
	"strings"

	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/utils"

	"github.com/gocolly/colly"
)

func ScrapeTable(scrapeUrl string) []models.Standing {
	collector := colly.NewCollector()
	headerSkipped := false

	var standings []models.Standing

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {
		standing := models.Standing{}

		if headerSkipped {
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				switch i {
				case 0:
					standing.Position = utils.ParseToInt(td.Text)
				case 2:
					standing.TeamName = td.ChildText("a span")
					standing.LogoUrl = td.ChildAttr("img", "data-url")

				case 3:
					standing.GamesPlayed = utils.ParseToInt(td.Text)

				case 4:
					standing.Wins = utils.ParseToInt(td.Text)
				case 5:
					standing.Draws = utils.ParseToInt(td.Text)
				case 6:
					standing.Losses = utils.ParseToInt(td.Text)

				case 7:
					parts := strings.Split(td.Text, ":")
					standing.GoalsFor = utils.ParseToInt(parts[0])
					standing.GoalsAgainst = utils.ParseToInt(parts[1])
					standing.GoalsDifference = standing.GoalsFor - standing.GoalsAgainst
				case 9:
					standing.TeamPoints = utils.ParseToInt(td.Text)

				case 10:
					var lastResults []models.MatchResult

					td.ForEach(".form__badge", func(i int, div *colly.HTMLElement) {
						lastResults = append(lastResults, ResolveMatchResult(div))
					})

					if len(lastResults) < 5 {
						for i := len(lastResults); i < 5; i++ {
							lastResults = append(lastResults, models.Unknown)
						}
					}

					standing.LastResults = lastResults

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

func ResolveMatchResult(div *colly.HTMLElement) models.MatchResult {
	if strings.Contains(div.Text, "P") {
		return models.Lost
	} else if strings.Contains(div.Text, "R") {
		return models.Draw
	} else {
		return models.Won
	}
}
