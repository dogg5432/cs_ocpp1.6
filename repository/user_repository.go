package repository

import (
	"context"

	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(ctx context.Context, user *models.User) error {
	_, err := database.Client.Collection("users").InsertOne(ctx, user)
	return err
}

func GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := database.Client.Collection("users").FindOne(ctx, bson.M{"userId": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByRFID(ctx context.Context, rfidTag string) (*models.User, error) {
	var user models.User
	err := database.Client.Collection("users").FindOne(ctx, bson.M{"rfidTag": rfidTag}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
