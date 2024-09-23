package repository

import (
	"context"

	"github.com/dogg5432/central_charger/database"

	"github.com/dogg5432/central_charger/models"
)

func CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := database.Client.Collection("transactions").InsertOne(ctx, transaction)
	return err
}
