package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// set /api path
	api := app.Group("/api")

	// set /api/v1 path
	v1 := api.Group("/v1")

	// setup user routing path
	RegisterUserRoute(v1)
}
