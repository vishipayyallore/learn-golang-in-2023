package main

import "fmt"

func main() {

	card := getNewCard()
	fmt.Println(card)

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
