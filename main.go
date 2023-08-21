package main

import (
	"log"

	"github.com/joho/godotenv"

	"poemonger/api/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.InitializeAPI()
}
