package users

import (
	"Belobetty-Starter/internal/git/entity"
	"Belobetty-Starter/internal/use_cases"
	"fmt"
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

	uc, err := use_cases.NewSenderQueue("git")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = uc.Exec(user, c.Get("token", "defaultValue"), actionPost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado criação do usuario %s.", user.UserName))
}
