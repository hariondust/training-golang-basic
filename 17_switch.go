package main

import "fmt"

func main() {
	name := "Fikri"

	switch name {
	case "Fikri":
		fmt.Println("Hello Fikri")
	case "Afi":
		fmt.Println("Hello Afi")
	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudaah sesuai")
	}

	// switch tanpa kondisi
	// not recommended (better use IF statement)
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length < 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah sesuai")
	}

}
