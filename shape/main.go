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

func main() {
	t := triangle{
		base:   10,
		height: 10,
	}
	s := square{
		sideLenght: 10,
	}

	fmt.Println("Triangle area:", t.getArea())
	fmt.Println("Square area:", s.getArea())
}
