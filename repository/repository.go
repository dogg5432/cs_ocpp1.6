package repository

import (
	"context"

	"github.com/dogg5432/central_charger/models"
)

// UserRepository defines methods to interact with user data.
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	GetUserByRFID(ctx context.Context, rfidTag string) (*models.User, error)
}

// ChargePointRepository defines methods to interact with charge point data.
type ChargePointRepository interface {
	CreateChargePoint(ctx context.Context, cp *models.ChargePoint) error
	GetChargePointByID(ctx context.Context, chargePointID string) (*models.ChargePoint, error)
	UpdateChargePointStatus(ctx context.Context, chargePointID string, cp *models.ChargePoint) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, trnx *models.Transaction) error
	GetTransactionByID(ctx context.Context, transactionID int) (*models.Transaction, error)
}
