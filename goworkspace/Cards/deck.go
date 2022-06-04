package main

import "fmt"

type singleCard string
type deck []string

func getNewDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(typeOfDeck string) {
	printHeader(fmt.Sprintf("Displaying the %s of Cards !", typeOfDeck))

	for i, card := range d {
		fmt.Println(i, card)
	}

	printFooter()
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
