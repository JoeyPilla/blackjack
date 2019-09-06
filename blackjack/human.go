package blackjack

import (
	"fmt"

	"./card"
)

type humanPlayer struct {
	bet     int
	balance int
	hand    []hand
}

func (ai *humanPlayer) SetBet() {
	bet := 10
	// fmt.Println("What would you like to bet?")
	// fmt.Scanf("%d", &bet)
	ai.bet = bet
}

func (ai *humanPlayer) DoubleDown() {
	ai.bet = ai.bet * 2
}

func (ai *humanPlayer) Split() {
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

func (ai *humanPlayer) GetBet() int {
	return ai.bet
}
func (ai *humanPlayer) SetBalance(payout int) {
	ai.balance += payout
}

func (ai *humanPlayer) GetBalance() int {
	return ai.balance
}

func (ai *humanPlayer) AddToHand(card card.Card, hand int) {
	ai.hand[hand] = append(ai.hand[hand], card)
}

func (ai *humanPlayer) GetHand() []hand {
	return ai.hand
}

func (ai *humanPlayer) NewHand() {
	ai.hand = []hand{hand{}}
}
func createHumanPlayer() *humanPlayer {
	return &humanPlayer{
		bet:     0,
		balance: 0,
		hand:    []hand{hand{}},
	}
}

func (ai humanPlayer) Play(hand hand, dealer card.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand, or (d)ouble down")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "split":
			return MoveSplit
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai humanPlayer) Results(dealer []hand) {
	fmt.Println("==FINAL HANDS==")
	for _, h := range ai.hand {
		fmt.Println("Player:", h, h.score())
	}
	fmt.Println("Dealer:", dealer[0], dealer[0].score())
}
