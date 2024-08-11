package main

import "fmt"

// interface kosong disebut juga type data 'any'
func Ups() any {
	// return 1
	// return true
	return "Ups"
}

// contoh penggunaan interface kosong asli
func Oops() interface{} {
	// return 1
	// return true
	return "Upsies"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)

	interfaceKosong := Oops()
	fmt.Println(interfaceKosong)
}
