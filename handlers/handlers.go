package handlers

import (
	"time"

	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

const defaultHeartbeatInterval = 600

type ChargingStationHandler struct{}

func (h *ChargingStationHandler) OnBootNotification(chargePoint ocpp16.ChargePoint, request *core.BootNotificationRequest) (confirmation *core.BootNotificationConfirmation, error error) {

	return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), defaultHeartbeatInterval, core.RegistrationStatusAccepted), nil
}

func (h *ChargingStationHandler) OnMeterValues(chargePoint ocpp16.ChargePoint, request *core.MeterValuesRequest) (confirmation *core.MeterValuesConfirmation, error error) {
	return core.NewMeterValuesConfirmation(), nil
}

func (h *ChargingStationHandler) OnAuthorize(chargePoint ocpp16.ChargePoint, request *core.AuthorizeRequest) (confirmation *core.AuthorizeConfirmation, error error) {
	return core.NewAuthorizationConfirmation(&types.IdTagInfo{}), nil
}

func (h *ChargingStationHandler) OnStatusNotification(chargePoint ocpp16.ChargePoint, request *core.StatusNotificationRequest) (confirmation *core.StatusNotificationConfirmation, error error) {
	return core.NewStatusNotificationConfirmation(), nil
}

func (h *ChargingStationHandler) OnHeartbeat(chargePoint ocpp16.ChargePoint, request *core.HeartbeatRequest) (confirmation *core.HeartbeatConfirmation, error error) {
	return core.NewHeartbeatConfirmation(types.NewDateTime(time.Now())), nil
}


