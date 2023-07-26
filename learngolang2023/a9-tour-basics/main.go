package main

// import "fmt"
// import "math"

/*
	By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package
	comprises files that begin with the statement package rand.
*/

// "factored" import statement.
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	generateRandomNumbers(20)

	fmt.Println("\nMath.Pi: ", math.Pi)

	fmt.Println("Add: ", add(42, 13))
	fmt.Println("Multiply: ", multiply(2, 3))
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

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x + y
}
