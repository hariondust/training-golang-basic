package main

import (
	"fmt"
	"go-basic-learning/database"
	_ "go-basic-learning/internal" // pakai _ agar bisa di init tanpa perlu dipanggil
)

func main() {
	fmt.Println(database.GetDatabase())
}
