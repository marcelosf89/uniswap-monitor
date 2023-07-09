package router

import (
	controllersv1 "uniswap-monitor/api/pkg/api/v1/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func AddPoolRoutesV1(fiberApp *fiber.App) {
	var gv1 = fiberApp.Group("v1/api")

	gv1.Get("/pool/:pool_id", controllersv1.Get)
	gv1.Get("/pool/:pool_id/historic", controllersv1.GetHistoric)
}
