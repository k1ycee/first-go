package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bill
func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Format the bill
func (b *Bill) format() string {
	fs := b.name + " Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ..... $%v\n", k+":", v)
		total += v
	}

	//show tip
	fs += fmt.Sprintf("%-25v ..... $%0.2f\n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ..... $%0.2f", "total:", total+b.tip)

	return fs
}

// update tip
func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

// update items
func (b *Bill) updateItems(name string, price float64) {
	b.items[name] = price
}
