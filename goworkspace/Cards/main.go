package main

import "fmt"

func main() {

	card := getNewCard()
	fmt.Println(card)

	cards := []string{"Ace of Dimonds", getNewCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func getNewCard() string {
	return "Five of Dimonds"
}
