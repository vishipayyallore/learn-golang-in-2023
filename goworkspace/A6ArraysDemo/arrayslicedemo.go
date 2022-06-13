package main

import(
	Employee "A6ArraysDemo/Employees"
	"fmt"
)

func main() {
	message := []string{"Hello", "World"}

	updateMessage(message)

	fmt.Println(message)

	alex := Employee.Person{"Alex", "Anderson"}
	fmt.Println(alex)
}

func updateMessage(message []string) {
	message[0] = "Challo"
}
