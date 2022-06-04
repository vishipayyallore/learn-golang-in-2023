package main

func main() {

	var card = getNewCard()
	card.print()

	cards := getNewDeck()
	cards = append(cards, "Six of Spades")
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
