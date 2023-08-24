package main

import (
	"log"

	"github.com/joho/godotenv"

	"poemonger/api/api"
	"poemonger/api/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	poemonger, client := db.InitializeDB("poemonger")
	api.InitializeAPI(poemonger, client)
}
