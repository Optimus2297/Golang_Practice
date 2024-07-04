package main

func main() {
	cards := newDeck()
	cards.print()
	dealtCards, remianingcards := deal(cards, 7)
	dealtCards.print()
	remianingcards.print()
	cards.saveToFile("my_cards_deck")
	nd := newDeckFromFile(("my_cards_deck"))
	nd.print()
	nd.shuffle()
	nd.print()
}