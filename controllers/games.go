package controllers

import (
	"scraper/ekstraklasa/queries"

	"github.com/gin-gonic/gin"
)

func GetUpcomingGames(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	games := queries.QueryUpcomingGames(queryParams)

	c.JSON(200, gin.H{
		"games": games,
	})
}
