package queries

import (
	"fmt"
	"log"
	"os"
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/scrapers/ekstraklasa"
	"scraper/ekstraklasa/scrapers/eurosport"
	"scraper/ekstraklasa/utils"
	"strings"
)

func QueryStandings(queryParams map[string][]string) []models.Standing {
	activeScraper := os.Getenv("ACTIVE_SCRAPER")

	urlEnvName := activeScraper + "_URL"
	scrapeUrl := os.Getenv(urlEnvName)

	if scrapeUrl == "" {
		message := fmt.Sprintf("No %s env variable found!", urlEnvName)
		log.Panic(message)
	}

	var standings []models.Standing

	switch activeScraper {
	case "EKSTRAKLASA":
		standings = ekstraklasa.ScrapeTable(scrapeUrl)
	case "EUROSPORT":
		standings = eurosport.ScrapeTable(scrapeUrl)
	default:
		message := fmt.Sprintf("No %s env variable found!", activeScraper)
		log.Panic(message)
		standings = []models.Standing{}
	}

	standings = FilterByTeamName(standings, queryParams)

	return standings
}

func FilterByTeamName(standings []models.Standing, queryParams map[string][]string) []models.Standing {
	teamNameParam := queryParams["teamName"]

	if len(teamNameParam) == 0 {
		return standings
	}

	teamName := teamNameParam[0]

	var filteredStandings []models.Standing
	for _, standing := range standings {

		if strings.Contains(utils.SanitizeString(standing.TeamName), utils.SanitizeString(teamName)) {
			filteredStandings = append(filteredStandings, standing)
		}
	}

	return filteredStandings
}
