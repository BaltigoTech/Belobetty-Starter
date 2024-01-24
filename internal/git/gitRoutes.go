package git

import (
	"Belobetty-Starter/internal/git/handlers/repositories"
	"Belobetty-Starter/internal/git/handlers/users"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router) {

	git := r.Group("/github/")

	users.SetRoutes(git)
	repositories.SetRoutes(git)

}
