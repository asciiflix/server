package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/asciiflix/server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO replace with dependency injection
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
	Global_mongo_context, cancell_ctx := context.WithTimeout(context.Background(), 10*time.Second)

	err = global_mongo_client.Connect(Global_mongo_context)
	panicWhenErr(err)
	// defer client.Disconnect(Global_mongo_context)
	
	config.Log.Info("Connected to MongoDB")
	cancell_ctx();
}

func panicWhenErr(err error) {
	if err != nil {
		panic(err)
	}
}
