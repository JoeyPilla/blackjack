package card

import "strconv"

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var rankString = []string{
	"",
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King"}

const (
	min = Ace
	max = King
)

func (r Rank) String() string {
	if r > Rank(len(rankString)-1) {
		return "Rank(" + strconv.FormatInt(int64(r), 10) + ")"
	}
	return rankString[r]
}
