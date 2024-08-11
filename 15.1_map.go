package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Fikri"
	// person["address"] = "Bekasi"

	person := map[string]string{
		"name":    "Fikri",
		"address": "Bekasi",
	}

	fmt.Println(person["name"])
	fmt.Println(person)

	fmt.Println(person["salah"]) // testing value yg salah
}
