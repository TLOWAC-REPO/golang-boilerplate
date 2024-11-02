package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func SetupMiddleware(app *fiber.App) *fiber.App {

	// http logging middleware
	app.Use(HttpLoggingMiddleware())

	// cors middleware
	app.Use(CorsMiddleware())

	return app
}
