package users

import (
	"Belobetty-Starter/internal/git/entity"
	"Belobetty-Starter/internal/use_cases"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func getByUser(c *fiber.Ctx) error {
	userJson, err := use_cases.GetOneSynced("git", strings.ToLower(c.Params("user")), c.Get("token"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	user := new(entity.User)
	err = json.NewDecoder(userJson.Body).Decode(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	defer userJson.Body.Close()
	if user == nil {
		return c.Status(http.StatusNoContent).JSON("")
	}
	return c.JSON(user)
}

func getAllUsers(c *fiber.Ctx) error {
	usersJson, err := use_cases.GetAllSynced("git", c.Get("token"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	var users []entity.User
	err = json.NewDecoder(usersJson.Body).Decode(&users)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	defer usersJson.Body.Close()
	if len(users) == 0 {
		return c.Status(http.StatusNoContent).JSON("")
	}
	return c.JSON(users)
}
