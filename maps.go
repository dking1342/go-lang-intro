package main

import "fmt"

func maps() {

	// practicing maps data type
	// similar to a hash, dictionary, object
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.99,
	}

	// print all menu
	fmt.Println(menu)

	// print one item of the menu
	fmt.Println(menu["pie"])
	for key, value := range menu {
		fmt.Printf("the menu item %v has a price of %v\n", key, value)
	}

	// ints as a key type
	phonebook := map[int]string{
		2345342: "mario",
		3423456: "luigi",
		2239956: "toad",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[2239956])

	// update item in map
	phonebook[2239956] = "princess"
	fmt.Println(phonebook)

}
