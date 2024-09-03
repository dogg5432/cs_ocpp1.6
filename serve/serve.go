package serve

import (
	"fmt"

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
	// handler := &CentralSystemHandler{}
	// centralSystem.SetCoreHandler(handler)

	// Start central system
	listenPort := 8887
	fmt.Printf("starting central system")
	centralSystem.Start(listenPort, "/{ws}") // This call starts server in daemon mode and is blocking
	fmt.Println("stopped central system")
}
