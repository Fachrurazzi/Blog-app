package main

import "fmt"

func main() {
	// 1. variabel manifest typing
	var firstName string = "razzi"

	var lastName string
	lastName = "tirta"

	// fungsi fmt.Printf
	fmt.Printf("hallo %s %s!\n", firstName, lastName)
	fmt.Printf("hallo razzi tirta!\n")
	fmt.Println("hallo",firstName, lastName + "!")

	// 2. variabel type inference
	var namaDepan string = "Andi"
	namaBelakang := "Setya"

	fmt.Printf("hello %s %s!\n", namaDepan, namaBelakang)

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	// var firstName = "Bambang"
	// tanpa var, tanpa tipe data, menggunakan perantara :=
	// lastName := "Setia"

	// penggunaan ":=" hanya digunakan diawal pada saat deklarasi ketika assigment nilai selanjutnya harus menggunakan "="
	// animal := "Kucing"
	// animal = "Buaya"
	// animal ="Gajah"

}