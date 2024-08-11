package main

import "fmt"

func main() {
	// "len(array)" = get array length
	// "array[index]" = get value from index
	// "array[index] = value" = assign value to index

	// declare array w/o size
	// the size will follow value size
	var number = [...]int{
		90,
		100,
		110,
		120,
	}

	// check array "number" values
	fmt.Println(number)

	// array "number" has 4 size
	fmt.Println(len(number))

	// first value of array "number"
	fmt.Println(number[0])

	// assign last value of array "number"
	number[len(number)-1] = 200

	// check array "number" values after changed
	fmt.Println(number)
}
