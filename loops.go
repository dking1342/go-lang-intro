package main

import "fmt"

func loops() {

	// for loop behaving like a while loop
	index := 0
	for index <= 5 {
		fmt.Println("value of index is: ", index)
		index++
	}

	// for loop
	for x := 0; x <= 5; x++ {
		fmt.Println("value of x is: ", x)
	}

	// looping through a slice
	names := []string{"kavooce", "jack", "jill", "joe"}

	for i := 0; i < len(names); i++ {
		fmt.Println("name: ", names[i])
	}

	// another looping strategy
	for index, value := range names {
		fmt.Printf("the value at index %v is %v\n", index, value)
	}
	// when you do not want to use index or value then just put a _ at that place
	// it will not modify value if you want to change it
	for _, value := range names {
		fmt.Printf("the value is %v\n", value)
	}

}
