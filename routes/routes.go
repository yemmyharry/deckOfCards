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
		api.Get("/create", deck.CreateDeck)
		api.Get("/open/:deck_id", deck.OpenDeck)
		api.Get("/draw/:deck_id", deck.DrawCard)
	}
}
