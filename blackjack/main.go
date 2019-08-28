package blackjack

import (
	"fmt"
)

func (game *Game) PlayHand() {
	var input string
	for i := range game.Players {
		for game.CurrentPlayer == i && game.Stage != DealerTurn {
			fmt.Printf("Player %d: %s\n", i+1, game.Players[i])
			fmt.Printf("%-8s: %s\n", fmt.Sprintf("Dealer"), game.Dealer.DealerString())
			fmt.Println()
			fmt.Print("What will you do? (h)it, (s)tand:")
			fmt.Scanf("%s", &input)
			fmt.Println()
			switch input {
			case "h":
				game.Hit()
			case "s":
				game.Stand()
			default:
				fmt.Println("Invalid option:", input)
			}
		}
	}
	for game.Dealer.Score() <= 16 || (game.Dealer.Score() == 17 && game.Dealer.Score() != 17) {
		game.Hit()
	}
}
