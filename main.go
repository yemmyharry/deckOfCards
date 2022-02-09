package main

import (
	"deckOfCards/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
