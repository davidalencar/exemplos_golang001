package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, l := range links {
		go checkLink(l)
	}
}

func checkLink(l string) {
	resp, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "--> Error:", err)
		return
	}

	fmt.Println(resp.Status, l)
}
