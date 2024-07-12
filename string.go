package main

import "fmt"

// Test string and string function
func main() {
	fmt.Println("This is a string!")
	fmt.Println("Fikri")
	fmt.Println("Nurul Fikri")

	// len: get the length of the string
	fmt.Println(len("Fikri"))

	// array: get the character according to array
	// array starts from 0, not 1
	// array will get any character including whitespace
	// array will get the byte value of the string
	// array byte need to be converted to get the string character
	fmt.Println("Nurul Fikri"[0])
	fmt.Println("Nurul Fikri"[4])
}
