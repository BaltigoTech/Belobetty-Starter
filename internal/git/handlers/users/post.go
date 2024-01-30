package users

import (
	"Belobetty-Starter/internal/git/dto"
	"Belobetty-Starter/internal/git/entity"
	"Belobetty-Starter/internal/git/use_cases"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const actionPost = "CreateUser"

func createUser(c *fiber.Ctx) error {

	var user = new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	err = user.Validate()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	uc, err := use_cases.NewSenderQueue("git")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	userDto := dto.NewMessageOut(user, c.Get("token", "defaultValue"), actionPost)

	err = uc.Exec(userDto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON("createUser")
}
