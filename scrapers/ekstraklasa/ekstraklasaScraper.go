package ekstraklasa

import (
	"strings"

	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ScrapeTable(scrapeUrl string) []models.Standing {
	collector := colly.NewCollector()

	var standings []models.Standing

	collector.OnHTML("app-league-standings-entry", func(tableRow *colly.HTMLElement) {
		standing := models.Standing{}
		var lastResults []models.MatchResult

		tableRow.DOM.Find("div span").Each(func(i int, span *goquery.Selection) {
			switch i {
			case 0:
				standing.Position = utils.ParseToInt(span.Text())
			case 2:
				standing.TeamName = span.Text()
			}
		})

		tableRow.DOM.Find("img").Each(func(i int, img *goquery.Selection) {
			switch i {
			case 0:
				standing.LogoUrl, _ = img.Attr("src")
			}
		})

		tableRow.DOM.Find("div.standings-data span").Each(func(i int, span *goquery.Selection) {
			switch i {
			case 0:
				standing.GamesPlayed = utils.ParseToInt(span.Text())
			case 1:
				standing.TeamPoints = utils.ParseToInt(span.Text())
			case 2:
				standing.Wins = utils.ParseToInt(span.Text())
			case 3:
				standing.Draws = utils.ParseToInt(span.Text())
			case 4:
				standing.Losses = utils.ParseToInt(span.Text())
			case 5:
				var dividerIndex = strings.Index(span.Text(), ":")
				var goalsFor = span.Text()[:dividerIndex]
				var goalsAgainst = span.Text()[dividerIndex+1:]
				standing.GoalsFor = utils.ParseToInt(goalsFor)
				standing.GoalsAgainst = utils.ParseToInt(goalsAgainst)
				standing.GoalsDifference = standing.GoalsFor - standing.GoalsAgainst
			}
		})

		tableRow.DOM.Find("app-last-five-matches div").Each(func(i int, item *goquery.Selection) {
			lastResults = append(lastResults, ResolveMatchResult(item))
		})

		standing.LastResults = lastResults
		standings = append(standings, standing)
	})

	collector.Visit(scrapeUrl)

	return standings
}

func ResolveMatchResult(div *goquery.Selection) models.MatchResult {
	var val, exists = div.Attr("class")

	if exists {
		if strings.Contains(val, "red") {
			return models.Lost
		} else if strings.Contains(val, "gray") {
			return models.Draw
		} else {
			return models.Won
		}
	}
	return models.Draw
}
