package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDB(dbName string) *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	var uri string
	if uri = os.Getenv("MONGODB_PASSWORD"); uri == "" {
		log.Fatal("You must set your 'MONGODB_PASSWPRD' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	mongoURL := fmt.Sprintf("mongodb+srv://aretheregods:%v@poems.o87xo5v.mongodb.net/?retryWrites=true&w=majority", uri)
	clientOptions := options.Client().
		ApplyURI(mongoURL).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return client.Database(dbName)
}

