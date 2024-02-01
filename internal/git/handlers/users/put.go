package users

import (
	"Belobetty-Starter/internal/git/entity"
	"Belobetty-Starter/internal/use_cases"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const actionPut = "UpdateUser"

func updateUser(c *fiber.Ctx) error {

	var user = new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	sender, err := use_cases.NewSenderQueue("git")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = sender.Exec(user, c.Get("token", "defaultValue"), actionPut)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("Solicitado atualização do usuario %s.", user.UserName))
}
