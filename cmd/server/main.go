package main

import (
	"Belobetty-Starter/internal/git"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	git.SetRoutes(v1)

	app.Listen(":9001")
}
