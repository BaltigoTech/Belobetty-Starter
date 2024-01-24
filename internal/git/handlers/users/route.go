package users

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {

	user := r.Group("/user")

	user.Post("/", createUser)
	user.Get("/", getAllUsers)
	user.Get("/:user", getByUser)
	user.Put("/:user", updateUser)
	user.Delete("/:user", deleteUser)

}
