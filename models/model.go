package models

import (
	"time"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChargePoint represents the charge point (charging station)
type ChargePoint struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty"`
	ChargePointID string                 `bson:"chargePointId"`
	Vendor        string                 `bson:"vendor"`
	Model         string                 `bson:"model"`
	Status        core.ChargePointStatus `bson:"status"`
	Connectors    []Connector            `bson:"connectors"`
	CreatedAt     time.Time              `bson:"createdAt"`
	UpdatedAt     time.Time              `bson:"updatedAt"`
}

// Connector represents a single connector at a charge point
type Connector struct {
	ConnectorID int                    `bson:"connectorId"`
	Type        string                 `bson:"type"`
	Status      core.ChargePointStatus `bson:"status"`
	Power       int                    `bson:"power"` // kW
}

// User represents the user information
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	RFIDTag   string             `bson:"rfidTag"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

// Transaction represents a charging session transaction
type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	TransactionID  int                `bson:"transactionId"`
	ChargePointID  string             `bson:"chargePointId"`
	ConnectorID    int                `bson:"connectorId"`
	UserID         string             `bson:"userId"`
	StartTime      time.Time          `bson:"startTime"`
	StopTime       time.Time          `bson:"stopTime"`
	MeterStart     int                `bson:"meterStart"`
	MeterStop      int                `bson:"meterStop"`
	EnergyConsumed int                `bson:"energyConsumed"`
	Status         string             `bson:"status"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
