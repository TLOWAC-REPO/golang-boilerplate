package routes

import (
	"golang-boilerplate/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(router fiber.Router) {
	// url path - api/v1/user
	user := router.Group("/user")

	user.Get("/", controllers.GetAllUsers)
	user.Post("/", controllers.CreateUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}
