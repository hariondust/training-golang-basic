package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("Your name is blacklisted", name)
	} else {
		fmt.Println("Welcome!", name)
	}
}

func main() {
	// Usage 1# - write the function on variable
	blacklist := func(name string) bool {
		return name == "anjing" // if name == anjing -> true
	}
	registerUser("Fikri", blacklist)

	// Usage 2# - write the function directly inside the parameter
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
