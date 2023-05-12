package main

import (
	"fmt"
	"time"
)

func main() {
	showBasicSwitchDemo()

	showCommasSeparateMultipleExpressionsDemo()

	showSwitchWithoutExpressionDemo()
}

// basic switch
func showBasicSwitchDemo() {
	fmt.Println("***** basic switch *****")
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}

// Commas to separate multiple expressions in the same case statement and optional default case as well.
func showCommasSeparateMultipleExpressionsDemo() {
	fmt.Println("\n***** Commas to separate multiple expressions in the same case statement *****")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}

// switch without an expression is an alternate way to express if/else logic.
// Here we also show how the case expressions can be non-constants.
func showSwitchWithoutExpressionDemo() {
	fmt.Println("\n***** switch without an expression *****")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon ", t.Hour())
	default:
		fmt.Println("It's after noon", t.Hour())
	}
}
