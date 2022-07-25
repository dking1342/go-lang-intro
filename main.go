package main

import "fmt"

func main() {
	// create new bill
	myBill := newBill("kavooce's bill")

	// add items
	myBill.addItem("onion soup", 4.5)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	// update tip
	myBill.updateTip(10)

	// get the bill (formatted)
	fmt.Println("********** get formatted bill **********")
	fmt.Println(myBill.format())

}
