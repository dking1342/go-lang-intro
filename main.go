package main

import "fmt"

var score = 99.5

func main() {
	fmt.Println("Hello world")

	greet("kavooce")

	for _, value := range points {
		fmt.Println("points: ", value)
	}

	showScore()
}

// to run the file run the command:
// go run <filename>
