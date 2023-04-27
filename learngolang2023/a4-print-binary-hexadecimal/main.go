package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("%v as binary, %#b \n", adams, adams)
	fmt.Printf("%v as hexadecimal, %#x \n", adams, adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", a, a, a)
	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", b, b, b)
	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", c, c, c)
	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", d, d, d)
	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", e, e, e)
	fmt.Printf("Decimal: %v \tBinary: %#b \tHexadecimal: %#x \n", f, f, f)

}
