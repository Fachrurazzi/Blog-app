package main

import (
	"fmt"
)

// func main() {
// 	// defer fmt.Println("halo")
// 	// fmt.Println("Selamat Datang")
// 	orderSomeFood("pizza")
// 	orderSomeFood("burger")
// }

// func orderSomeFood(menu string)  {
// 	defer fmt.Println("Terima Kasih, silahkan tunggu")

// 	if menu == "pizza" {
// 		fmt.Println("Pilihan tepat!", " ")
// 		fmt.Println("Pizza ditempat kami paling enak!", "\n")
// 		return
// 	}

// 	fmt.Println("Pesanan anda:", menu)
// }

// func main() {
// 	number := 3

// 	if number == 3 {
// 		fmt.Println("halo 1")
// 		defer fmt.Println("halo 3")
// 	}
// 	fmt.Println("halo 2")
// }

func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		func() {
			fmt.Println("halo 3")
		}()

		fmt.Println("halo 2")
	}
}