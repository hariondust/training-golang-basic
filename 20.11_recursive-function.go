package main

import "fmt"

// jika menggunakan for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// jika menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	// cara biasa
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)

	// cara loop
	fmt.Println(factorialLoop(10))

	// cara recursive
	fmt.Println(factorialRecursive(10))
}
