package main

import "fmt"

func main() {
	fmt.Println("Arithmetic - Unary Operator")

	// unary operator
	var num = 10
	fmt.Println("num = ", num)

	// ++ > a = a + 1
	num++
	fmt.Println("num++ = num + 1 = ", num)

	// -- > a = a - 1
	num--
	fmt.Println("num-- = num - 1 = ", num)

	var negativeNum = -num
	fmt.Println("Negative num = ", negativeNum)

	var positiveNum = +num
	fmt.Println("Positive num = ", positiveNum)

	fmt.Println("Reverse Boolean")
	var initBoolean = true
	fmt.Println("InitialBoolean = ", initBoolean)

	var reverseBoolean = !initBoolean
	fmt.Println("!initialBoolean = ", reverseBoolean)

}
