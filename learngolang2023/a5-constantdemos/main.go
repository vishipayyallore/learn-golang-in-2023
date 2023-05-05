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
	fmt.Printf("%v of type %T \n", n, n)

	const d = 3e20 / n
	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Printf("%v of type %T \n", d, d)
	fmt.Printf("%v of type %T \n", int64(d), int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Printf("%v of type %T \n", math.Sin(n), math.Sin(n))

}
