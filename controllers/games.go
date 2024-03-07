package controllers

import (
	"scraper/ekstraklasa/queries"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	games := queries.QueryGames(queryParams)

	c.JSON(200, gin.H{
		"games": games,
	})
}
