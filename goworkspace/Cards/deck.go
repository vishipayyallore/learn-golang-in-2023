package main

import "fmt"

type singleCard string
type deck []string

func newDeck() deck {

}

func (this deck) print() {

	for i, card := range this {
		fmt.Println(i, card)
	}

}

func (this singleCard) print() {
	fmt.Println(this)
}
