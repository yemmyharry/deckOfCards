package controllers

import (
	"deckOfCards/data"
	"deckOfCards/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

type DeckController struct {
	DeckService map[string]*models.Deck
}

func (d *DeckController) CreateDeck(c *fiber.Ctx) error {

	var deck = models.Deck{
		ID:        uuid.New().String(),
		Shuffled:  false,
		Remaining: len(data.Data),
		Cards:     data.Data,
	}
	d.DeckService[deck.ID] = &deck

	res := struct {
		ID        string `json:"id"`
		Shuffled  bool   `json:"shuffled"`
		Remaining int    `json:"remaining"`
	}{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
	return c.JSON(res)
}

func (d *DeckController) OpenDeck(c *fiber.Ctx) error {
	id := c.Params("deck_id")
	deck, ok := d.DeckService[id]
	if !ok {
		return c.Status(404).JSON(fiber.Map{"message": "Deck not found"})
	}
	deck.Remaining = len(deck.Cards)
	return c.JSON(deck)
}

func (d *DeckController) DrawCard(c *fiber.Ctx) error {
	id := c.Params("deck_id")
	deck, ok := d.DeckService[id]
	if !ok {
		return c.Status(404).JSON(fiber.Map{"message": "Deck not found"})
	}

	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		count = 1
	}
	var randomCards []models.Card
	for i := 0; i < count; i++ {
		randIndex := rand.Intn(len(deck.Cards))
		randomCards = append(randomCards, *deck.Cards[randIndex])
		deck.Cards = append(deck.Cards[:randIndex], deck.Cards[randIndex+1:]...)
	}

	return c.JSON(fiber.Map{"cards": randomCards})
}
