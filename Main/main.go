package main

import (
	"fmt"

	blackjack "github.com/gomark/blackjack_ai/blackjack_w_ai/blackjack_ai/blackjack_pkg"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)

}
