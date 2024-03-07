package queries

import (
	"log"
	"os"
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/scrapers/gol"
	"scraper/ekstraklasa/utils"
	"sort"
	"strings"
	"time"
)

func QueryGames(queryParams map[string][]string) []models.Game {

	scrapeUrl := os.Getenv("GOL_URL")

	if scrapeUrl == "" {
		log.Panic("No GOL_URL env variable found!")
	}

	games := gol.ScrapeGames(scrapeUrl + "/terminarz")
	games = FilterByTeam(games, queryParams)
	games = FilterByRound(games, queryParams)
	games = FilterUpcomingGames(games, queryParams)

	sort.Sort(utils.GameSlice(games))

	return games
}

func FilterByTeam(games []models.Game, queryParams map[string][]string) []models.Game {
	teamNameParam := queryParams["teamName"]

	if len(teamNameParam) == 0 {
		return games
	}

	teamName := teamNameParam[0]

	var filteredGames []models.Game
	for _, game := range games {

		if strings.Contains(utils.SanitizeString(game.Host), utils.SanitizeString(teamName)) || strings.Contains(utils.SanitizeString(game.Guest), utils.SanitizeString(teamName)) {
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

func FilterUpcomingGames(games []models.Game, queryParams map[string][]string) []models.Game {
	roundParam := queryParams["upcoming"]

	if len(roundParam) == 0 {
		return games
	}

	oneWeek := time.Now().Add(7 * 24 * time.Hour)
	var filteredGames []models.Game

	for _, game := range games {
		if game.DateTime.After(time.Now()) && game.DateTime.Before(oneWeek) {
			filteredGames = append(filteredGames, game)
		}
	}

	return filteredGames
}
