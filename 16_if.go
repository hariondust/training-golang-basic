package main

import "fmt"

func main() {
	name := "Afi"

	if name == "Fikri" {
		fmt.Println("Hello Fikri")
	} else if name == "Dani" {
		fmt.Println("Hello Dani")
	} else if name == "Afi" {
		fmt.Println("Hello Afi")
	} else if name == "Idea" {
		fmt.Println("Hello Idea")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Name terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
