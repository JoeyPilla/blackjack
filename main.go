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
	fmt.Println(float64(totalWinnings) / float64(len(results)) / float64(hands))
}

func main() {

	options := blackjack.Options{
		NumberOfHands:   99999999,
		NumberOfAI:      5,
		NumberOfHumans:  0,
		NumberOfDecks:   6,
		BlackjackPayout: 1.5,
	}

	game := blackjack.CreateGame(options)
	results := game.Play()
	print(results, options.NumberOfHands)
}
