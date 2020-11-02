package main

import "fmt"

func main() {
	// fungsi len() menghitung elemen slice
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// var fruits = []string{"apple", "grape", "banana"}
	// fmt.Println(len(fruits))

	// fungsi cap() menghitung lebar atau kapasitas elemen
	// var aFruits = fruits[0:3]
	// fmt.Println(len(aFruits))
	// fmt.Println(cap(aFruits))

	// var bFruits = fruits[1:4]
	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	//fungsi append()
	// var cFruits = append(fruits, "papaya")
	// fmt.Println(fruits)
	// fmt.Println(cFruits)

	// var bFruits = fruits[0:2]
	// fmt.Println(cap(bFruits))
	// fmt.Println(len(bFruits))

	// fmt.Println(fruits)
	// fmt.Println(bFruits)

	// var cFruits = append(bFruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	// fmt.Println(cFruits)

	// fungsi copy()

	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)
	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// dst := []string{"potato", "potato", "potato"}
	// src := []string{"watermelon", "pinnaple"}
	// n := copy(dst, src)
	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	fruits := []string{"apple", "grape", "banana"}

	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))

}