/**
Package adalah tempat u/ mengorganizir kode program yg dibuat di Go-Lang
Dengan package, kita bisa mengelompokan kode program agar lebih rapih
Package sebenernya hanya direktori folder di sistem operasi kita
*/

package helper

// penulisan fungsi pascal case agar bisa di access package yg lain
func SayHello(name string) string {
	return "Hello " + name
}

// penulisan fungsi camel case agar tidak bisa di access package lain
// tujuannya yaitu untuk memodif fungsi ini di package tersebut
func checkVersion(version string) string {
	return "This is version " + version
}

// selain nama function, access modifier juga berlaku untuk variable
var version = "1.0.0"      // tidak bisa diakses dari luar package
var Application = "golang" // bisa diakses dari luar package
