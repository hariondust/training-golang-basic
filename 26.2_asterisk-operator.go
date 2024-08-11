package main

import "fmt"

type Address struct {
	City, Province, Country string
}

/**
Ketika Pointer value diubah typenya, maka hanya mengubah value pointer tsb
Jika ingin mengubah value dari variable Parentnya juga, gunakan *
*/

func main() {
	// Contoh Penggunaan Pointer
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // kalau pointer depan type dikasih *
	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// Jika hanya ingin mengubah Address2 tanpa Parent pakai code dibawah
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // Address1 tidak ikut berubah
	// fmt.Println(address2) // Address2 sudah memiliki type data baru

	// Jika ingin mengubah Address2 dan Parent pakai code dibawah
	// Tambahkan bintang dan tidak memakai type data pointer lagi
	*address2 = Address{"Jogja", "DIY Jogjakarta", "Indonesia"}
	fmt.Println(address1) // Address1 mengikuti Address2
	fmt.Println(address2)

}
