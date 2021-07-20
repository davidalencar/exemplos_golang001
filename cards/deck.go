package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" de "+suit)
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	data, e := ioutil.ReadFile(fileName)

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}

	innerText := string(data)
	s := strings.Split(innerText, ",")

	return deck(s)
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

func (d deck) shuffler() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
