package users

import (
	"github.com/gofiber/fiber/v2"
)

func getByUser(c *fiber.Ctx) error {
	//var user = new(entity.User)
	return c.JSON("getByUser")
}

func getAllUsers(c *fiber.Ctx) error {
	//var users []entity.User
	return c.JSON("getAllUsers")
}
