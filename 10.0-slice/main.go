package main

import "fmt"

func main() {
	var fruits = []string{"banana", "watermelon", "orange", "strawberry"}
	// fmt.Println(fruits[0])
	// var newFruits = fruits[0:2]
	// var newFruits = fruits[0:4]
	// var newFruits = fruits[0:0]
	// var newFruits = fruits[4:4]
	// var newFruits = fruits[:]
	// var newFruits = fruits[2:]
	// var newFruits = fruits[:2]
	// fmt.Println(newFruits)

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = fruits[1:2]
	var baFruits = fruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	baFruits[0] = "grape"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
}