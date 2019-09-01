package card

import "strconv"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // this is a special case
)

var suitsString = []string{"Spade", "Diamond", "Club", "Heart", "Joker"}
var suits = [...]Suit{Spade, Diamond, Club, Heart}

func (s Suit) String() string {
	if s > Suit(len(suitsString)-1) {
		return "Suit(" + strconv.FormatInt(int64(s), 10) + ")"
	}
	return suitsString[s]
}
