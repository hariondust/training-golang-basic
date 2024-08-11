package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(20, 20, 20))

	// if want to use slice to insert the value
	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...)) //add 3 dots behind the slice variable
}
