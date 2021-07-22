package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Call with go run
	if len(os.Args) < 2 {
		fmt.Println("Missing file name")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
