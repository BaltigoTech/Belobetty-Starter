package users

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func createUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("createUser")
}
