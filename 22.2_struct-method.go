package main

import "fmt"

type User struct {
	Name, Email string
	Age         int
}

func (user User) welcomeUser(name string) {
	fmt.Println("Welcome", name)
}

func main() {
	fikri := User{"Nurul Fikri", "nf@gmail.com", 27}
	fmt.Println(fikri)

	// struct method can only be used on a variable with the struct variable
	fikri.welcomeUser(fikri.Name)
}
