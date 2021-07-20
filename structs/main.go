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

func (p person) print() {

	fmt.Printf("%+v\n", p)
}

/*O operador * antes do tipo indica que a função
espera um ponteiro apontando para uma variável
daquele tipo (byRef no C#)*/
func (p *person) updFirstName(name string) {
	p.firstName = name
}

func main() {
	david := person{
		firstName: "David",
		lastName:  "Alencar",
		contact: contactInfo{
			email:   "david@gmail.com",
			zipCode: 10000,
		},
	}

	david.print()

	davidPointer := &david //O operador & cria um ponteiro para variável
	davidPointer.updFirstName("Sr. David")
	davidPointer.print()
	david.print() // O valor alterado no ponteiro davidPointer reflete em david

}
