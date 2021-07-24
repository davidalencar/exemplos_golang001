package main

import "fmt"

type salesline struct {
	salesid    string
	itemid     string
	qty        float64
	unitPrince float64
	discAmount float64
	lineAmount float64
}

func calcLineAmount(l *salesline) {
	l.lineAmount = (l.unitPrince - l.discAmount) * l.qty
}

func (l *salesline) calcLineAmount() {
	l.lineAmount = (l.unitPrince - l.discAmount) * l.qty
}
func main() {
	l := salesline{
		salesid:    "001",
		itemid:     "cadeira001",
		qty:        1,
		unitPrince: 100,
		discAmount: 10,
	}

	calcLineAmount(l)
	l.calcLineAmount()

	fmt.Printf("%+v\n", l)
}
