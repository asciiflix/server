package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/asciiflix/server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var global_mongo_client *mongo.Client
var global_mongo_context context.Context

func ConnectToMongo() {
	var err error

	dsn := url.URL{
		User: url.UserPassword(config.MongoDB.User, config.MongoDB.Password),
		Host: fmt.Sprintf("%s:%s", config.MongoDB.Host, config.MongoDB.Port),
	}

	//Create Client
	global_mongo_client, err = mongo.NewClient(options.Client().ApplyURI("mongodb:" + dsn.String()))

	//Panic
	panicWhenErr(err)

	//Context

	global_mongo_context = context.Background()

	err = global_mongo_client.Connect(global_mongo_context)
	panicWhenErr(err)
	// defer client.Disconnect(Global_mongo_context)

	config.Log.Info("Connected to MongoDB")
}

func panicWhenErr(err error) {
	if err != nil {
		panic(err)
	}
}
