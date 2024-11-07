package serve

import (
	"fmt"

	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/handlers"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
)

func Run() {
	configApp := config.ConfigApp.Server
	centralSystem := ocpp16.NewCentralSystem(nil, nil)

	// Set callback handlers for connect/disconnect
	centralSystem.SetNewChargePointHandler(func(chargePointId ocpp16.ChargePointConnection) {
		fmt.Printf("new charge point %+v connected\n", chargePointId)
	})
	centralSystem.SetChargePointDisconnectedHandler(func(chargePointId ocpp16.ChargePointConnection) {
		fmt.Printf("charge point %v disconnected\n", chargePointId)
	})

	// Set handler for profile callbacks
	Charginghandler := &handlers.ChargingStationHandler{}
	centralSystem.SetCoreHandler(Charginghandler)

	// Start central system
	listenPort := configApp.Port
	fmt.Printf("starting central system url: ws://localhost:%d%s\n", listenPort, configApp.Path)
	centralSystem.Start(listenPort, "/ocpp16/{chargepoint}") // This call starts server in daemon mode and is blocking
	fmt.Println("stopped central system")
}
