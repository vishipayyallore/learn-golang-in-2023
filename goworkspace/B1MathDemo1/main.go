package main

import (
	"fmt"
	"math"
)

func main() {

	value1 := 5.2
	tothepower := 8.0

	result := math.Pow(value1, tothepower)
	fmt.Println(value1, "to the power of ", tothepower, " = ", result)
}
