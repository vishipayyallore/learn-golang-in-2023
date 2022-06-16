package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	// triangle is a type that implements the bot interface
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	// triangle implements the getArea method of the shape interface
	return 0.5 * t.base * t.height
}

type square struct {
	// square is a type that implements the bot interface
	sideLength float64
}

func (s square) getArea() float64 {
	// square implements the getArea method of the shape interface
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {

	t := triangle{base: 3, height: 4}
	printArea(t)

	s := square{sideLength: 5}
	printArea(s)
}
