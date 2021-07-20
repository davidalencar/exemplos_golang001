package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	david := person{
		firstName: "David",
		lastName:  "Alencar"}

	var bruce person

	fmt.Println(david, bruce)
	bruce.firstName = "Bruce"
	bruce.lastName = "Alencar"

	fmt.Println(david, bruce)
	fmt.Printf("%+v", david)
}
