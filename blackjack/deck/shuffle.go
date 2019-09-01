package deck

import (
	"math/rand"
	"time"

	"../card"
)

//shuffles current deck
func (d *Deck) Shuffle() {
	ret := make([]card.Card, len(d.Deck))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	newOrder := r.Perm(len(d.Deck))
	for i, rand := range newOrder {
		ret[i] = d.Deck[rand]
	}
	d.Deck = ret
}
