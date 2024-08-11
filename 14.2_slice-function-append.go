package main

import "fmt"

func main() {

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println("Days Init: ", days)

	weekdaySlice := days[5:]
	fmt.Println("Weekday Slice = ", weekdaySlice) // Senin, Selasa, Rabu, Kamis, Jumat

	fmt.Println("Pointer Slice = 5, Length = Slice Max (7) - Slice Pointer (5), Capacity = Array Size (7) - Slice Pointer (5)")
	fmt.Println("Length of slice: ", len(weekdaySlice))
	fmt.Println("Capacity of slice: ", cap(weekdaySlice))

	// when change slice data, array will also changed
	// it happen because slice = actual array value
	fmt.Println("Days Index 6 = ", days[6])
	fmt.Println("weekdaySlice Index 1 = ", weekdaySlice[1])

	weekdaySlice[1] = "Minggu Berkah"
	fmt.Println("Days After Slice Update: ", days) // Minggu -> Minggu Berkah

	weekdaySlice2 := append(weekdaySlice, "Minggu Biasa")
	fmt.Println("Weekday Slice 2 = ", weekdaySlice2)
	fmt.Println("Days After Weekday Slice 2: ", days) // The value is the same

	weekdaySlice2[0] = "Sabtu Lama"
	fmt.Println("Weekday Slice 2 = ", weekdaySlice2)
	fmt.Println("Days After Update Weekday Slice 2: ", days) // The value is the same
}
