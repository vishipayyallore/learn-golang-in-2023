package main

import (
	"bankappdemo/app"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go!")

	customer := app.Customer{Name: "John", City: "New York", Zipcode: "12345"}
	fmt.Println(customer)

	customer2 := app.CustomerV2{Name: "Jean", City: "New Jersey", Zipcode: "23456"}
	fmt.Println(customer2)

	app.StartBankServer()
}
