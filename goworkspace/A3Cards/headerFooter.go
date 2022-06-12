package main

import "fmt"

func printHeader(headerTitle string) {
	fmt.Println("=======================================================")
	fmt.Println(headerTitle)
	fmt.Println("=======================================================")
}

func printFooter() {
	fmt.Println("-------------------------------------------------------")
	fmt.Println()
}
