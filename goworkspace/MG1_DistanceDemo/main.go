package main

import "fmt"

const secondsInHour = 3600
const milesConversion = 0.621

func main() {
	fmt.Println("Number of seconds in hour is: ", secondsInHour)

	distance := 60.8
	fmt.Printf("The distance in miles is %f \n", (distance * milesConversion))
}
