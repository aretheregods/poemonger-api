package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func InitializeDB() *firestore.Client {
	config := &firebase.Config{ProjectID: "poemonger"}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return client
}
