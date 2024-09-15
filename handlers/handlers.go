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
