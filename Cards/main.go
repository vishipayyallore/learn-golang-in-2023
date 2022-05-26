package main

import "fmt"

func main() {

	card := getNewCard()

	fmt.Println(card)
}

func getNewCard() string {
	return "Five of Dimonds"
}
