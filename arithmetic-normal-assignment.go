package main

import "fmt"

func main() {
	// math operation on go:
	// + = addition; - = substraction;
	// * = multiplication; / = division;
	// % = modulo;

	fmt.Println("Normal Arithmetic Assignment")

	var a = 10
	var b = 5
	fmt.Println("Var A: ", a, " Var B: ", b)

	// normal assignment
	var c = a + b
	fmt.Println("Addition: a + b = ", c)

	var d = a - b
	fmt.Println("Substraction: a - b = ", d)

	var e = a * b
	fmt.Println("Multiplication: a * b = ", e)

	var f = a / b
	fmt.Println("Division: a / b = ", f)

	var g = a % 3
	fmt.Println("Modulo: a % 3 = ", g)
}
