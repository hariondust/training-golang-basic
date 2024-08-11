package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("Selesai")

	// For dengan statement
	for number := 1; number <= 10; number++ {
		fmt.Println("Angka saat ini: ", number)
	}
	fmt.Println("10 angka berhasil ditambahkan")

	// For Range
	// Manual Range
	names := []string{"Nurul", "Fikri"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Automatic Range
	// First variable = variable name of the index / key (if using map)
	// Second variable = variable name of the value
	for index, name := range names {
		fmt.Println("Index: ", index, " = ", name)
	}

	// Automatic Range
	// Jika tidak butuh index (karna di go semua variable wajib dipanggil)
	for _, name := range names {
		fmt.Println(name)
	}

}
