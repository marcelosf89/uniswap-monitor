package controllers

import (
	services "brahmafi-build-it/api/pkg/services/mgn"

	fiber "github.com/gofiber/fiber/v2"
)

func GeHealth(c *fiber.Ctx) error {

	response, err := services.HandleHealthCheck()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
