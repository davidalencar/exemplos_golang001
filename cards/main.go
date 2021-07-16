package main

func main() {
	cards := newDeck()

	cards = append(cards, "6 de Espadas")
	cards.print()
}

func newCard() string {
	return "5 de espadas"
}
