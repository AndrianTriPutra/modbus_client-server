package server

import (
	"log"
	"modbus_server/config"
	"os"

	"github.com/simonvetter/modbus"
)

// Example handler object, passed to the NewServer() constructor above.
type exampleHandler struct {
	// these are here to hold client-provided (written) values, for both coils and
	// holding registers
	coils       [100]bool
	holdingReg1 uint16
	holdingReg2 uint16

	// this is a 16-bit signed integer
	holdingReg3 int16
}

func ModbusServer() {
	var eh *exampleHandler

	// create the handler object
	eh = &exampleHandler{}

	// create the server object
	server, err := modbus.NewServer(&modbus.ServerConfiguration{
		// listen on url
		URL: config.ServerSetting.Url,
		// close idle connections after xs of inactivity
		Timeout: config.ServerSetting.Timeout,
		// accept max concurrent connections max.
		MaxClients: uint(config.ServerSetting.MaxClients),
	}, eh)
	if err != nil {
		log.Printf("failed to create server: %v\n", err)
		os.Exit(1)
	}

	// start accepting client connections
	// note that Start() returns as soon as the server is started
	err = server.Start()
	if err != nil {
		log.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("starting server")
	}

}
