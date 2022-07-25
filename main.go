package main

import "fmt"

func main() {
	myBill := newBill("kavooce's bill")
	fmt.Println(myBill)
	copy := myBill.format()
	fmt.Println(copy)
}
