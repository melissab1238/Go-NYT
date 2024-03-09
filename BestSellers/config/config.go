// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var APIKEY string

func init() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set the APIKEY variable from the .env file
	APIKEY = os.Getenv("APIKEY")

}
