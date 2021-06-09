package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoDatabaseEntry struct {
	ID    primitive.ObjectID `bson:"_id,omitempty`
	Video string             `bson:"title,omitempty`
}

// TODO replace with dependency injection
var Global_mongo_client *mongo.Client
var Global_mongo_context context.Context

func ConnectToMongo() {
	var err error
	Global_mongo_client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	panicWhenErr(err)
	Global_mongo_context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Global_mongo_client.Connect(Global_mongo_context)
	panicWhenErr(err)
	// defer client.Disconnect(Global_mongo_context)
	panicWhenErr(err)
	fmt.Println("Connected to Mongo")
}

func panicWhenErr(err error) {
	if err != nil {
		panic(err)
	}
}
