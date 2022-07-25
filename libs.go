package main

import (
	"fmt"
	"sort"
	"strings"
)

func libs() {
	greeting := "howdy there!"

	// practicing strings package methods
	fmt.Println(strings.Contains(greeting, "howdy"))
	fmt.Println(strings.ReplaceAll(greeting, "howdy", "hello"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "o"))
	fmt.Println(strings.Split(greeting, " "))

	// original string
	fmt.Println("original string is: ", greeting)

	// practicing sort package methods
	// sort will modify original set of data
	ages := []int{45, 20, 35, 55, 65, 33, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	fmt.Println("ages slice: ", ages)

	// check for exisiting number
	indexThere := sort.SearchInts(ages, 35)
	fmt.Println(indexThere)

	// check for non exisiting number. it will show the index of where it would be if it were in the slice
	indexNotThere := sort.SearchInts(ages, 30)
	fmt.Println(indexNotThere)

	// slices with strings
	names := []string{"kavooce", "jack", "joe", "jill", "jonah", "julia"}
	sort.Strings(names)
	fmt.Println(names)

	// search string in slice
	fmt.Println(sort.SearchStrings(names, "kavooce"))

}
