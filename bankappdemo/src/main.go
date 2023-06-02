package main

import (
	"bankappdemo/app"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go!")

	customer := app.Customer{Name: "John", City: "New York", Zipcode: "12345"}
	fmt.Println(customer)

	customer2 := app.CustomerV2{Name: "John", City: "New York", Zipcode: "12345"}
	fmt.Println(customer2)

	// app.startBankServer()
}
