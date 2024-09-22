package handlers

import (
	"context"
	"fmt"
	"time"
	
	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/models"
	"github.com/dogg5432/central_charger/repository"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

type ChargingStationHandler struct{}

func (h *ChargingStationHandler) OnBootNotification(chargePoint string, request *core.BootNotificationRequest) (confirmation *core.BootNotificationConfirmation, error error) {
	fmt.Printf("OnBootNotification => %s %v\n", chargePoint, request)
	chargePointModel,err := repository.GetChargePointByID(context.Background(), chargePoint)
	if err != nil {
		fmt.Println(err)
		return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), config.ConfigApp.Server.HeartbeatInterval, core.RegistrationStatusRejected), err
	}
	if chargePointModel != nil {
		return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), config.ConfigApp.Server.HeartbeatInterval, core.RegistrationStatusAccepted), nil
	}
	chargePointModel = &models.ChargePoint{
		ChargePointID: chargePoint,
		Vendor:        request.ChargePointVendor,
		Model:         request.ChargePointModel,
		Status:        "Available",
		Connectors:    []models.Connector{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = repository.CreateChargePoint(context.Background(), chargePointModel)
	if err != nil {
		fmt.Println(err)
		return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), config.ConfigApp.Server.HeartbeatInterval, core.RegistrationStatusRejected), err
	}
	return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), config.ConfigApp.Server.HeartbeatInterval, core.RegistrationStatusAccepted), nil
}

func (h *ChargingStationHandler) OnMeterValues(chargePoint string, request *core.MeterValuesRequest) (confirmation *core.MeterValuesConfirmation, error error) {
	fmt.Printf("OnMeterValues => %s %v\n", chargePoint, request)
	return core.NewMeterValuesConfirmation(), nil
}

func (h *ChargingStationHandler) OnAuthorize(chargePoint string, request *core.AuthorizeRequest) (confirmation *core.AuthorizeConfirmation, error error) {
	fmt.Printf("OnAuthorize => %s %v\n", chargePoint, request)
	return core.NewAuthorizationConfirmation(&types.IdTagInfo{}), nil
}

func (h *ChargingStationHandler) OnStatusNotification(chargePoint string, request *core.StatusNotificationRequest) (confirmation *core.StatusNotificationConfirmation, error error) {
	fmt.Printf("OnStatusNotification => %s %v\n", chargePoint, request)
	return core.NewStatusNotificationConfirmation(), nil
}

func (h *ChargingStationHandler) OnHeartbeat(chargePoint string, request *core.HeartbeatRequest) (confirmation *core.HeartbeatConfirmation, error error) {
	fmt.Printf("OnHeartbeat => %s %v\n", chargePoint, request)
	return core.NewHeartbeatConfirmation(types.NewDateTime(time.Now())), nil
}

func (h *ChargingStationHandler) OnDataTransfer(chargePoint string, request *core.DataTransferRequest) (confirmation *core.DataTransferConfirmation, error error) {
	fmt.Printf("OnDataTransfer => %s %v\n", chargePoint, request)
	return core.NewDataTransferConfirmation(core.DataTransferStatusAccepted), nil
}

func (h *ChargingStationHandler) OnStartTransaction(chargePoint string, request *core.StartTransactionRequest) (confirmation *core.StartTransactionConfirmation, error error) {
	fmt.Printf("OnStartTransaction => %s %v\n", chargePoint, request)
	return core.NewStartTransactionConfirmation(&types.IdTagInfo{}, 123), nil
}

func (h *ChargingStationHandler) OnStopTransaction(chargePoint string, request *core.StopTransactionRequest) (confirmation *core.StopTransactionConfirmation, error error) {
	fmt.Printf("OnStopTransaction => %s %v\n", chargePoint, request)
	return core.NewStopTransactionConfirmation(), nil
}
