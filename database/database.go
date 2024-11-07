package database

import (
	"context"
	"fmt"

	"github.com/dogg5432/central_charger/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Database

func Connect() error {
	// Set client options and connect to MongoDB
	clientOptions := options.Client().ApplyURI(config.ConfigApp.Database.Uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	// defer client.Disconnect(context.TODO())

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	Client = client.Database("central_system")

	return nil
}
