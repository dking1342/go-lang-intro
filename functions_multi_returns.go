package main

import (
	"fmt"
	"strings"
)

// return multiple values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func multiple_returns() {
	firstName, lastName := getInitials("joe moe")
	fmt.Println(firstName, lastName)

	firstName2, lastName2 := getInitials("kavooce")
	fmt.Println(firstName2, lastName2)
}
