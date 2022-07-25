package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	formattedName := strings.Title(strings.ToLower(name))
	b := bill{
		name:  fmt.Sprintf("%v's bill", formattedName),
		items: map[string]float64{},
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

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	// writing a file
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
