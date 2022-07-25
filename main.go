package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		formattedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("invalid price. please try again")
			promptOptions(b)
		}
		b.addItem(name, formattedPrice)
		fmt.Printf("you chose %v that has a price of $%0.2f\n", name, formattedPrice)
		promptOptions(b)

	case "s":
		fmt.Println("you chose to save the bill")
		fmt.Println(b.format())
	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)
		formattedTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("invalid tip. please try again")
			promptOptions(b)
		}
		b.updateTip(formattedTip)
		fmt.Printf("you gave a tip of $%0.2f\n", formattedTip)
		promptOptions(b)

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
