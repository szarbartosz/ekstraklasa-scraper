package main

import (
	"os"
	"scraper/ekstraklasa/initializers"
	"scraper/ekstraklasa/queries"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	router := gin.Default()

	router.GET("/table", queries.GetTable)

	router.Run()
}
