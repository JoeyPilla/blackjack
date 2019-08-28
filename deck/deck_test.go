package deck

import (
	"testing"

	"../card"
)

func TestShuffle(t *testing.T) {
	d := NewDeck()
	shuffledD := NewDeck()
	shuffledD.Shuffle()
	diff := false
	for i, deckCard := range d.Deck {
		if deckCard != shuffledD.Deck[i] {
			diff = true
		}
	}
	if !diff {
		t.Errorf("Deck not shuffled")
	}
}

func TestShuffleLength(t *testing.T) {
	d := NewDeck()
	d.Shuffle()
	// 13 ranks * 4 suits
	if len(d.Deck) != 13*4 {
		t.Errorf("Want %d, Have %d.", len(d.Deck), 13*4)
	}
}

func TestInit(t *testing.T) {
	d := NewDeck()
	// 13 ranks * 4 suits
	if len(d.Deck) != 13*4 {
		t.Errorf("Want %d, Have %d.", len(d.Deck), 13*4)
	}
}

func TestSort(t *testing.T) {
	SortedDeck := NewDeck()
	testDeck := NewDeck()
	SortedDeck.Shuffle()
	SortedDeck.Sort()

	diff := false
	for i, deckCard := range testDeck.Deck {
		if deckCard != SortedDeck.Deck[i] {
			diff = true
		}
	}
	if diff {
		t.Errorf("Deck not sorted")
	}
}

func TestFilterOneRank(t *testing.T) {
	deck := NewDeck()
	deck.Filter(card.Ace)
	hasAce := false
	for _, deckCard := range deck.Deck {
		if deckCard.Rank == card.Ace {
			hasAce = true
		}
	}
	if hasAce {
		t.Errorf("Ace not filtered out of deck")
	}
}

func TestFilterAllRanks(t *testing.T) {
	deck := NewDeck()
	deck.Filter(card.Ace)
	deck.Filter(card.Two)
	deck.Filter(card.Three)
	deck.Filter(card.Four)
	deck.Filter(card.Five)
	deck.Filter(card.Six)
	deck.Filter(card.Seven)
	deck.Filter(card.Eight)
	deck.Filter(card.Nine)
	deck.Filter(card.Ten)
	deck.Filter(card.Jack)
	deck.Filter(card.Queen)
	deck.Filter(card.King)
	if deck.Len() > 0 {
		t.Errorf("Deck still has cards after all ranks filtered.")
	}
}

func TestAddOneJoker(t *testing.T) {
	deck := NewDeck()
	deck.AddJokers(1)
	hasJoker := false
	for _, deckCard := range deck.Deck {
		if deckCard.Suit == card.Joker {
			hasJoker = true
			break
		}
	}
	if !hasJoker {
		t.Errorf("No Joker Added")
	}
}

func TestAddManyJokers(t *testing.T) {
	deck := NewDeck()
	amount := 5
	deck.AddJokers(amount)
	jokerCount := 0
	for _, deckCard := range deck.Deck {
		if deckCard.Suit == card.Joker {
			jokerCount++
		}
	}
	if jokerCount != amount {
		t.Errorf("Incorrect amount of Jokers. Have %d, Want %d.", jokerCount, amount)
	}
}

func TestAddOneDeck(t *testing.T) {
	deck := NewDeck()
	deck.AddDecks(1)
	if deck.Len() != 2*13*4 {
		t.Errorf("Incorrect card count after deck added. Have %d, want %d.", deck.Len(), 2*13*4)
	}
}

func TestAddManyDecks(t *testing.T) {
	deck := NewDeck()
	deck.AddDecks(5)
	if deck.Len() != 6*13*4 {
		t.Errorf("Incorrect card count after deck added. Have %d, want %d.", deck.Len(), 6*13*4)
	}
}

func TestShuffleMultiDeck(t *testing.T) {
	d := NewDeck()
	d.AddDecks(1)
	shuffledD := NewDeck()
	shuffledD.AddDecks(1)
	shuffledD.Shuffle()
	diff := false
	for i, deckCard := range d.Deck {
		if deckCard != shuffledD.Deck[i] {
			diff = true
		}
	}
	if !diff {
		t.Errorf("Multiple decks not shuffled")
	}
}

func TestSortMultiDeck(t *testing.T) {
	SortedDeck := NewDeck()
	SortedDeck.AddDecks(1)
	testDeck := NewDeck()
	SortedDeck.Shuffle()
	SortedDeck.Sort()

	diff := false
	for i, deckCard := range testDeck.Deck {
		for j := 0; j < 2; j++ {
			if deckCard != SortedDeck.Deck[j+2*i] {
				diff = true
			}
		}
	}
	if diff {
		t.Errorf("Multi deck not sorted")
	}
}

func TestFilterOneRankMultiDeck(t *testing.T) {
	deck := NewDeck()
	deck.AddDecks(3)
	deck.Filter(card.Ace)
	hasAce := false
	for _, deckCard := range deck.Deck {
		if deckCard.Rank == card.Ace {
			hasAce = true
		}
	}
	if hasAce {
		t.Errorf("Ace not filtered out of deck")
	}
}
func TestFilterAllRanksMultiDeck(t *testing.T) {
	deck := NewDeck()
	deck.AddDecks(3)
	deck.Filter(card.Ace)
	deck.Filter(card.Two)
	deck.Filter(card.Three)
	deck.Filter(card.Four)
	deck.Filter(card.Five)
	deck.Filter(card.Six)
	deck.Filter(card.Seven)
	deck.Filter(card.Eight)
	deck.Filter(card.Nine)
	deck.Filter(card.Ten)
	deck.Filter(card.Jack)
	deck.Filter(card.Queen)
	deck.Filter(card.King)
	if deck.Len() > 0 {
		t.Errorf("Deck still has cards after all ranks filtered.")
	}
}
