package routes

import (
	"deckOfCards/controllers"
	"deckOfCards/models"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	deck := controllers.DeckController{DeckService: make(map[string]*models.Deck)}
	api := app.Group("/api")
	{
		api.Get("/", deck.CreateDeck)
		api.Get("/:deck_id", deck.OpenDeck)
		api.Get("/:deck_id/draw", deck.DrawCard)
	}
}
