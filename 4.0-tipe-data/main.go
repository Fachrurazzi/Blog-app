package main

import "fmt"

func main() {
	// tipe data non desimal
	var positiveNumber uint8  = 78
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif %d\n", positiveNumber)
	fmt.Printf("bilangan negatif %d\n", negativeNumber)

	// tipe data desimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal %f\n", decimalNumber)
	fmt.Printf("bilangan desimal %.3f\n", decimalNumber)

	// tipe data boolean
	var exist bool = true

	fmt.Printf("exist ? %t \n", exist)

	// tipe data string
	var message string ="Hallo"
	fmt.Printf("message: %s\n", message)

	var pesan = `Nama saya adalah "Razzi".
	Salam kenal.
	Mari belajar "Golang"`
	fmt.Printf(pesan)
}