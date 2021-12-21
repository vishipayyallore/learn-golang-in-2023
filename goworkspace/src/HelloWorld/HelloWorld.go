package main

import "fmt"

func main() {

	n, _ := fmt.Println("Hello World !!")
	fmt.Println(n)

	sayHello()

	for i := 0; i <= 10; i++ {
		n, err := fmt.Println("Number: ", i, "; also boolean: ", true)
		fmt.Println(n, err)
	}
}

func sayHello() {
	fmt.Println("Hello from sayHello()")
}
