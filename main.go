package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	cards = newDeckFromFile("cards.txt")

	cards.shuffle()

	hand, restOfDeck := deal(cards, 5)

	fmt.Println("HAND")
	hand.print()
	fmt.Println("REST OF DECK")
	restOfDeck.print()
}
