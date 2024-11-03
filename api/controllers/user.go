package controllers

import (
	"golang-boilerplate/api/services"

	"github.com/gofiber/fiber/v2"
)

type UserCtrl interface {
	GetAllUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserCtrlImpl struct {
	userService services.UserService
}

func NewUserCtrl() UserCtrl {
	return &UserCtrlImpl{services.NewUserService()}
}

func (u *UserCtrlImpl) GetAllUsers(c *fiber.Ctx) error {
	userList, err := u.userService.GetAllUsers()
	if err != nil {
		return nil
	}

	return c.JSON(fiber.Map{"data": userList})
}
func (u *UserCtrlImpl) CreateUser(c *fiber.Ctx) error { return c.JSON(fiber.Map{"success": true}) }
func (u *UserCtrlImpl) UpdateUser(c *fiber.Ctx) error { return c.JSON(fiber.Map{"success": true}) }
func (u *UserCtrlImpl) DeleteUser(c *fiber.Ctx) error { return c.JSON(fiber.Map{"success": true}) }
