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

	option, _ := getInput("choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(option)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)

}
