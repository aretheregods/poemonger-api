package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDB(dbName string) (*mongo.Database, *mongo.Client) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	var password string
	var username string
	var location string
	if password = os.Getenv("MONGODB_PASSWORD"); password == "" {
		log.Fatal("You must set your 'MONGODB_PASSWORD' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if username = os.Getenv("MONGODB_USERNAME"); username == "" {
		log.Fatal("You must set your 'MONGODB_USERNAME' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if location = os.Getenv("MONGODB_LOCATION"); location == "" {
		log.Fatal("You must set your 'MONGODB_LOCATION' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	mongoURL := fmt.Sprintf("mongodb+srv://%v:%v@%v/?retryWrites=true&w=majority", username, password, location)
	clientOptions := options.Client().
		ApplyURI(mongoURL).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := client.Database("poemonger").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	return client.Database(dbName), client
}
