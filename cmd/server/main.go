package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok start")
	})

	app.Listen(":9001")
}
