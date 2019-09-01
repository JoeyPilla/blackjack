package blackjack

import (
	"fmt"

	"./card"
	"./deck"
)

type Game struct {
	deck          deck.Deck
	stage         int
	players       []hand
	dealer        hand
	currentPlayer int
}

const (
	playerTurn int = iota
	dealerTurn
	handOver
)

func StartGame(numberOfPlayers, decks int) Game {
	game := Game{}
	game.deck = deck.NewDeck()
	game.deck.AddDecks(decks - 1)
	game.deck.Shuffle()
	for i := 0; i < numberOfPlayers; i++ {
		game.players = append(game.players, nil)
	}
	game.dealer = nil
	return game
}

func (game *Game) Deal() {
	d := game.deck
	var c card.Card
	for i := 0; i < 2; i++ {
		for j := range game.players {
			c, d = draw(d)
			game.players[j] = append(game.players[j], c)
		}
		c, d = draw(d)
		game.dealer = append(game.dealer, c)
	}
	game.stage = playerTurn
	game.currentPlayer = 0
	game.deck = d
}

func (game *Game) EndGame() {
	fmt.Println("==FINAL HANDS==")
	dScore := game.dealer.score()
	for i, player := range game.players {
		fmt.Printf("%-8s: %2d %s\n", fmt.Sprintf("Player %d", i+1), player.score(), results(player.score(), dScore))
		game.players[i] = nil
	}
	fmt.Printf("%-8s: %d\n", "Dealer", dScore)
	fmt.Println("=======================================")
	game.dealer = nil
}

func (game Game) playerCount() int {
	return len(game.players)
}

func (game *Game) hit() {
	if game.stage == dealerTurn {
		card, deck := draw(game.deck)
		game.dealer = append(game.dealer, card)
		game.deck = deck
		return
	}
	p := game.currentPlayer
	card, deck := draw(game.deck)
	game.players[p] = append(game.players[p], card)
	game.deck = deck
}

func (game *Game) stand() {
	if game.currentPlayer < game.playerCount()-1 {
		game.currentPlayer++
	} else {
		game.stage = dealerTurn
	}
}

func draw(deck deck.Deck) (card.Card, deck.Deck) {
	card := deck.Draw(1)
	return card[0], deck
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
