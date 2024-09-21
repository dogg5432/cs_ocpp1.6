package serve

import (
	"fmt"

	"github.com/dogg5432/central_charger/handlers"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
)

func Run() {
	centralSystem := ocpp16.NewCentralSystem(nil, nil)

	// Set callback handlers for connect/disconnect
	centralSystem.SetNewChargePointHandler(func(chargePointId ocpp16.ChargePointConnection) {
		fmt.Printf("new charge point %v connected", chargePointId)
	})
	centralSystem.SetChargePointDisconnectedHandler(func(chargePointId ocpp16.ChargePointConnection) {
		fmt.Printf("charge point %v disconnected", chargePointId)
	})

	// Set handler for profile callbacks
	Charginghandler := &handlers.ChargingStationHandler{}
	centralSystem.SetCoreHandler(Charginghandler)

	// Start central system
	listenPort := 8887
	fmt.Printf("starting central system")
	centralSystem.Start(listenPort, "/ocpp16") // This call starts server in daemon mode and is blocking
	fmt.Println("stopped central system")
}
