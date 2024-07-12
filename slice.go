package main

import "fmt"

func main() {
	// slice derived from array
	// slice size can be changed
	// slice and array is connected
	// slize can access some or whole array data

	// slice has 3 data: pointer, length, capacity
	// pointer: pointer first array data on slice
	// length: size of slice
	// capacity: capacity of slice
	// length can't be higher than capacity

	// create slice from array
	// arrayName[arrayIndexStart:arrayIndexEnd]
	// arrayIndexStart default (if left empty) = first index
	// arrayIndexEnd default (if left empty) = last index
	// example if left empty: arrayName[:]
	// example if exist: array[0:4]
	// example if only one of them defined: array[:4]

	// IF array has length of 12
	// slice = array[4:7] | pointer = 4 | length = slice max index (7) -  slice min index (4) = 3 | capacity = array size (12) - slice pointer (4) = 8

	names := [...]string{"Nurul", "Fikri", "Agus", "Joko", "Budi"}
	slice1 := names[3:]
	// pointer = 3
	// length = 5 - 3 = 2
	fmt.Println("Length of slice1: ", len(slice1))
	// capacity = 5 - 3 = 2
	fmt.Println("Capacity of slice1: ", cap(slice1))
	fmt.Println(slice1)

	slice2 := names[0:4]
	// pointer = 0
	// length = 4 - 0 = 4
	fmt.Println("Length of slice2: ", len(slice2))
	// capacity = 5 - 0 = 5
	fmt.Println("Capacity of slice2: ", cap(slice2))
	fmt.Println(slice2)

}
