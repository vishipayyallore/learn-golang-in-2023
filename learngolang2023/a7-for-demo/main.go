package main

import (
	"fmt"
	"math/rand"
)

func main() {
	showForLoopsDemo()

	generateRandomNumbers(20)
}

func generateRandomNumbers(stopValue int) {
	i := 1
	fmt.Println("Random numbers:")
	for i <= stopValue {
		fmt.Print(getRandomNumber(10*i), " | ")
		i = i + 1
	}
}

func getRandomNumber(input int) int {
	return rand.Intn(input)
}

func showForLoopsDemo() {

	fmt.Println("for .. with a single condition.")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("for .. classic initial/condition/after.")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("for .. Without Condition.")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("for .. Without continue.")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
