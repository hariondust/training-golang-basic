package main

import "fmt"

type Filter func(string) string // Create Type Declaration to Make Usage Short

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

// using function as parameter with type declaration
func sayHelloWithFilterAndTypeDeclaration(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName, "This is Type Declaration Test Case")
}

func main() {
	// usage 1# - use function as parameter
	sayHelloWithFilter("Fikri", spamFilter)

	// usage 2# - use "function as value" as parameter
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	sayHelloWithFilterAndTypeDeclaration("Anjing", spamFilter)
}
