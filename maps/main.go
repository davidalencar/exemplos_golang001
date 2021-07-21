package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff9886",
		"green": "#ff9886",
		"white": "#ff9886",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
