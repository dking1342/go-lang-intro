package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
	return b
}

// get the bill
func (b *bill) format() string {
	formattedString := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		formattedString += fmt.Sprintf("%-25v ... $%0.2f\n", key+":", value)
		total += value
	}

	// tip
	total += b.tip
	formattedString += fmt.Sprintf("%-25v ... $%0.2f\n", "tip:", b.tip)

	// total
	formattedString += fmt.Sprintf("%-25v ... $%0.2f\n", "total:", total)

	return formattedString
}

// update the bill - tip
// make sure to add asterick to object for the pointer *bill
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// update the bill - items
// make sure to add asterick to object for the pointer *bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
