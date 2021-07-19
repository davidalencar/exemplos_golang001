package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("./cards.txt")
	cards2 := newDeckFromFile("./cards.txt")
	fmt.Println(cards2)

}
