package blackjack

import (
	"./card"
)

type Player interface {
	SetBet(int)
	GetBet() int
	DoubleDown()
	SetBalance(int)
	GetBalance() int
	Split()
	AddToHand(card.Card, int)
	GetHand() []hand
	NewHand()
	Play(hand, card.Card) Move
	Results([]hand)
}

type aiPlayer struct {
	bet     int
	balance int
	hand    []hand
}

func (ai *aiPlayer) SetBet(count int) {
	bet := 10
	if count >= 10 {
		bet = 50
	}
	if count >= 5 {
		bet = 20
	}
	ai.bet = bet
}

func (ai *aiPlayer) Split() {
	h := ai.hand
	ai.hand = []hand{
		hand{
			h[0][0],
		},
		hand{
			h[0][1],
		},
	}
}

func (ai *aiPlayer) DoubleDown() {
	ai.bet = ai.bet * 2
}

func (ai *aiPlayer) GetBet() int {
	return ai.bet
}

func (ai *aiPlayer) SetBalance(payout int) {
	ai.balance += payout
}

func (ai *aiPlayer) GetBalance() int {
	return ai.balance
}

func (ai *aiPlayer) AddToHand(card card.Card, h int) {
	ai.hand[h] = append(ai.hand[h], card)
}

func (ai *aiPlayer) GetHand() []hand {
	return ai.hand
}

func (ai *aiPlayer) NewHand() {
	ai.hand = []hand{hand{}}
}

func createAIPlayer() *aiPlayer {
	return &aiPlayer{
		bet:     0,
		balance: 0,
		hand:    []hand{hand{}},
	}
}

func (ai aiPlayer) Play(hand hand, dealer card.Card) Move {
	if len(hand) == 2 {
		if hand[0] == hand[1] {
			score := hand[0].BlackjackValue()
			if score >= 8 && score != 10 {
				return MoveSplit
			}
		}
		if (hand.score() == 10 || hand.score() == 11) && hand.score() == hand.minScore() {
			return MoveDouble
		}
		dealerScore := dealer.BlackjackValue()
		if dealerScore >= 5 || dealerScore <= 6 {
			return MoveStand
		}
		if hand.score() < 13 {
			return MoveHit
		}
	}

	return MoveStand
}

func (ai aiPlayer) Results(dealer []hand) {
	// TODO(JoeyPilla): implement
}
