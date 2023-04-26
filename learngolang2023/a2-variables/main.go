package main

import "fmt"

func main() {

	fmt.Print("***** Printing Values *****")
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Print("***** Printing Values ***** \n\n")

	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)

	var aa = "initial"
	fmt.Println(aa)

	var bb, cc int = 1, 2
	fmt.Println(bb, cc)

	var dd = true
	fmt.Println(dd)

	var e int
	fmt.Println(e)

	ff := "apple"
	fmt.Println(ff)

}
