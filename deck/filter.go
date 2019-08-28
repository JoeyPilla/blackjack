package deck

import "../card"

// filter a card out of the deck
func (d *Deck) Filter(rank card.Rank) {
	cards := []card.Card{}
	for _, card := range d.Deck {
		if card.Rank != rank {
			cards = append(cards, card)
		}
	}
	d.Deck = cards
}
