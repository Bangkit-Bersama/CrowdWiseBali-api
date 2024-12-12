package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Production bool
var GMPAPIKey string
var InferenceServerUrl string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file! Ignoring...")
	}

	production, _ := strconv.Atoi(os.Getenv("PRODUCTION"))
	Production = production > 0
	GMPAPIKey = os.Getenv("GMP_API_KEY")
	InferenceServerUrl = os.Getenv("INFERENCE_SERVER_URL")
}
