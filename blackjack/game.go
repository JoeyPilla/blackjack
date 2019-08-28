package blackjack

import (
	"fmt"

	"../card"
	"../deck"
)

type Game struct {
	Deck          deck.Deck
	Stage         int
	Players       []Hand
	Dealer        Hand
	CurrentPlayer int
}

const (
	PlayerTurn int = iota
	DealerTurn
	HandOver
)

func (game Game) PlayerCount() int {
	return len(game.Players)
}

func (game *Game) Deal() {
	d := game.Deck
	var c card.Card
	for i := 0; i < 2; i++ {
		for j := range game.Players {
			c, d = Draw(d)
			game.Players[j] = append(game.Players[j], c)
		}
		c, d = Draw(d)
		game.Dealer = append(game.Dealer, c)
	}
	game.Stage = PlayerTurn
	game.CurrentPlayer = 0
	game.Deck = d
}

func (game *Game) Hit() {
	if game.Stage == DealerTurn {
		card, deck := Draw(game.Deck)
		game.Dealer = append(game.Dealer, card)
		game.Deck = deck
		return
	}
	p := game.CurrentPlayer
	card, deck := Draw(game.Deck)
	game.Players[p] = append(game.Players[p], card)
	game.Deck = deck
}

func (game *Game) Stand() {
	if game.CurrentPlayer < game.PlayerCount()-1 {
		game.CurrentPlayer++
	} else {
		game.Stage = DealerTurn
	}
}

func Draw(deck deck.Deck) (card.Card, deck.Deck) {
	card := deck.Draw(1)
	return card[0], deck
}

func StartGame(numberOfPlayers, Decks int) Game {
	game := Game{}
	game.Deck = deck.NewDeck()
	game.Deck.AddDecks(Decks - 1)
	game.Deck.Shuffle()
	for i := 0; i < numberOfPlayers; i++ {
		game.Players = append(game.Players, nil)
	}
	game.Dealer = nil
	return game
}

func (game *Game) EndGame() {
	fmt.Println("==FINAL HANDS==")
	dScore := game.Dealer.Score()
	for i, player := range game.Players {
		fmt.Printf("%-8s: %2d %s\n", fmt.Sprintf("Player %d", i+1), player.Score(), results(player.Score(), dScore))
		game.Players[i] = nil
	}
	fmt.Printf("%-8s: %d\n", "Dealer", dScore)
	fmt.Println("=======================================")
	game.Dealer = nil
}

func results(player, dealer int) string {
	ret := ""
	switch {
	case player > 21:
		ret = fmt.Sprintf("you busted.")
	case dealer > 21:
		ret = fmt.Sprintf("dealer busted, you win!")
	case player > dealer:
		ret = fmt.Sprintf("you win!")
	case dealer > player:
		ret = fmt.Sprintf("you lose.")
	case dealer == player:
		ret = fmt.Sprintf("you draw.")
	}
	return ret
}
