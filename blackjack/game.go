package blackjack

import (
	"errors"
	"fmt"

	"./card"
	"./deck"
)

// Game is the structure of the blackjack game.
type Game struct {
	numberOfHands int
	numberOfDecks int
	playerCount   int

	deck          deck.Deck
	stage         int
	currentPlayer int

	players []Player
	dealer  Player

	blackjackPayout float64
}

// Options is the options of the blackjack game.
type Options struct {
	NumberOfHands   int
	NumberOfHumans  int
	NumberOfAI      int
	NumberOfDecks   int
	BlackjackPayout float64
}

const (
	playerTurn int = iota
	dealerTurn
	handOver
)

// CreateGame creates a new blackjack game with the provided options.
func CreateGame(options Options) Game {
	game := Game{}
	game.deck = deck.NewDeck()
	game.deck.AddDecks(options.NumberOfDecks - 1)
	game.deck.Shuffle()
	for i := 0; i < options.NumberOfAI; i++ {
		game.players = append(game.players, createAIPlayer())
	}
	for i := 0; i < options.NumberOfHumans; i++ {
		game.players = append(game.players, createHumanPlayer())
	}
	game.dealer = createDealer()
	game.numberOfHands = options.NumberOfHands
	game.numberOfDecks = options.NumberOfDecks
	game.blackjackPayout = options.BlackjackPayout
	return game
}

// Play starts the new blackjack game.
func (game *Game) Play() []int {
	min := 52 * game.numberOfDecks / 3
	for i := 0; i < game.numberOfHands; i++ {
		if len(game.deck.Deck) < min {
			game.deck = deck.NewDeck()
			game.deck.AddDecks(game.numberOfDecks - 1)
			game.deck.Shuffle()
		}
		game.playerCount = len(game.players)
		game.Deal()
		game.Bet()
		game.PlayHand()
		game.EndGame()
	}
	results := []int{}
	for _, player := range game.players {
		results = append(results, player.GetBalance())
	}
	return results
}

func (game *Game) Bet() {
	for _, player := range game.players {
		player.SetBet()
	}
}

func (game *Game) Deal() {
	d := game.deck
	var c card.Card
	for i := 0; i < 2; i++ {
		for _, player := range game.players {
			c, d = draw(d)
			player.AddToHand(c, 0)
		}
		c, d = draw(d)
		game.dealer.AddToHand(c, 0)
	}
	game.stage = playerTurn
	game.currentPlayer = 0
	game.deck = d
}

func (game *Game) EndGame() {
	dealerHand := game.dealer.GetHand()
	dScore := dealerHand[0].score()
	for i, player := range game.players {
		for _, h := range player.GetHand() {
			game.getResults(i, h.score(), dScore)
		}
		game.players[i].Results(dealerHand)
		game.players[i].NewHand()
	}
	game.dealer.NewHand()
}

var (
	errBust  = errors.New("Hand score exceeded 21")
	errStand = errors.New("Player Stood")
)

type Move func(*Game, int) error

func MoveHit(game *Game, h int) error {
	var card card.Card
	score := 0
	card, game.deck = draw(game.deck)
	if game.stage == dealerTurn {
		fmt.Println("dealer turn")
		game.dealer.AddToHand(card, 0)
		score = game.dealer.GetHand()[0].score()
	} else {
		fmt.Println("player Turn")
		game.players[game.currentPlayer].AddToHand(card, h)
		score = game.players[game.currentPlayer].GetHand()[h].score()
		fmt.Println(score)
	}
	if score > 21 {
		return errBust
	}
	return nil
}

func MoveStand(game *Game, h int) error {
	if game.currentPlayer == game.playerCount-1 && h == len(game.players[game.currentPlayer].GetHand())-1 {
		game.stage = dealerTurn
	} else if game.stage == dealerTurn {
		game.stage = handOver
	}
	return errStand
}

func MoveDouble(game *Game, h int) error {
	player := game.players[game.currentPlayer]
	if len(player.GetHand()[0]) > 2 {
		return errors.New("can only double on a hand with 2 cards")
	}
	player.DoubleDown()
	MoveHit(game, 0)
	return MoveStand(game, 0)
}

func MoveSplit(game *Game, h int) error {
	player := game.players[game.currentPlayer]
	playerHand := player.GetHand()
	if playerHand[0][0].BlackjackValue() != playerHand[0][1].BlackjackValue() {
		return errors.New("Cannot split different cards")
	}
	if len(playerHand) > 1 {
		return errors.New("Can't split already split hand")
	}
	if len(playerHand[0]) > 2 {
		return errors.New("Can only split two cards")
	}
	player.Split()
	d := game.deck
	var c card.Card
	c, d = draw(d)
	player.AddToHand(c, 0)
	c, d = draw(d)
	player.AddToHand(c, 1)
	var err error = nil
	for err == nil {
		err = player.Play(player.GetHand()[0], game.dealer.GetHand()[0][1])(game, 0)
	}
	var err2 error = nil
	for err2 == nil {
		err2 = player.Play(player.GetHand()[1], game.dealer.GetHand()[0][1])(game, 1)
	}

	return err
}

func draw(deck deck.Deck) (card.Card, deck.Deck) {
	card := deck.Draw(1)
	return card[0], deck
}

func (game *Game) getResults(i, pScore, dScore int) {
	player := game.players[i]
	bet := player.GetBet()
	switch {
	case pScore > 21:
		player.SetBalance(-bet)
	case dScore > 21:
		payout := int(float64(bet) * game.blackjackPayout)
		player.SetBalance(payout)
	case pScore > dScore:
		payout := int(float64(bet) * game.blackjackPayout)
		player.SetBalance(payout)
	case dScore > pScore:
		player.SetBalance(-bet)
	case dScore == pScore:
	}
}
