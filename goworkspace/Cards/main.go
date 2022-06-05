package main

import "fmt"

func main() {

	var card = getNewCard()
	card.print()

	cards := getNewDeck()
	cards = append(cards, "Six of Spades")
	cards.print("Current Deck")

	hand, remainingCards := deal(cards, 5)
	hand.print("Hand Deck")
	remainingCards.print("Remaining Deck")

	// Converting the Deck to String
	fmt.Println(cards.toString())

	// Saving the Deck to a File
	cards.saveToFile("./Data1/deskofcards.txt")
}
