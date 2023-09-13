package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Favorite number is: %d\n", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Printf("PI Value is %g.\n", math.Pi)
}

func displayMessage(message string) {
	fmt.Println(message)
}
