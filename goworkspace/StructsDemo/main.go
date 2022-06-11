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

type Person1 struct {
	FirstName string
	LastName string
	// ContactInfo is a pointer to a contactInfo struct
	contactInfo
}

func main() {
	// Create a new Person struct (Method 1)
	alex := Person{"Alex", "Anderson", contactInfo{}}	
	alex.print()

	// Create a new Person struct (Method 2)
	alex = Person{FirstName: "Alex", LastName: "Anderson", Contact: contactInfo{}}
	alex.print()

	// Create a new Person struct (Method 3). Assigns ZERO Values to the fields
	var alex2 Person
	alex2.FirstName = "Alex"
	alex2.LastName = "Anderson"
	alex2.Contact.Email = "a.b@c.com"
	alex2.Contact.ZipCode = 12345
	alex2.print()

	jim := Person{
		FirstName: "Jim", 
		LastName: "Anderson",
		Contact: contactInfo{
			Email: "a.b@c.com",
			ZipCode: 12345,
		},
	}
	jim.print()

	jill := Person1{
		FirstName: "Jill", 
		LastName: "Jackson",
		contactInfo: contactInfo{
			Email: "Jill.Jackson@Sample.com",
			ZipCode: 12345,
		},
	}
	// It will not update the FirstName field. As this is Pass By Value
	jill.updateFirstName("Jilly")
	jill.print()
}

func (p Person) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}

func (p Person1) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}

func (p Person1) updateFirstName(firstName string) {
	p.FirstName = firstName
}