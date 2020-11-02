package main

import "fmt"

func main() {
	// variabel underscore
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "razzi", "tirta"
	fmt.Println(name)
	// variabel dengan keyword new
	nama := new(string)

	fmt.Println(nama) // nilai yang dihasilkan berupa heksadesimal
	fmt.Println(*nama) // kosong ""

	//variabel dengan keyword make
	// channel
	// slice
	// map
	
}