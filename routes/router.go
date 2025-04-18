package routes

import (
	"go-delivery-api/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/calculate", controller.CalculateCost)
}
