package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}

// you can initialize a variable
// and assign it a value later
// func main() {
//   var deckSize int
//   deckSize = 52
//   fmt.Println(deckSize)
// }
