package main

import (
	"fmt"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	fmt.Printf("s: %v || type:%T \n", s, s)

	var n1 []int
	fmt.Println("uninit:", n1, n1 == nil, len(n1) == 0)
	fmt.Printf("n1: %v || type:%T \n", n1, n1)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	fmt.Printf("s: %v || type:%T \n", s, s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	fmt.Printf("s: %v || type:%T \n", s, s)

}
