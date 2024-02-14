package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("No .env file found - that's okay if running on docker tho.")
		} else {
			log.Print("Error loading .env file!")
		}
	}
}
