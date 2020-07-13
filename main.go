package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://netflix.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.de",
	}

	// create a channel of string
	c := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, c)
	}

	for {

		go func() {
			time.Sleep(5 * time.Second)
			checkLinkStatus(<-c, c)
		}()
	}

}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
