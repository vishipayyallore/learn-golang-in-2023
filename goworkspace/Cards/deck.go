package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		print("Error occurred at newDeckFromFile() - ", err)
		os.Exit(1)
	}

	deckFromFile := strings.Split(string(bs), ",")

	return deck(deckFromFile)
}
