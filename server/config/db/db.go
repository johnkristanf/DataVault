package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify that the client has connected successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func UserCollection() *mongo.Collection {
	client := InitMongoDB()

	if client == nil {
		log.Fatal("Cannot Connect to the Database")
		return nil
	}

	return client.Database("DataVault").Collection("users")
}
