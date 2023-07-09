package router

import (
	controllersmgn "brahmafi-build-it/api/pkg/api/mgn/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func AddApiHealthCheck(fiberApp *fiber.App) {
	var gv1 = fiberApp.Group("")

	gv1.Get("/health", controllersmgn.GeHealth)
}
