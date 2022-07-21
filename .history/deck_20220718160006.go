package main

import (
	"fmt"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// should create and return a list of playing cards
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// d is the receiver argument
// receiver arguments are like 'this' or 'self'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (deck, deck) tells us what the return should be
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	[]string(d)

}
