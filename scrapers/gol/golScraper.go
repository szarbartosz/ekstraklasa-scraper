package gol

import (
	"log"
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/utils"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeGames(scrapeUrl string) []models.Game {
	collector := colly.NewCollector()

	var games []models.Game
	currentRound := -1
	currentDate := ""
	headerSkipped := false

	collector.OnHTML("tr", func(tr *colly.HTMLElement) {

		switch tr.Attr("class") {
		case "kolejkaData":
			roundText := tr.ChildText(".rozgrywka")
			currentRound = ParseRoundText(roundText)

		case "dzien":
			currentDate = strings.ReplaceAll(tr.Text, "\n", "")

		case "spotkanie":
			time := tr.ChildText(".godzina")
			result := tr.ChildText(".wynik")
			host := ""
			guest := ""

			tr.ForEach(".nazwa", func(i int, a *colly.HTMLElement) {
				switch i {
				case 0:
					host = a.Text
				case 1:
					guest = a.Text
				}
			})

			parsedDateTime, err := utils.ParseDateTime(time, currentDate)

			if err != nil {
				log.Panic("Error while parsing date and time: ", err)
			}

			game := models.Game{
				Round:    currentRound,
				DateTime: parsedDateTime,
				Host:     host,
				Guest:    guest,
				Result:   result,
			}
			if headerSkipped {
				games = append(games, game)
			} else {
				headerSkipped = true
			}

		}

	})

	collector.Visit(scrapeUrl)

	return games
}

func ParseRoundText(roundText string) int {
	if roundText == "" {
		return -1
	}
	parsedString := strings.ReplaceAll(roundText, ". kolejka", "")

	return utils.ParseToInt(parsedString)
}
