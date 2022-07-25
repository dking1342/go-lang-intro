package main

import "fmt"

func conditionals() {
	age := 40

	// using operators to check truthy or false
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	// using a conditional if/else if/else
	if age < 30 {
		fmt.Println("age is less than or equal to 50")
	} else if age < 40 {
		fmt.Println("age is less than or equal to 40")
	} else {
		fmt.Println("age is greater than 40")
	}

	// nesting conditional inside of loop
	names := []string{"kavooce", "jack", "jill", "joe"}

	for index, value := range names {
		if index == 1 {
			// using the continue command
			// continues loop without executing anything else in loop
			fmt.Println("continuing at position:", index)
			continue
		}
		if index > 2 {
			// using the break command
			// breaks loop and stops
			fmt.Println("breaking at position:", index)
			break
		}
		fmt.Printf("the value at position %v is %v\n", index, value)
	}

	colors := []string{"red", "blue", "yellow", "kavooce"}

	for _, value := range colors {

		switch value {
		case "red":
			fmt.Println("this loop is red")
			break
		case "yellow":
			fmt.Println("this loop is yellow")
			break
		default:
			fmt.Println("not a valid color")
			break
		}
	}

}
