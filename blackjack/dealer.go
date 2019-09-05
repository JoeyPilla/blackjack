package blackjack

import (
	"fmt"

	"./card"
)

func (h hand) dealerString() string {
	return fmt.Sprintf("%15s, %15s", "***HIDDEN***", h[1].String())
}

type dealerPlayer struct {
	hand hand
}

func (ai *dealerPlayer) SetBalance(payout int) {
	//
}

func (ai *dealerPlayer) GetBalance() int {
	return 0
}

func (d dealerPlayer) GetBet() int {
	// here to satisfy Player interface
	return 1
}

func (d dealerPlayer) SetBet() {
	// here to satisfy Player interface
}

func (ai *dealerPlayer) AddToHand(card card.Card, h int) {
	ai.hand = append(ai.hand, card)
}

func (ai *dealerPlayer) GetHand() []hand {
	return []hand{ai.hand}
}

func (ai *dealerPlayer) NewHand() {
	ai.hand = hand{}
}

func createDealer() *dealerPlayer {
	return &dealerPlayer{
		hand: hand{},
	}
}

func (ai *dealerPlayer) Split() {

}
func (d dealerPlayer) Play(hand hand, dealer card.Card) Move {
	dScore := hand.score()
	if dScore <= 16 || (dScore == 17 && hand.minScore() != 17) {
		return MoveHit
	}
	return MoveStand
}

func (ai *dealerPlayer) DoubleDown() {
	// nil
}

func (d dealerPlayer) Results(dealer []hand) {
	// here to satisfy Player interface
}
