package db

import (
	"commerce/handlers/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB(cfg *config.Config) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d",
		cfg.MONGODB.HOST, cfg.MONGODB.PORT))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	// Return the connected client and collection for further use
	return client, nil
}
