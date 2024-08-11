package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Fikri")
	fmt.Println(result)

	fmt.Println(getHello("Idea"))
}
