package main

import (
	"fmt"

	"./blackjack"
)

func print(results []int, hands int) {
	totalWinnings := 0
	for i, result := range results {
		fmt.Printf("Player %d winnings $%d.\n", i+1, result)
		totalWinnings += result
	}
	fmt.Println(totalWinnings / len(results))
}

func main() {

	options := blackjack.Options{
		NumberOfHands:   1,
		NumberOfAI:      0,
		NumberOfHumans:  2,
		NumberOfDecks:   3,
		BlackjackPayout: 1.5,
	}

	game := blackjack.CreateGame(options)
	results := game.Play()
	print(results, options.NumberOfHands)
}
