package main

import "fmt"

func print() {

	// prints content without a new a line
	fmt.Print("hello")    // without new line carriage
	fmt.Print("howdy \n") // with a new line carriage

	// prints content then makes new line
	fmt.Println("world")

	age := 30
	name := "kavooce"

	fmt.Println("my name is", name, "and my age is", age)

	// formatted string
	// formatting takes into account the data type
	// %_ is a format specifier
	// %v is the default format
	// %q is the quote format
	// %T is the type format
	// %<specifier>f is the float format with a significant digit specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)

	// sprintf saves formmated strings
	var output = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is: ", output)
}
