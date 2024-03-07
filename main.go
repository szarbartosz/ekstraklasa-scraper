package main

import (
	"os"
	"scraper/ekstraklasa/controllers"
	"scraper/ekstraklasa/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	router := gin.Default()

	router.GET("/table", controllers.GetTable)
	router.GET("/games/upcoming", controllers.GetUpcomingGames)

	router.Run()
}
