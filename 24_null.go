package main

import "fmt"

/**
Pada bahasa pemrograman lain, value null dimiliki oleh variable kosong
Namun pada Go-Lang, variable kosong berisi default value
Terkecuali untuk beberapa tipe data seperti function, interface, map,
slice, pointer dan channel
*/

// function dibawah akan error, karna output berupa string
// func ContohError(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // pada Go-Lang ditulis nil bukan null
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	fmt.Println(data)

	// cek jika 'data' berupa nil
	if data == nil {
		fmt.Println("Data map masih kosong")
	}

	// cek jika nil berbeda dengan string kosong
	mapKosong := map[string]string{}
	fmt.Println(mapKosong)

	if mapKosong == nil {
		fmt.Println("mapKosong map masih kosong")
	} else {
		fmt.Println("mapKosong map tidak kosong")
	}

	// map tidak dianggap kosong karna default valuenya tetap ada
}
