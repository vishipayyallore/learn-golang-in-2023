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
	fmt.Printf("Sum of %d and %d is %d.\n", num1, num2, addNumbersV1(num1, num2))

	a, b := "world", "hello"
	fmt.Printf("Values are %s and %s.\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("Swapped values are %s and %s.\n", a, b)
}

func addNumbers(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func addNumbersV1(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}
