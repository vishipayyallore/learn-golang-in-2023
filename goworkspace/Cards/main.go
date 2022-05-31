package main

func main() {

	var card = singleCard(getNewCard())
	card.print()

	cards := getNewDeck()

	cards = append(cards, "Six of Spades")

	cards.print()
}
