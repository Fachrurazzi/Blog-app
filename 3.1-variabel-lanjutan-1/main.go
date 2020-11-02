package main

import "fmt"

func main() {
	// deklarasi multi variabel

	// cara pertama
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// cara kedua
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// cara ketiga
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// dengan teknik type intefence kita bisa memasukkan nilai yang tipe datanya berbeda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Printf("%s %s %s\n", first, second, third)
	fmt.Printf("%s %s %s\n", fourth, fifth, sixth)
	fmt.Printf("%s %s %s\n", seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
}