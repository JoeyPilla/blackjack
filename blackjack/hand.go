package blackjack

import (
	"fmt"
	"strings"

	"../card"
)

type Hand []card.Card

func (h Hand) Score() int {
	ret := h.MinScore()
	for _, c := range h {
		if c.Rank == card.Ace {
			if ret <= 11 {
				ret += 10
			}
		}
	}
	return ret
}

func (h Hand) MinScore() int {
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

func (h Hand) String() string {
	cards := make([]string, len(h))
	for i := range h {
		cards[i] = fmt.Sprintf("%15s", h[i].String())
	}
	return strings.Join(cards, ", ")
}
