package deck

import "sort"

// length of current deck
func (d Deck) Len() int {
	return len(d.Deck)
}

// swap deck values
func (d Deck) Swap(i, j int) {
	d.Deck[i], d.Deck[j] = d.Deck[j], d.Deck[i]
}

// default less function
func (d Deck) Less(i, j int) bool {
	return d.Deck[i].Value() < d.Deck[j].Value()
}

// default sort function
func (d *Deck) Sort() {
	sort.Sort(d)
}
