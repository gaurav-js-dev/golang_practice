package main

import "fmt"

func main() {

	// cards := newDeck()
	// hand, remaingCards := deal(cards, 5)

	// hand.print()
	// remaingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())

}
