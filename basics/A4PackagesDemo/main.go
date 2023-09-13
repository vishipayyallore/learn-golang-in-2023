package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Favorite number is: %d\n", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	displayMessage("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")
}

func displayMessage(message string) {
	fmt.Println(message)
}
