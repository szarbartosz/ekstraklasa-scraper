package controllers

import (
	"scraper/ekstraklasa/queries"

	"github.com/gin-gonic/gin"
)

func GetTable(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	standings := queries.QueryStandings(queryParams)

	c.JSON(200, gin.H{
		"standings": standings,
	})
}
