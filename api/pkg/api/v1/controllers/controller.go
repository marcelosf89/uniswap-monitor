package controllers

import (
	services "brahmafi-build-it/api/pkg/services/v1"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	blockQueryString := c.Query("block")
	blocks := strings.Split(blockQueryString, ",")

	response, _ := services.HandleGetPoolBalanceByBlocks(c.Params("pool_id"), blocks)

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetHistoric(c *fiber.Ctx) error {
	response, _ := services.HandleGetPoolHistoric(c.Params("pool_id"))

	return c.Status(fiber.StatusOK).JSON(response)
}
