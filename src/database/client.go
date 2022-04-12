package database

import (
	"context"
	"log"

	"github.com/Neoration/dot-canvas/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client  *mongo.Client
	Canvas  *mongo.Collection
	Users   *mongo.Collection
	Records *mongo.Collection
)

func init() {
	var err error

	if Client, err = mongo.Connect(
		context.Background(), options.Client().ApplyURI(config.MongoDBURI)); err != nil {
		log.Fatalf("Error: %v", err)
	}

	if err = Client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("MongoDB Connected.")

	Canvas = Client.Database(config.DatabaseName).Collection("canvas")
	Users = Client.Database(config.DatabaseName).Collection("users")
	Records = Client.Database(config.DatabaseName).Collection("records")
}
