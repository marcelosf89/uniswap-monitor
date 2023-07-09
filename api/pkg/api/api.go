package api

import (
	"encoding/json"
	"fmt"
	"runtime"

	fiber "github.com/gofiber/fiber/v2"

	routermgn "brahmafi-build-it/api/pkg/api/mgn/router"
	routerv1 "brahmafi-build-it/api/pkg/api/v1/router"
	"brahmafi-build-it/api/pkg/configurations"
)

var fiberApp *fiber.App

func Initialize() {
	fiberApp = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	setupApiRouter()

	go func() {
		// Run server.
		if err := fiberApp.Listen(fmt.Sprintf(":%v", configurations.GetPort())); err != nil {
		}
	}()
}

func Shoutdown() {
	runtime.GC()
}

func setupApiRouter() {
	routerv1.AddPoolRoutesV1(fiberApp)
	routermgn.AddApiHealthCheck(fiberApp)
}
