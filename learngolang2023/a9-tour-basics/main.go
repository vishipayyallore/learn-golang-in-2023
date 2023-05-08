package main

// import "fmt"
// import "math"

// "factored" import statement.
import (
	"fmt"
	"math/rand"
)

func main() {
	generateRandomNumbers(20)

	fmt.Println(add(42, 13))
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
