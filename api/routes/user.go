package routes

import "github.com/gofiber/fiber/v2"

func userRoute(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"success": true})
	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"success": true})
	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"success": true})
	})
}
