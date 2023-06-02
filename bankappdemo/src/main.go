package main

import (
	"bankappdemo/app"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go!")

	customer := app.Customer{Name: "John", City: "New York", Zipcode: "12345"}
	fmt.Println(customer)

	app.StartBankServer()
}
