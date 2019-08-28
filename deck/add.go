package deck

import (
	"../card"
)

func (d *Deck) AddJokers(amount int) {
	for ; amount > 0; amount-- {
		d.Deck = append(d.Deck, card.Card{
			Suit: card.Joker,
		})
	}
}

func (d *Deck) AddDecks(amount int) {
	for ; amount > 0; amount-- {
		d.Deck = append(d.Deck, NewDeck().Deck...)
	}
}
