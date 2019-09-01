package deck

import "../card"

func (d *Deck) Draw(n int) []card.Card {
	if d.Len() < n {
		n = d.Len()
	}
	ret := make([]card.Card, n)
	for i := 0; i < n; i++ {
		ret[i] = d.Deck[i]
	}
	d.Deck = d.Deck[n:]
	return ret
}
