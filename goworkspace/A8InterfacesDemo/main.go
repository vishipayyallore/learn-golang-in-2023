package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	// englishBot is a type that implements the bot interface

}

func (eb englishBot) getGreeting() string {
	// englishBot implements the getGreeting method of the bot interface
	return "Hi there!"
}

// func printGreeting(eb englishBot) {
// 	// printGreeting function accepts a bot interface as a parameter
// 	// and prints the greeting of the bot
// 	fmt.Println(eb.getGreeting())
// }

type spanishBot struct {
	// spanishBot is a type that implements the bot interface

}

func (sb spanishBot) getGreeting() string {
	// spanishBot implements the getGreeting method of the bot interface
	return "Hola!"
}

// We cannot use the printGreeting function to print the greeting of the spanishBot
// func printGreeting(sb spanishBot) {
// 	// printGreeting function accepts a bot interface as a parameter
// 	// and prints the greeting of the bot
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	// printGreeting function accepts a bot interface as a parameter
	// and prints the greeting of the bot
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
