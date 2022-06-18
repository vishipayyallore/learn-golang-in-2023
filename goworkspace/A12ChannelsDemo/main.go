package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// // Each message will be received in each print statement
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// // All messages will be received in one print statement
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		// fmt.Println(link, "might be down !")
		c <- fmt.Sprintf("%s might be down", link)

		return
	}

	// fmt.Println(link, "is up !!")
	c <- fmt.Sprintf("%s is up !!", link)
}
