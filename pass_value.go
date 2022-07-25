package main

import "fmt"

// non-pointer values
// strings
// ints
// floats
// booleans
// arrays
// structs

// pointer wrapper values
// slices
// maps
// functions

// depending on data type it will behave differently due to how it is stored in memory
// pointer is a data type itself

// updates copy of name
func updateName(n string) string {
	n = "gogo"

	// returning the value will allow it to modify original value
	return n
}

// updates map due to the pointer in memory
func updateMenu(m map[string]float64) {
	m["coffee"] = 2.99
}

// updates the pointer to update the value
// astrick will denote the value at the memory location
func updateLocation(location *string) {
	// use the astrick to denote the value in memory
	*location = "France"
}

func main() {
	name := "hoho"

	// makes a copy of the original
	// function updates the copy, not the original
	updateName(name)

	// shows the original
	fmt.Println(name)

	// correctly setting the name to a new value
	// setting name to the value created in the function, return value
	name = updateName(name)

	// shows the new value
	fmt.Println(name)

	// passing values of other data types, slices, maps, functions, etc

	menu := map[string]float64{
		"soup": 4.99,
		"pie":  7.99,
	}

	fmt.Println(menu)

	// setting new value
	updateMenu(menu)
	fmt.Println(menu)

	// pointers //////
	location := "USA"

	// & will give the address in the memory
	locationPointer := &location
	fmt.Println("memory address of location is: ", locationPointer)
	fmt.Println("value at memory address is: ", *locationPointer)
	updateLocation(locationPointer)
	fmt.Println("location: ", *locationPointer)

}
