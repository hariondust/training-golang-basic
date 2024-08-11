package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// new operator hanya dapat digunakan pada pointer type
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
