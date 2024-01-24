package repositories

import (
	"github.com/gofiber/fiber/v2"
)

func getByRepository(c *fiber.Ctx) error {
	//var user = new(entity.User)
	return c.JSON("getByUser")
}

func getAllRepository(c *fiber.Ctx) error {
	//var users []entity.User
	return c.JSON("getAllUsers")
}
