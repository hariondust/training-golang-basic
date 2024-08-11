package main

import "fmt"

func main() {
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Fikri"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println("ups key has been deleted!")

	fmt.Println(book)
}
