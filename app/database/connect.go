package database

import (
	"Biocad/app/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client         *mongo.Client
	fileCollection *mongo.Collection
)

func Connect() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURL))
	if err != nil {
		log.Println(err.Error())
		return
	}

	fileCollection = client.Database("test").Collection("test")
	fmt.Println("Connected to MongoDB!")
}
