package main

import "fmt"

func main() {

	// strings
	// you can type the variable initially or go will infer the type
	var name string = "kavooce" // init typed
	var greeting = "howdy"      // go inferred the type
	var friend string           // variable created and typed for later use

	fmt.Println(greeting + " " + name + " my name is " + friend)

	// setting existing variable
	friend = "jack"

	fmt.Println(greeting + " " + name + " my name is " + friend)

	// init variable
	// this syntax to set a variable can only be done inside a function
	location := "japan"

	fmt.Println(greeting + " " + name + " my name is " + friend + ". I live in " + location)

	// int- data type
	// syntax for int variables
	var age int = 30 // variable set with type included
	var price = 5    // variable set without type included but go will infer the type
	score := 7       // variable set with accompanied syntax

	fmt.Println(age, price, score)

	// bits and memory
	// int<number> will specify the size of the number
	// int can be a catch all and go will infer
	var days int8 = 7
	var temp int8 = -100
	var weeks uint = 52 // uint is only positive and you can go higher with the positive number

	fmt.Println(days, temp, weeks)

	// floats
	var height float32 = 2.5
	var weight float64 = 2333333333333333333333.7
	radius := 1.5

	fmt.Println(height, weight, radius)
}
