package main

func main() {

	cards := newDeck()

	// declare 2 deck type variables, from deal() function
	// hand, remainingCards := deal(cards, 5)

	cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")

	newCards.shuffle()
	newCards.print()
}
