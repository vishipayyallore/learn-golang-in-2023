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

	showVariablesDemo()
}

func showVariablesDemo() {
	value1 := 128

	fmt.Println("Value 1: ", value1)

	value1 = 256
	fmt.Println("Modified -> Value 1: ", value1)

	value2 := 256
	sum := value1 + value2
	fmt.Println("Value 1: ", value1, " + Value 2: ", value2, " Sum = ", sum)
}

func sayHello() {
	fmt.Println("Hello from sayHello()")
}
