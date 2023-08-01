package main

import (
	"fmt"
)

func main() {
	output := findTwoNumbersSumBruteForce([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	fmt.Println(output[0], output[1])
}

func findTwoNumbersSumBruteForce(inputNumbers []int, sumToMatch int) []int {
	for i := 0; i < len(inputNumbers)-1; i++ {
		for j := i + 1; j < len(inputNumbers); j++ {
			if inputNumbers[i]+inputNumbers[j] == sumToMatch {
				return []int{inputNumbers[i], inputNumbers[j]}
			}
		}
	}

	// Common Logic B: Return an empty slice when no match is found
	return []int{}
}
