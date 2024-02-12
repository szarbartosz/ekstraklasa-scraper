package main

import (
	"scraper/ekstraklasa/models"
	"scraper/ekstraklasa/scraper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/table", func(c *gin.Context) {
		scraper.Scrape()

		c.JSON(200, gin.H{
			"standings": []models.Standing{
				{
					Position: "1",
					Team:     "Team A",
					Points:   45,
				},
			},
		})
	})
	r.Run()
}
