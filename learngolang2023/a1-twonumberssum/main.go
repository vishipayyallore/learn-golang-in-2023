package main

import (
	"fmt"
)

func main() {
	inputSets := generateTwoNumbersSumInputs()

	for i, inputSet := range inputSets {
		fmt.Printf("Input Set %d:\n", i+1)
		if inputSet.InputNumbers == nil {
			fmt.Println("Error: Input array is nil.")
			continue
		}
		fmt.Printf("Input Array: %v\n", inputSet.InputNumbers)
		fmt.Printf("Sum to Match: %d\n\n", inputSet.SumToMatch)
	}
}

type twoNumbersSumInputSet struct {
	InputNumbers []int
	SumToMatch   int
}

func generateTwoNumbersSumInputs() []twoNumbersSumInputSet {
	return []twoNumbersSumInputSet{
		{InputNumbers: []int{3, 5, -4, 8, 11, 1, -1, 6}, SumToMatch: 10},
		{InputNumbers: []int{2, 4, 6, 8, 10}, SumToMatch: 12},
		{InputNumbers: []int{-3, 1, 5, 8, -2}, SumToMatch: 6},
		{InputNumbers: []int{0, 1, 2, 3, 4}, SumToMatch: 8},
		{InputNumbers: nil, SumToMatch: 15}, // Test case with nil InputNumbers
	}
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
