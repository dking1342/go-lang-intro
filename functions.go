package main

import (
	"fmt"
	"math"
)

// creating a function outside the main function
// paramter in the function
func greeting(name string) {
	fmt.Printf("good day %v\n", name)
}

func parting(name string) {
	fmt.Printf("goodbye %v\n", name)
}

func cycleNames(names []string, fn func(string)) {
	for _, value := range names {
		fn(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func functions() {

	// calling functions with primitive data types
	greeting("mario")
	parting("kavooce")

	// calling function with callback function
	names := []string{"joe", "jack", "jill"}
	cycleNames(names, greeting)
	cycleNames(names, parting)

	// calling function with a return statement
	area := circleArea(10)
	fmt.Println("the area of the circle is: ", area)
	// formatting the output
	fmt.Printf("circle has an area of %0.2f", area)

}
