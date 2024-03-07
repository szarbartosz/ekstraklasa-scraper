package queries

import (
	"log"
	"os"
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/scrapers/gol"
	"scraper/ekstraklasa/utils"
	"strings"
)

func QueryGames(queryParams map[string][]string) []models.Game {
	scrapeUrl := os.Getenv("GOL_URL")

	if scrapeUrl == "" {
		log.Panic("No GOL_URL env variable found!")
	}

	games := gol.ScrapeGames(scrapeUrl + "/terminarz")
	games = FilterByHost(games, queryParams)

	return games
}

func FilterByHost(games []models.Game, queryParams map[string][]string) []models.Game {
	hostNameParam := queryParams["host"]

	if len(hostNameParam) == 0 {
		return games
	}

	hostName := hostNameParam[0]

	var filteredGames []models.Game
	for _, game := range games {

		if strings.Contains(utils.SanitizeString(game.Host), utils.SanitizeString(hostName)) {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}
