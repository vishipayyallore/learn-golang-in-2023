package main

import "fmt"

type singleCard string

func getNewCard() singleCard {
	return "Five of Dimonds"
}

func (sc singleCard) print() {
	printHeader("Displaying the Current Card !")

	fmt.Println(sc)

	printFooter()
}
