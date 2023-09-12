package main

import (
	"basics/A5VariablesDemo/doctor"
	"fmt"
	"math/rand"
)

func main() {

	displayMessage(doctor.Intro())

	favoriteNumberMessage := fmt.Sprintf("Favorite number is: %d", rand.Intn(10))
	fmt.Println(favoriteNumberMessage)

	message := "Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓"
	displayMessage(message)

	var userName string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&userName)
	displayMessage(fmt.Sprintf("Hello: %s. Welcome to GoLang", userName))

	var userAge int = 18
	displayMessage(fmt.Sprintf("Your Age is : %d", userAge))
}

func displayMessage(message string) {
	fmt.Println(message)
}
