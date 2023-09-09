package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Favorite number is", rand.Intn(10))

	displayMessage("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")
}

func displayMessage(message string) {
	fmt.Println(message)
}
