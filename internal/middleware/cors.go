package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept",
	})
}
