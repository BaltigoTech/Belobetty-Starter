package handlers

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {

	git := r.Group("/github/")

	repo := git.Group("/repository")

	repo.Post("/", nil)
	repo.Get("/", nil)
	repo.Get("/:id", nil)
	repo.Put("/:id", nil)
	repo.Delete("/:id", nil)

	user := git.Group("/user")

	user.Post("/", nil)
	user.Get("/", nil)
	user.Get("/:id", nil)
	user.Put("/:id", nil)
	user.Delete("/:id", nil)

}
