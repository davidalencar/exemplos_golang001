package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logW struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	lw := logW{}
	io.Copy(lw, resp.Body)
}

func (logW) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 1, nil
}
