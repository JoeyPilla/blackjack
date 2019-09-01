package card

import (
	"testing"
)

var cardTests = []struct {
	c    Card
	want string
}{
	{
		Card{Rank: Ace, Suit: Heart},
		"Ace of Hearts",
	},
	{
		Card{Rank: Two, Suit: Spade},
		"Two of Spades",
	},
	{
		Card{Rank: Nine, Suit: Diamond},
		"Nine of Diamonds",
	},
	{
		Card{Rank: Jack, Suit: Club},
		"Jack of Clubs",
	},
	{
		Card{Suit: Joker},
		"Joker",
	},
}

func TestCard(t *testing.T) {
	for i, test := range cardTests {
		if test.c.String() != test.want {
			t.Errorf("Failed Test %d. Want %s, Have %s.", i, test.want, test.c.String())
		}
	}
}

func TestNew(t *testing.T) {
	cards := NewDeck()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("Want %d, Have %d.", len(cards), 13*4)
	}
}
