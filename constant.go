package main

import "fmt"

func main() {
	// unlike var, const can be declared without being used
	const firstName string = "Nurul"
	const lastName = "Fikri"

	// error
	// const cannot be re-assigned a value
	// firstName = "Nurul Fikri"

	// declare multiple const
	const (
		address   = "Bandung"
		workplace = "Xtremax"
	)

	fmt.Println(address)
	fmt.Println(workplace)
}
