package main

import "fmt"

func main() {
	// array size need to be declared
	// array size cannot be changed after declared

	// array of string
	// array size = 2
	var names [2]string
	names[0] = "Nurul"
	names[1] = "Fikri"

	fmt.Println(names[0])
	fmt.Println(names[1])

	// declare array w/ value
	var number = [3]int{
		24,
		04,
		22,
	}
	fmt.Println(number)

	// if array size not assigned value, it will use default
	var integer = [3]int{06, 24}
	// array 2 will contain 0 as default value
	fmt.Println(integer[2])

}
