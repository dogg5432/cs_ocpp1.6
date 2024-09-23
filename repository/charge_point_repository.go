package repository

import (
	"context"
	"errors"

	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChargePoint(ctx context.Context, chargePoint *models.ChargePoint) error {
	_, err := database.Client.Collection("charge_points").InsertOne(ctx, chargePoint)
	return err
}

func GetChargePointByID(ctx context.Context, chargePointID string) (*models.ChargePoint, error) {
	var chargePoint models.ChargePoint
	err := database.Client.Collection("charge_points").FindOne(ctx, bson.M{"chargePointId": chargePointID}).Decode(&chargePoint)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chargePoint, nil
}

func UpdateChargePoint(ctx context.Context, chargePointID string, chargePoint *models.ChargePoint) error {
	_, err := database.Client.Collection("charge_points").UpdateOne(ctx, bson.M{"chargePointId": chargePointID}, bson.M{"$set": chargePoint})
	return err
}
