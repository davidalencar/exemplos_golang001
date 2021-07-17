package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	fmt.Println("Na minha mão")
	hand.print()

	fmt.Println("Na mesa")
	remainingCards.print()

	fmt.Println("E minhas cartas estao como?")
	cards.print()
}
