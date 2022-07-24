package main

import "fmt"

func arrays() {

	// arrays
	// setting arrays with different syntax
	var counts [3]int = [3]int{1, 2, 3} // type on the left hand side
	var ages = [4]int{12, 22, 32, 44}   // go infer the type from right hand side
	names := [4]string{"kavooce", "jack", "jill", "joe"}

	fmt.Println(counts, len(counts))
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{12, 23, 55} // no length specified

	fmt.Println(scores, len(scores))

	// setting value within a slice
	scores[2] = 1

	fmt.Println(scores, len(scores))

	// appending value to slice
	scores = append(scores, 77)

	fmt.Println(scores, len(scores))

	// slice ranges
	// includes first parameter but one less than second parameter
	// if first parameter not present then it will start at the beginning of the array
	// if second paramter not present then it will go until the end of the array
	nameRangeFixed := names[1:3]
	nameRangeStart := names[:2]
	nameRangeEnd := names[2:]

	fmt.Println(nameRangeFixed)
	fmt.Println(nameRangeStart)
	fmt.Println(nameRangeEnd)

	// append to slice
	nameRangeFixed = append(nameRangeFixed, "koopa")
	fmt.Println(nameRangeFixed)

}
