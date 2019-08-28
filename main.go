package main

import (
	"flag"
	"fmt"

	"./blackjack"
)

func main() {
	players := flag.Int("players", 1, "help message for flagname")
	decks := flag.Int("decks", 2, "help message for flagname")
	flag.Parse()
	game := blackjack.StartGame(*players, *decks)
	var input string
	for input != "no" {
		game.Deal()
		game.PlayHand()
		game.EndGame()
		if input == "yes" {

		}
		fmt.Println("Would you like to play another hand?")
		fmt.Scanf("%s\n", &input)
	}
}
