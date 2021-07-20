package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	var bruce person

	bruce.firstName = "Bruce"
	bruce.lastName = "Alencar"
	bruce.contact.email = "bruce@gmail.com"
	bruce.contact.zipCode = 10000

	fmt.Printf("%+v", bruce)

	david := person{
		firstName: "David",
		lastName:  "Alencar",
		contact: contactInfo{
			email:   "david@gmail.com",
			zipCode: 10000,
		},
	}

	fmt.Printf("%+v", david)
}
