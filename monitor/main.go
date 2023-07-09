package main

import (
	"context"
	"os"
	"os/signal"

	"brahmafi-build-it/monitor/pkg/configurations"
	"brahmafi-build-it/monitor/pkg/database"
	"brahmafi-build-it/monitor/pkg/models"
	"brahmafi-build-it/monitor/pkg/services"
)

func main() {
	var ctx, cancel = context.WithCancel(context.Background())

	database.Init(configurations.GetDatabaseConnectionString(), database.NewMongoDb(), ctx)
	database.OpenConnection()

	defer database.CloseConnection()
	defer cancel()

	createDummyMonitor()

	services.InitEthClient(services.NewGoEthClient())
	services.StartEthClient()

	monitorsStates := database.GetAllMonitorStates()
	for _, state := range monitorsStates {
		services.StartMonitor(services.NewUniswapV3Monitor(), &state, ctx)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}

func createDummyMonitor() {
	monitorsStates := database.GetAllMonitorStates()

	if len(monitorsStates) <= 0 {
		state := models.MonitorState{
			Id:            "1",
			Address:       "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
			LastBlock:     17651601,
			Token0Balance: "0",
			Token1Balance: "0",
			Enabled:       true,
		}
		database.SaveMonitorState(state)
	}
}
