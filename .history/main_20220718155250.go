package main

func main() {
	cards := newDeck()
	// add a new card

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

// you can initialize a variable
// and assign it a value later
// func main() {
//   var deckSize int
//   deckSize = 52
//   fmt.Println(deckSize)
// }
