package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Database struct {
	Database *mongo.Database
}


func NewDatabase(uri string) *Database {
	return &Database{connectMongoDB(uri)}
}


func  connectMongoDB(uri string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	return client.Database("product-api")
}
