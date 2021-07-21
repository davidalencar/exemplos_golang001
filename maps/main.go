package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	fmt.Println(colors)
	delete(colors, "red")

	fmt.Println(colors)
}
