package main

import "fmt"

// harap gunakan closure dengen bijak agar tidak membuat dev lain bingung

func main() {
	counter := 0

	// function ini berinteraksi dengan data di sekitarnya (counter berada diatasnya)
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	// asumsikan disini ada banyak code
	// sampai sini

	// namun function (closure) ini dipanggil setelah banyak code tadi
	// jika tidak dipakai dengan bijak akan bikin bingung
	// karna referencenya terlalu jauh diatas
	increment()
	increment()
	increment()

	fmt.Println(counter)
}
