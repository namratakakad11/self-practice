package main

func main() {

	cards := newDeck()

	cards.shuffle()

	cards.print()

	// cards.saveToFile("mycard")

	// cards.print()
	// hand, remainCards := deal(cards,3)

	// hand.print()
	// remainCards.print()

}
