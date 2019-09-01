package deck

import (
	"../card"
)

// a deck of cards
type Deck struct {
	Deck []card.Card
}

// get a new deck of cards
func NewDeck() Deck {
	return Deck{
		card.NewDeck(),
	}
}
