package main

import "fmt"

func main() {
	// var getMinMax = func(n []int) (int, int) {
	// 	var min, max int
	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	// 	return min, max
	// }

	// var numbers = []int{2, 3, 4, 3, 4, 2, 6}
	// var min, max = getMinMax(numbers)
	// fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	//Immediately-Invoked Function Expression(IIFE)
	// var numbers = []int{2,3,2,4,5,3,2,1,0,5,2}

	// var newNumbers = func(min int) []int {
	// 	var r []int
	// 	for _, e  := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 		fmt.Print(e , "<", min," = ")
	// 		fmt.Println(r)
	// 	}
	// 	return r
	// }(3)

	// fmt.Println("original number :", numbers)
	// fmt.Println("filtered number :", newNumbers)

	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}
// closure deng nilai kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
			fmt.Println(e, " <= ", max, "=", res)
		}
	}

	var getNumbers = func() []int {
		return res
	}
	return len(res), getNumbers
}

