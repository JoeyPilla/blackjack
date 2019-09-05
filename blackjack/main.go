package blackjack

func (game *Game) PlayHand() {
	if game.stage != dealerTurn {
		for i := range game.players {
			for game.currentPlayer == i && game.stage == playerTurn {
				playerHand := game.players[i].GetHand()
				dealerHand := game.dealer.GetHand()[0]
				for i := range playerHand {
					stand := false
					for !stand {
						err := game.players[i].Play(playerHand[i], dealerHand[1])(game)
						switch err {
						case errBust:
							MoveStand(game)
							stand = true
						case errStand:
							stand = true
						case nil:
							//nothing
						default:
							panic(err)
						}
					}
				}
				game.currentPlayer++
			}
		}
	}
	for game.stage != handOver {
		dealerHand := game.dealer.GetHand()[0]
		game.dealer.Play(dealerHand, dealerHand[1])(game)
	}
}
