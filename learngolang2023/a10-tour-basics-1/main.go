package main

import (
	"fmt"
)

func main() {
	fmt.Println("Add: ", add(42, 13))
	fmt.Println("Multiply: ", multiply(2, 3))

	a, b := swap("Go World !!", "Hello")
	fmt.Println(a, b)

	x, y := split(17)
	fmt.Println("Named Returns: ", x, y)
}

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return // A "return" statement without arguments returns the named return values. This is known as a "naked" return.
	// return x, y
}
