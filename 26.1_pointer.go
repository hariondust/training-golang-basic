package main

import "fmt"

/**
Secara default, variable Go-Lang di passing by value, bkn reference
Artinya, jika mengirim variable dalam function, method atau variable lain
sebenernya yang dikirim adalah duplikasi dari valuenya
*/

type Address struct {
	City, Province, Country string
}

func main() {
	// Contoh Pass by Value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah, karna value duplicate
	fmt.Println(address2)

	// Contoh Pass by Reference (Pointer)
	// Jadi datanya tidak di duplicate, melainkan pakai data yang sama
	// Penggunaan pointer memakai operator "&" diikuti nama variable

	address3 := Address{"Madiun", "Jawa Timur", "Indonesia"}
	address4 := &address3 // pointer

	address4.City = "Malang"

	fmt.Println(address3) // address3 berubah, karna address4 pakai pointer
	fmt.Println(address4)

}
