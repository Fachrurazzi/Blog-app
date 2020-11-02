package main

import "fmt"

func main() {
	// operator aritmatika + - * / %
	var nilai = (((2 + 6) % 3) * 4 - 2) / 3

	fmt.Println("hasil:", nilai)

	// operator perbandingan true dan false == != > >= < <=
	var value  = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai dari %d (%t) \n", value, isEqual)

	// operator logika && || !
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t)\n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t)\n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left \t\t(%t)\n", leftReverse)


}