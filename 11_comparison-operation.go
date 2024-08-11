package main

import "fmt"

func main() {
	// comparison operation = comparing 2 variable
	// comparison operation = return boolean value
	// comparison operation return TRUE if correct
	// comparison operation return FALSE if incorrect

	// comparison operator for number (>, <, >=, <=)
	num1 := 5
	num2 := 10
	num3 := 5
	fmt.Println("num1 = ", num1, " num2 = ", num2, " num3 = ", num3)

	moreThan := num1 > num2
	fmt.Println("num1 > num2 = ", moreThan)

	lessThan := num1 < num2
	fmt.Println("num1 < num2 = ", lessThan)

	moreThanEqual := num1 >= num3
	fmt.Println("num1 >= num3 = ", moreThanEqual)

	lessThanEqual := num1 <= num2
	fmt.Println("num1 <= num2 = ", lessThanEqual)

	equal := num1 == num2
	fmt.Println("num1 == num2 = ", equal)

	notEqual := num1 != num2
	fmt.Println("num1 != num2 = ", notEqual)
}
