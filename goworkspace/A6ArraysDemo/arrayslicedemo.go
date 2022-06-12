package main

import "fmt"

func main() {
	message := []string{"Hello", "World"}

	updateMessage(message)

	fmt.Println(message)
}

func updateMessage(message []string) {
	message[0] = "Challo"
}
