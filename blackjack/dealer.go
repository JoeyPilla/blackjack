package blackjack

import "fmt"

func (h Hand) DealerString() string {
	return fmt.Sprintf("%15s, %15s", "***HIDDEN***", h[1].String())
}
