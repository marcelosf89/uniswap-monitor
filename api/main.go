package main

import (
	"context"
	"os"
	"os/signal"

	"brahmafi-build-it/api/pkg/api"
	"brahmafi-build-it/api/pkg/configurations"
	"brahmafi-build-it/api/pkg/database"
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
