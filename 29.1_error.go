package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		// gunakan interface errors built-in golang
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		// because its interface, we can use nil
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil pembagian adalah", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
