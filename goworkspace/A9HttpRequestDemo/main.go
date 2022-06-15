package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println("Output Received: ", string(bs))

	defer resp.Body.Close()
}
