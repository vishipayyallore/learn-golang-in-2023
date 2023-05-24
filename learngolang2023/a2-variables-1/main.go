package main

import "fmt"

var c, python, java bool
var n1, n2 int = 1, 2

func main() {

	fmt.Println("***** Printing Variables with default Values *****")
	printVariablesWithDefaults()

	fmt.Println("\n***** Printing Variables Initializers Values *****")
	variablesWithInitializers()

	fmt.Println("\n***** Printing Variables With Short Assignment *****")
	variablesWithShortAssignment()
}

func printVariablesWithDefaults() {
	var i int
	var name string

	fmt.Printf("a: %v || type:%T \n", c, c)
	fmt.Printf("python: %v || type:%T \n", python, python)
	fmt.Printf("java: %v || type:%T \n", java, java)
	fmt.Printf("i: %v || type:%T \n", i, i)
	fmt.Printf("name: %v || type:%T \n", name, name)
}

func variablesWithInitializers() {
	var c1, python1, java1 = "no!", true, false
	k := 3.6

	fmt.Printf("n1: %v || type:%T \n", n1, n1)
	fmt.Printf("n2: %v || type:%T \n", n2, n2)
	fmt.Printf("c1: %v || type:%T \n", c1, c1)
	fmt.Printf("python1: %v || type:%T \n", python1, python1)
	fmt.Printf("java1: %v || type:%T \n", java1, java1)
	fmt.Printf("k: %v || type:%T \n", k, k)
}

func variablesWithShortAssignment() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Printf("i: %v || type:%T \n", i, i)
	fmt.Printf("j: %v || type:%T \n", j, j)
	fmt.Printf("k: %v || type:%T \n", k, k)

	fmt.Printf("c: %v || type:%T \n", c, c)
	fmt.Printf("python: %v || type:%T \n", python, python)

	fmt.Printf("java: %v || type:%T \n", java, java)
}
