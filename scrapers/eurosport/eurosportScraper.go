package eurosport

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
				case 1:
					standing.Position = utils.ParseToInt(td.Text)
				case 3:
					standing.TeamName = td.ChildText("a.hidden")
					standing.LogoUrl = td.ChildAttr("source", "srcset")
				case 4:
					var lastResults []models.MatchResult

					td.ForEach("svg", func(i int, svg *colly.HTMLElement) {
						lastResults = append(lastResults, ResolveMatchResult(svg))

					})

					for i, j := 0, len(lastResults)-1; i < j; i, j = i+1, j-1 {
						lastResults[i], lastResults[j] = lastResults[j], lastResults[i]
					}

					if len(lastResults) < 5 {
						for i := len(lastResults); i < 5; i++ {
							lastResults = append(lastResults, models.Unknown)
						}
					}

					standing.LastResults = lastResults
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

	return standings
}

func ResolveMatchResult(svg *colly.HTMLElement) models.MatchResult {
	if strings.Contains(svg.Attr("data-testid"), "lost") {
		return models.Lost
	} else if strings.Contains(svg.Attr("data-testid"), "draw") {
		return models.Draw
	} else {
		return models.Won
	}
}
