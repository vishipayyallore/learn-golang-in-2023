package main

import "fmt"

type deck []string

func (this deck) print() {

	for i, card := range this {
		fmt.Println(i, card)
	}

}
