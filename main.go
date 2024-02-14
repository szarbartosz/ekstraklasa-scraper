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

	r.GET("/table", scrapers.ScrapeTable)

	r.Run()
}
