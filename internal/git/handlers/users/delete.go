package users

import (
	"Belobetty-Starter/internal/git/entity"
	"Belobetty-Starter/internal/use_cases"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

const actionDelete = "DeleteUser"

func deleteUser(c *fiber.Ctx) error {
	userJson, err := use_cases.GetOneSynced("git", strings.ToLower(c.Params("user")), c.Get("token"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	user := new(entity.User)
	err = json.NewDecoder(userJson.Body).Decode(user)

	sender, err := use_cases.NewSenderQueue("git")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = sender.Exec(user, c.Get("token", "defaultValue"), actionDelete)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado remoção do usuario %s.", user.UserName))
}
