package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HttpLoggingMiddleware() fiber.Handler {
	// 00:41:02 | 500 |      19.334µs | 127.0.0.1 | POST | /api/post | -
	// app.Use(logger.New())

	// [01:01:21 | 127.0.0.1]:58681 |     298.708µs | 200 - POST /api/post
	// app.Use(logger.New(logger.Config{
	// 	Format: "[${time} | ${ip}]:${port} | ${latency} | ${status} - ${method} ${path}\n",
	// }))

	httpLoggingFormat := logger.Config{
		Format: "[${time} | ${ip}]:${port} | ${latency} | ${status} - ${method} ${path}\n",
	}

	return logger.New(httpLoggingFormat)
}
