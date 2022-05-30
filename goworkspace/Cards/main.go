package main

func main() {

	var card = singleCard(getNewCard())
	// fmt.Println(card)
	card.print()

	cards := deck{"Ace of Dimonds", getNewCard()}
	// fmt.Println(cards)
	// cards.print()

	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	// cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func getNewCard() string {
	return "Five of Dimonds"
}
