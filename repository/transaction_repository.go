package repository

import (
	"context"

	"github.com/dogg5432/central_charger/database"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/dogg5432/central_charger/models"
)

func CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := database.Client.Collection("transactions").InsertOne(ctx, transaction)
	return err
}

func GetTransactionByID(ctx context.Context, transactionID int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := database.Client.Collection("transactions").FindOne(ctx, bson.M{"transactionId": transactionID}).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func UpdateTransaction(ctx context.Context, transactionID int, transaction *models.Transaction) error {
	_, err := database.Client.Collection("transactions").UpdateOne(ctx, bson.M{"transactionId": transactionID}, bson.M{"$set": transaction})
	return err
}
