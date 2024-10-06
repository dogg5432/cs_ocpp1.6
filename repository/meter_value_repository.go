package repository

import (
	"context"

	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Save MeterValues data to MongoDB
func SaveMeterValuesToDB(chargePointId string, connectorId int,transactionID*int,sampledValue *models.SampledValue) error {
    // Convert meter values to BSON
    data := bson.M{
        "chargePointId": chargePointId,
        "connectorId":   connectorId,
        "transactionId": transactionID,
        "meterValues":   sampledValue,
    }

    _, err := database.Client.Collection("meter_values").InsertOne(context.Background(), data)
    if err != nil {
        return err
    }

    return nil
}