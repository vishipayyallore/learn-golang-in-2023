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

	num1 := 2
	num2 := 3
	fmt.Printf("Sum of %d and %d is %d.\n", num1, num2, addNumbers(num1, num2))

	num1 = 5
	num2 = 7
	fmt.Printf("Sum of %d and %d is %d.\n", num1, num2, addNumbers(num1, num2))
}

func addNumbers(x int, y int) int {
	return x + y
}
