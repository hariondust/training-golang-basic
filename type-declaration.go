package main

import "fmt"

func main() {
	// Type Declaration = create new data type from existing data type
	// Type Declaration Usage = creating an alias

	// create new data type called "NoKTP"
	// "NoKTP" is an alias of string data type
	type NoKTP string

	var ktpFikri NoKTP = "1234567890"
	fmt.Println(ktpFikri)
	fmt.Println(NoKTP("5555555555"))
}
