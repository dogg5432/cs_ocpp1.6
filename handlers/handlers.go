package handlers

import (
	"context"
	"fmt"
	"slices"
	"strconv"
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
	chargePointModel, err := repository.GetChargePointByID(context.Background(), chargePoint)
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
	fmt.Printf("Received MeterValues from Charge Point %s", chargePoint)

	// Iterate over MeterValues in the request
	for _, meterValue := range request.MeterValue {
		fmt.Printf("Timestamp: %s", meterValue.Timestamp)

		// Loop through SampledValues to log or process the data

		for _, sampledValue := range meterValue.SampledValue {
			value, _ := strconv.Atoi(sampledValue.Value)
			sampleValueData := models.SampledValue{
				Measurand: sampledValue.Measurand,
				Value:     value,
				Unit:      sampledValue.Unit,
			}
			fmt.Printf("Measurand: %s, Value: %s %s", sampledValue.Measurand, sampledValue.Value, sampledValue.Unit)
			repository.SaveMeterValuesToDB(chargePoint, request.ConnectorId, request.TransactionId, &sampleValueData)
		}
	}

	// Return confirmation response
	return core.NewMeterValuesConfirmation(), nil
}

func (h *ChargingStationHandler) OnAuthorize(chargePoint string, request *core.AuthorizeRequest) (confirmation *core.AuthorizeConfirmation, error error) {
	fmt.Printf("OnAuthorize => %s %v\n", chargePoint, request)
	return core.NewAuthorizationConfirmation(&types.IdTagInfo{Status: types.AuthorizationStatusAccepted}), nil
}

func (h *ChargingStationHandler) OnStatusNotification(chargePoint string, request *core.StatusNotificationRequest) (confirmation *core.StatusNotificationConfirmation, error error) {
	chargePointModel, err := repository.GetChargePointByID(context.Background(), chargePoint)
	if err != nil {
		fmt.Println(err)
		return core.NewStatusNotificationConfirmation(), err
	}
	if chargePointModel == nil {
		return core.NewStatusNotificationConfirmation(), nil
	}
	if len(chargePointModel.Connectors) == 0 {
		chargePointModel.Connectors = append(chargePointModel.Connectors, models.Connector{
			ConnectorID: request.ConnectorId,
			Status:      request.Status,
		})
		chargePointModel.Status = request.Status
		err = repository.UpdateChargePoint(context.Background(), chargePoint, chargePointModel)
		if err != nil {
			fmt.Println(err)
			return core.NewStatusNotificationConfirmation(), err
		}
	}
	cpConnector := slices.IndexFunc(chargePointModel.Connectors, func(connector models.Connector) bool {
		return connector.ConnectorID == request.ConnectorId
	})
	if cpConnector == -1 {
		chargePointModel.Connectors = append(chargePointModel.Connectors, models.Connector{
			ConnectorID: request.ConnectorId,
			Status:      request.Status,
		})
		chargePointModel.Status = request.Status
		err = repository.UpdateChargePoint(context.Background(), chargePoint, chargePointModel)
		if err != nil {
			fmt.Println(err)
			return core.NewStatusNotificationConfirmation(), err
		}
	}
	chargePointModel.Connectors[cpConnector].Status = request.Status
	err = repository.UpdateChargePoint(context.Background(), chargePoint, chargePointModel)
	if err != nil {
		fmt.Println(err)
		return core.NewStatusNotificationConfirmation(), err
	}
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
	chargePointModel, err := repository.GetChargePointByID(context.Background(), chargePoint)
	trnxID := int(time.Now().Unix())
	if err != nil {
		return core.NewStartTransactionConfirmation(&types.IdTagInfo{Status: types.AuthorizationStatusInvalid}, trnxID), err
	}
	if chargePointModel == nil {
		return core.NewStartTransactionConfirmation(&types.IdTagInfo{Status: types.AuthorizationStatusInvalid}, trnxID), nil
	}
	trnx := models.Transaction{
		TransactionID: trnxID,
		ChargePointID: chargePoint,
		ConnectorID:   request.ConnectorId,
		UserID:        request.IdTag,
		StartTime:     time.Now(),
		MeterStart:    request.MeterStart,
		Status:        "START",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = repository.CreateTransaction(context.Background(), &trnx)
	if err != nil {
		return core.NewStartTransactionConfirmation(&types.IdTagInfo{Status: types.AuthorizationStatusAccepted}, trnxID), err
	}
	return core.NewStartTransactionConfirmation(&types.IdTagInfo{Status: types.AuthorizationStatusAccepted}, trnxID), nil
}

func (h *ChargingStationHandler) OnStopTransaction(chargePoint string, request *core.StopTransactionRequest) (confirmation *core.StopTransactionConfirmation, error error) {
	trnx, err := repository.GetTransactionByID(context.Background(), request.TransactionId)
	if err != nil {
		return core.NewStopTransactionConfirmation(), err
	}
	if trnx == nil {
		return core.NewStopTransactionConfirmation(), nil
	}
	trnx.StopTime = time.Now()
	trnx.Status = "STOP"
	trnx.MeterStop = request.MeterStop
	trnx.UpdatedAt = time.Now()
	err = repository.UpdateTransaction(context.Background(), trnx.TransactionID, trnx)
	if err != nil {
		return core.NewStopTransactionConfirmation(), err
	}
	return core.NewStopTransactionConfirmation(), nil
}
