package repositories

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {

	repo := r.Group("/repository")

	repo.Post("/", createRepository)
	repo.Get("/", getAllRepository)
	repo.Get("/:id", getByRepository)
	repo.Put("/:id", updateRepository)
	repo.Delete("/:id", deleteRepository)

}
