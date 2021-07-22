package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func main() {
	printArea(triangle{base: 10, height: 10})
	printArea(square{sideLenght: 10})
}
