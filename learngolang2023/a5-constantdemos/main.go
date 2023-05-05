package main

import (
	"fmt"
	"math"
)

const s string = "constant"

var age int = 10

func main() {

	fmt.Printf("%v of type %T \n", s, s)
	fmt.Printf("%v of type %T \n", age, age)

	// Short hand declaration can be used to declare and initialize a variable. It is allowed only inside a function.
	name := "John"
	fmt.Println(name)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
