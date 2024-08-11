package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Implementasi Interface 1#
type Person struct {
	Name string
}

// karna person punya method GetName, maka dia masuk ke interface HasName
func (person Person) GetName() string {
	return person.Name
}

// Implementasi Interface 2#
type Animal struct {
	Name  string
	Color string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Fikri"}
	SayHello(person)

	animal := Animal{Name: "Zebra", Color: "Black-White"}
	SayHello(animal)
}
