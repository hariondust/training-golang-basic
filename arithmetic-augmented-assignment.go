package main

import "fmt"

func main() {
	fmt.Println("Augmented Arithmetic Assignment")

	var a = 10
	var b = 5
	fmt.Println("Var A: ", a, " Var B: ", b)

	// augmented assignment
	a += b
	fmt.Println("Augmented Assignment Addition: a += b: ", a)

	a -= b
	fmt.Println("Augmented Assignment Substraction: a -= b: ", a)

	a *= b
	fmt.Println("Augmented Assignment Multiplication: a *= b: ", a)

	a /= b
	fmt.Println("Augmented Assignment Division: a /= b: ", a)

	a %= 3
	fmt.Println("Augmented Assignment Modulo: a %= 3: ", a)
}
