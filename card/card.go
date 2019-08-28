package card

import (
	"fmt"
)

type Deck struct {
	cards []Card
}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func NewDeck() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := min; rank <= max; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}

func (c Card) Value() int {
	return int(c.Suit)*13 + int(c.Rank)
}
