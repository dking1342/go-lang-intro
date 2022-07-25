package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func newReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)
	return reader
}

func createBill() bill {
	// new reader
	reader := newReader()

	// prompt for name
	name, _ := getInput("create a new bill name: ", reader)

	// create the new bill with the name
	b := newBill(name)
	fmt.Println("created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := newReader()

	option, _ := getInput("choose option (a - add item, s - save bill, t - add tip, x - exit): ", reader)

	switch option {
	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)
		fmt.Printf("you chose %v that has a price of $%v", name, price)
	case "s":
		fmt.Println("you chose s")
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)
		fmt.Printf("you gave a tip of $%v", tip)
	case "x":
		fmt.Println("thank you... see you later!")
		break
	default:
		fmt.Println("invalid choice. please try again")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
