package main

import "fmt"

type contactInfo struct {
	Email string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName string
	// ContactInfo is a pointer to a contactInfo struct
	Contact contactInfo
}

func main() {
	// Create a new Person struct (Method 1)
	alex := Person{"Alex", "Anderson", contactInfo{}}	
	printPerson(alex)

	// Create a new Person struct (Method 2)
	alex = Person{FirstName: "Alex", LastName: "Anderson", Contact: contactInfo{}}
	printPerson(alex)

	// Create a new Person struct (Method 3). Assigns ZERO Values to the fields
	var alex2 Person
	alex2.FirstName = "Alex"
	alex2.LastName = "Anderson"
	alex2.Contact.Email = "a.b@c.com"
	alex2.Contact.ZipCode = 12345
	printPerson(alex2)

	jim := Person{
		FirstName: "Jim", 
		LastName: "Anderson",
		Contact: contactInfo{
			Email: "a.b@c.com",
			ZipCode: 12345,
		},
	}
	printPerson(jim)
}

func printPerson(p Person) {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}
