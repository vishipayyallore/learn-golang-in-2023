package main

func main() {

	var card = singleCard(getNewCard())
	card.print()

	cards := deck{"Ace of Dimonds", getNewCard()}

	cards = append(cards, "Six of Spades")

	cards.print()
}

func getNewCard() string {
	return "Five of Dimonds"
}
