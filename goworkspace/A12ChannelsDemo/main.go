package main

import (
	"fmt"
	"net/http"
	"time"
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

	// // All messages will be received in one print statement
	// Version 1
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Version 2
	// for l := range c {
	// 	time.Sleep(time.Second * 5)

	// 	go checkLink(l, c)
	// }

	// Version 3
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)

			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down !")
		c <- link

		return
	}

	fmt.Println(link, "is up !!")
	c <- link
}

// // Each message will be received in each print statement
// fmt.Println(<-c)
// fmt.Println(<-c)
// fmt.Println(<-c)
// fmt.Println(<-c)
// fmt.Println(<-c)
