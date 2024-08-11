package main

import "fmt"

/**
Type assertion adalah kemampuan merubah tipe data
Sering digunakan pada interface kosong
*/

func random() interface{} {
	// return "OK"
	// return 5
	return true
}

func main() {
	var result any = random()
	// di konversi menjadi string secara manual
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// jangan sembarangan konversi
	// jika salah konversi data type akan terjadi panic seperti kode dibawah
	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	// cara yang paling aman adalah memakai switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value) // value disini langsung menjadi string
	case int:
		fmt.Println("Int", value) // value disini langsung menjadi int
	default:
		fmt.Println("Unknown", value)
	}

}
