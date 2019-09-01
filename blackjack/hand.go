package blackjack

import (
	"fmt"
	"strings"

	"./card"
)

type hand []card.Card

func (h hand) String() string {
	cards := make([]string, len(h))
	for i := range h {
		cards[i] = fmt.Sprintf("%15s", h[i].String())
	}
	return strings.Join(cards, ", ")
}

func (h hand) score() int {
	ret := h.minScore()
	for _, c := range h {
		if c.Rank == card.Ace {
			if ret <= 11 {
				ret += 10
			}
		}
	}
	return ret
}

func (h hand) minScore() int {
	ret := 0
	for _, c := range h {
		ret += min(int(c.Rank), 10)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
