package main

import "fmt"

func main() {

	var name = "Nurul Fikri"
	// e will result in a byte var
	// byte = uint8
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
