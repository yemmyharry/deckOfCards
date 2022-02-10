package test

import (
	"deckOfCards/models"
	"deckOfCards/routes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var (
	app *fiber.App
	id  string
)

func TestMain(m *testing.M) {
	app = fiber.New()
	routes.Setup(app)
	m.Run()
}
func TestCreateDeck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/create", nil)
	if err != nil {
		log.Print(err)
	}
	resp, err := app.Test(req)
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	var deck models.Deck
	err = json.Unmarshal(body, &deck)
	if err != nil {
		log.Print(err)
	}
	id = deck.ID
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if deck.Remaining != 52 {
		t.Errorf("Expected 52, got %d", deck.Remaining)
	}
	if deck.Shuffled != false {
		t.Errorf("Expected false, got %t", deck.Shuffled)
	}
}

func TestOpenDeck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/open/%s", id), nil)
	if err != nil {
		log.Print(err)
	}
	resp, err := app.Test(req)
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	var deck models.Deck
	err = json.Unmarshal(body, &deck)
	if err != nil {
		log.Print(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if deck.Remaining != 52 {
		t.Errorf("Expected 52, got %d", deck.Remaining)
	}
	if deck.Shuffled != false {
		t.Errorf("Expected false, got %t", deck.Shuffled)
	}
	if deck.Cards[0].Value != "Ace" {
		t.Errorf("Expected Ace, got %s", deck.Cards[0].Value)
	}
	if deck.Cards[deck.Remaining-1].Code != "KC" {
		t.Errorf("Expected KC, got %s", deck.Cards[deck.Remaining-1].Code)
	}

}

func TestDrawCard(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/draw/%s", id), nil)
	t.Run("Draw 1 cards from deck", func(t *testing.T) {
		url := req.URL.Query()
		url.Set("count", "4")
		req.URL.RawQuery = url.Encode()

		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)
		cards := struct {
			Cards []models.Card `json:"cards"`
		}{}
		_ = json.Unmarshal(body, &cards)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		if len(cards.Cards) != 4 {
			t.Errorf("Expected 1, got %d", len(cards.Cards))
		}
	})

	t.Run("Draw 4 cards from deck", func(t *testing.T) {
		url := req.URL.Query()
		url.Set("count", "4")
		req.URL.RawQuery = url.Encode()

		resp, _ := app.Test(req)
		body, _ := ioutil.ReadAll(resp.Body)
		cards := struct {
			Cards []models.Card `json:"cards"`
		}{}
		_ = json.Unmarshal(body, &cards)
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		if len(cards.Cards) != 4 {
			t.Errorf("Expected 4, got %d", len(cards.Cards))
		}
	})
}
