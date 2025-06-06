package controller

import (
	"go-delivery-api/model"
	"go-delivery-api/service"

	"github.com/gofiber/fiber/v2"
)

func CalculateCost(c *fiber.Ctx) error {
	var order model.OrderRequest
	if err := c.BodyParser(&order); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	cost := service.CalculateMinimumCost(order.Products)
	return c.JSON(fiber.Map{"minimum_cost": cost})
}
