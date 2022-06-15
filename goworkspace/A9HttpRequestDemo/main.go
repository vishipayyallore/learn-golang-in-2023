package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {

	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)

		panic(err)
	}

	fmt.Println("Status Code: ", resp.Status)
	fmt.Println("Output Received: ", resp)
	fmt.Println("Output Received: ", resp.Body)

	// Readers are the most common way to read data from a response body.
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("Output Received: ", string(bs))

	// Writers are the most common way to write data to a response body.
	// io.Copy(os.Stdout, resp.Body)

	lw := logWritter{}
	io.Copy(lw, resp.Body)

	defer resp.Body.Close()
}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
