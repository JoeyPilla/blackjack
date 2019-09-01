package blackjack

import "fmt"

func (h hand) dealerString() string {
	return fmt.Sprintf("%15s, %15s", "***HIDDEN***", h[1].String())
}
