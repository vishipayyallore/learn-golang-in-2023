package main

import "fmt"

func main() {

	value1 := 128

	fmt.Println("Value 1: ", value1)

	value1 = 256
	fmt.Println("Modified -> Value 1: ", value1)

	value2 := 256
	sum := value1 + value2
	fmt.Println("Value 1: ", value1, " + Value 2: ", value2, " Sum = ", sum)
}
