package main

import (
	"fmt"
)

func newGetFullName() (string, string) {
	return "Fikri", "Murul"
}

func main() {
	firstName, lastName := newGetFullName()
	fmt.Println(firstName, lastName)

	newFirstName, _ := newGetFullName()
	fmt.Println(newFirstName)
}
