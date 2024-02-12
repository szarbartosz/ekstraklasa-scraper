package main

import "github.com/gin-gonic/gin"

type Standings struct {
	Position string
	Team     string
	Points   int
}

func main() {
	r := gin.Default()
	r.GET("/table", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"standings": []Standings{
				{
					Position: "1",
					Team:     "Team A",
					Points:   45,
				},
				{
					Position: "2",
					Team:     "Team B",
					Points:   43,
				},
			},
		})
	})
	r.Run()
}
