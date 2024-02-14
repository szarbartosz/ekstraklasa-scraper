package main

import (
	"scraper/ekstraklasa/initializers"
	"scraper/ekstraklasa/scrapers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	// r.GET("/table", func(c *gin.Context) {
	// 	standings := scraper.Scrape()

	// 	c.JSON(200, gin.H{
	// 		"standings": standings,
	// 	})
	// })
	r.GET("/table", scrapers.ScrapeTable)

	r.Run()
}
