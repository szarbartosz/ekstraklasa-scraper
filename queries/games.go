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
	games = FilterByGuest(games, queryParams)
	games = FilterByRound(games, queryParams)

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

func FilterByGuest(games []models.Game, queryParams map[string][]string) []models.Game {
	guestNameParam := queryParams["guest"]

	if len(guestNameParam) == 0 {
		return games
	}

	guestName := guestNameParam[0]

	var filteredGames []models.Game
	for _, game := range games {

		if strings.Contains(utils.SanitizeString(game.Guest), utils.SanitizeString(guestName)) {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}

func FilterByRound(games []models.Game, queryParams map[string][]string) []models.Game {
	roundParam := queryParams["round"]

	if len(roundParam) == 0 {
		return games
	}

	round := utils.ParseToInt(roundParam[0])

	var filteredGames []models.Game
	for _, game := range games {

		if game.Round == round {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}
