package blackjack

import (
	"fmt"
)

func (game *Game) PlayHand() {
	var input string
	for i := range game.players {
		for game.currentPlayer == i && game.stage != dealerTurn {
			fmt.Printf("Player %d: %s\n", i+1, game.players[i])
			fmt.Printf("%-8s: %s\n", fmt.Sprintf("Dealer"), game.dealer.dealerString())
			fmt.Println()
			fmt.Print("What will you do? (h)it, (s)tand:")
			fmt.Scanf("%s", &input)
			fmt.Println()
			switch input {
			case "h":
				game.hit()
			case "s":
				game.stand()
			default:
				fmt.Println("Invalid option:", input)
			}
		}
	}
	for game.dealer.score() <= 16 || (game.dealer.score() == 17 && game.dealer.score() != 17) {
		game.hit()
	}
}
