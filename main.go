package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"poemonger/api/api"
	"poemonger/api/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	poemonger, client := db.InitializeDB("poemonger")
	api.InitializeAPI(poemonger)

	defer closeDBConnection(client)
}

func closeDBConnection(c *mongo.Client) {
	if err := c.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
