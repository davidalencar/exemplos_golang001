package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Espadas", "Paus", "Diamantes", "Coração"}
	cardValues := []string{"As", "Dois", "Tres", "Quatro"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" de "+value)
		}
	}

	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	data := []byte(d.toString())
	return ioutil.WriteFile(fileName, data, 0666)
}
