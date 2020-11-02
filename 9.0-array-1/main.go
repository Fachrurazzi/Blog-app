package main

import "fmt"

func main() {
	var names[4] string
	names[0] = "Catur"
	names[1] ="Razzi"
	names[2] = "Dimas"
	names[3] = "Fachri"

	fmt.Println(names[0], names[1], names[2], names[3])

	// cara horizontal
	var fruits = [4]string{"banana", "apple", "melon", "mango"}

	// cara vertikal
	/*
	var fruits = [4]string{
		"banana",
		"apple",
		"melon",
		"mango"
	}
	*/

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// inisialisasi array tanpa jumlah elemen
	var numbers = [...]int{5,6,3,7,6,1,37}

	fmt.Println("data array \t", numbers)
	fmt.Println("jumlah array \t", len(numbers))
}