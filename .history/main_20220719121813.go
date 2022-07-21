package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

// you can initialize a variable
// and assign it a value later
// func main() {
//   var deckSize int
//   deckSize = 52
//   fmt.Println(deckSize)
// }
