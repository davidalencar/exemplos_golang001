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

	// Criando ponteiro
	davidPointer := &david //O operador & cria um ponteiro para variável
	davidPointer.updFirstName("Sr. David")
	davidPointer.print()
	david.print() // O valor alterado no ponteiro davidPointer reflete em david
	// Alterando um valor com uma funcao que espera um ponteiro
	david.updFirstName("Minha nossa")
	david.print()

	// Testando sintaxe para criar um pointeiro.
	(&david).updFirstName("Eita nois")
	david.print()
	// Caso eu tenha uma func que recebe um valor comum eu posso usar
	// essa sintaxe para mandar um ponteiro
	fmt.Println(&david)
}
