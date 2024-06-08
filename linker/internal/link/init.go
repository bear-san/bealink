package link

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var client *mongo.Client
var database *mongo.Database
var collection *mongo.Collection

func init() {
	ctx := context.Background()

	dbUri := os.Getenv("DB_URL")
	if dbUri == "" {
		panic("DB_URL is required")
	}

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	database = client.Database("bealink")
	collection = database.Collection("links")
}
