package main

import "fmt"

func getFullName() (string, string) {
	return "Nurul", "Fikri"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
