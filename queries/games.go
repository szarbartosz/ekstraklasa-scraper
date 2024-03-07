package queries

import (
	"log"
	"os"
	"scraper/ekstraklasa/scrapers/gol"
)

func QueryUpcomingGames(queryParams map[string][]string) {
	scrapeUrl := os.Getenv("GOL_GAMES_URL")

	if scrapeUrl == "" {
		log.Panic("No GOL_GAMES_URL env variable found!")
	}

	gol.ScrapeTable(scrapeUrl)

}
