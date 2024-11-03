package routes

import (
	"golang-boilerplate/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(router fiber.Router) {
	// url path - api/v1/user
	userRoute := router.Group("/user")

	userCtrl := controllers.NewUserCtrl()

	userRoute.Get("/", userCtrl.GetAllUsers)
	userRoute.Post("/", userCtrl.CreateUser)
	userRoute.Patch("/:id", userCtrl.UpdateUser)
	userRoute.Delete("/:id", userCtrl.DeleteUser)
}
