package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	// disini address tidak akan berubah karna yg dikirim adalah duplicatenya
	// bukan value dari address diatas
	ChangeCountryToIndonesia(address)
	fmt.Println(address)

	address = Address{}
	// disini address akan berubah karna yg dikirim adalah value sebenernya
	ChangeCountryToIndonesiaPointer(&address)
	fmt.Println(address)

	// cara lain mengirim pointer
	var addressPointer *Address = &Address{}
	ChangeCountryToIndonesiaPointer(addressPointer)
	fmt.Println(addressPointer)
}
