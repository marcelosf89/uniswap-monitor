package main

import (
	"context"
	"os"
	"os/signal"

	"uniswap-monitor/api/pkg/api"
	"uniswap-monitor/api/pkg/configurations"
	"uniswap-monitor/api/pkg/database"
)

func main() {
	var ctx, cancel = context.WithCancel(context.Background())

	database.Init(configurations.GetDatabaseConnectionString(), database.NewMongoDb(), ctx)
	database.OpenConnection()

	api.Initialize()

	defer api.Shoutdown()
	defer database.CloseConnection()
	defer cancel()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}
