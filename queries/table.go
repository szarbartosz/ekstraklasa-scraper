package queries

import (
	"fmt"
	"log"
	"os"
	"scraper/ekstraklasa/scrapers/ekstraklasa"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) {
	scrapeUrl := os.Getenv("EKSTRAKLASA_URL")

	if scrapeUrl == "" {
		log.Panic("No EKSTRAKLASA_URL env variable found!")
	}

	queryParams := c.Request.URL.Query()
	fmt.Println(queryParams)

	standings := ekstraklasa.ScrapeTable(scrapeUrl)

	c.JSON(200, gin.H{
		"standings": standings,
	})
}
