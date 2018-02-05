package main

func main() {
	/*
		var card string = "Ace of Spades"
		card := newCard() //"Ace of spades"
		card = "Five of Diamonds"
	*/

	cards := newDeck()

	/*
		cards.saveToFile("my_cards")

		newCards := newDeckFromFile("my_cards")

		hand, remainingDeck := deal(newCards, 5)

		hand.print()
		remainingDeck.print()
	*/

	cards.shuffle()

	cards.print()
}
