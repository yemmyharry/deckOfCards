package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {

	api := app.Group("/api")
	{
		api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello World!")
		})
	}
}
