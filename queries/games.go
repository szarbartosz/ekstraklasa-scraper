package queries

import (
	"log"
	"os"
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/scrapers/gol"
)

func QueryGames(queryParams map[string][]string) []models.Game {
	scrapeUrl := os.Getenv("GOL_GAMES_URL")

	if scrapeUrl == "" {
		log.Panic("No GOL_GAMES_URL env variable found!")
	}

	return gol.ScrapeGames(scrapeUrl + "/terminarz")
}
