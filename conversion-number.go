package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	// success because max value is 2 billion (9 zero after number)
	fmt.Println(nilai32)

	// success because max value is 9 sextilion (18 zero after number)
	fmt.Println(nilai64)

	// number overflow because max value is 32767
	// when overflow, it will start from min value which is -32768
	fmt.Println(nilai16)

	// change int32 to 32769 to make it overflow by 2 number
	// because overflow 2 number -> will start from -32678 to number above it
	nilai32 = 32769
	nilai16 = int16(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai16)
}
