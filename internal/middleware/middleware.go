package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupMiddleware(app *fiber.App) *fiber.App {

	// http logging middleware
	app.Use(requestid.New())
	app.Use(HttpLoggingMiddleware())

	// cors middleware
	app.Use(CorsMiddleware())

	// panic recovery middleware
	app.Use(recover.New())

	return app
}
