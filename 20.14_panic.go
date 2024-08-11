package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Application error")
	}
}

func main() {
	// test jika panic = false
	runApp(false)

	// test jika panic = true (defer akan tetap di eksekusi)
	// meskipun begitu, log defer akan muncul diatas (berbeda dengan umumnya dibawah)
	runApp(true)
}
