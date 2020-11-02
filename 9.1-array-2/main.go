package main

import "fmt"

func main() {
	// array multi dimensi
	var numbers1 = [2][3]int{[3]int{4,3,6}, [3]int{9,1,5}}
	var numbers2 = [2][3]int{{2,5,1}, {2,1,6 }}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruits = [4]string{"apple", "orange", "banana", "grape"}

	i := 0
	for  i < len(fruits) {
		fmt.Println("elemen", i, ":", fruits[i])
		i++
	}

		fmt.Println()
	// array dengan keyword for - range
	var animals = [...]string{"chicken", "goat", "cow", "camel"}

	for i, animal := range animals {
		fmt.Printf("elemen %d %s\n", i, animal)
	}

	fmt.Println()
	// penggunaan variabel underscore pada for - range ketika kita ingin menampilkan hanya elementnya saja

	var pemainBola = [...]string{"Messi", "Neymar", "Ronaldo", "Mbappe"}

	for _, pemain := range pemainBola {
		fmt.Printf("Nama Pemain: %s\n", pemain)
	}

	fmt.Println()

	var bahasa = make([]string, 2)
	bahasa[0] = "Golang"
	bahasa[1] = "Kotlin"

	fmt.Println(bahasa)
}