package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	// if want to assign function to variable, dont use ()
	goodbye := getGoodBye
	fmt.Println(goodbye("Fikri"))
}
