package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Perulangan genap akan dilewati")
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}
}
