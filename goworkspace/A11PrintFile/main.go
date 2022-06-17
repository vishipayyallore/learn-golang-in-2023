package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Arguments Received: ", os.Args)

	if len(os.Args) > 1 {
		printFile(os.Args[1])
	} else {
		fmt.Println("Usage: C:>\"go run .\\main.go ./Data/Sample.txt\"")
	}
}

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func printFile(fn string) {
	file, err := os.Open(fn)

	if err != nil {
		fmt.Println(err)
		return
	}

	lw := logWritter{}
	io.Copy(lw, file)

	defer file.Close()
}
