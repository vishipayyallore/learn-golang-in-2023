package main

import "fmt"

func main(){
	// Method 2: Create a map of strings 
	// var colors map[string]string

	// Method 3: Create a map of strings 
	// colors := make(map[string]string)

	// Method 1: Create a map of strings 
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["whiteD"] = "#ffffff"

	fmt.Println("Maps Demo", colors)

	delete(colors, "whiteD")
	fmt.Println("Maps Demo", colors)

	printMap(colors)

}

func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("Hex code for", color, "is", hex)
	}
}
