package main

import "fmt"

func endApplication() {
	fmt.Println("Application ended")

	message := recover() // recover will get value from panic
	fmt.Println("panic is called with message:", message)
}

func runApplicationWithRecover(error bool) {
	defer endApplication()

	if error {
		panic("ERROR")
	}
}

func main() {
	runApplicationWithRecover(true)
	// if no recover inside defer, this Println wont be called
	fmt.Println("Test panic continue or not after recover")
}
