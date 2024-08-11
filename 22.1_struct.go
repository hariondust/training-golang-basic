package main

import "fmt"

// struct mirip seperti class pada OOP
// penulisan pada struct ditulis dengan PascalCase

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var fikri Customer
	fikri.Name = "Nurul Fikri"
	fikri.Address = "Bekasi"
	fikri.Age = 27
	fmt.Println(fikri)

	fmt.Println(fikri.Name)
	fmt.Println(fikri.Address)
	fmt.Println(fikri.Age)

	// struct literals 1#
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	fmt.Println(joko)

	// struct literals 2#
	budi := Customer{"Budi", "Malang", 28}
	fmt.Println(budi)
}
