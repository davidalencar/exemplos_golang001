package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, c)
		}(l)
	}

}

func checkLink(l string, c chan string) {
	resp, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "--> Error:", err)
		c <- l
		return
	}

	fmt.Println(resp.Status, l)
	c <- l
}
