package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	fmt.Println("Run application")
	// error
	logging() // disini umumnya tidak dipanggil karna atasnya error
}

func runApplicationWithDefer() {
	// meskipun defer ditaruh diatas, tetap dijalankan setelah fungsi selesai
	defer logging()

	fmt.Println("Run application")
}

func main() {
	// defer akan di eksekusi setelah semua fungsi selesai dipanggil
	// defer akan tetap dijalankan walaupun terdapat error di fungsi sebelumnya

	runApplication()
	runApplicationWithDefer()
}
