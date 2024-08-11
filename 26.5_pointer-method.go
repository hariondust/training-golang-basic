package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

// jika ingin pass by reference, perlu diganti disini sebagai pointer
func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fikri := Man{"Fikri"}

	/**  Disini tidak berubah karna yg dikirim adalah duplicate "fikri"
	Bukan value sebenernya, oleh karna itu harus pass by reference
	Atau memakai pointer */
	fikri.Married()
	fmt.Println(fikri)

	fikri.MarriedPointer()
	fmt.Println(fikri)
}
