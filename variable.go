package main

import "fmt"

// Test variable usage on Go
func main() {
	// variable can be declared with data type
	// if declared with data type; don't need to assign value
	var name string

	name = "Fikri"
	fmt.Println(name)

	// Update variable value
	name = "Nurul Fikri"
	fmt.Println(name)

	// variable can be declared WITHOUT data type
	// if declared without data type; need to assign value
	var jobs = "Software Developer"
	fmt.Println(jobs)

	// variable can also be declared INITIAlLY, WITHOUT "var"
	// if declared without "var" need to assign value and use ":="
	workplace := "Andal"
	fmt.Println(workplace)

	// to update variable without "var" can use the same way as previous
	workplace = "Xtremax"
	fmt.Println(workplace)

	// we can also declare more than 1 variable at once
	var (
		firstName = "Nurul"
		lastName  = "Fikri"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
