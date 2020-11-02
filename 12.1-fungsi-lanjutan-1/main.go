package main

import "fmt"
import "strings"

// fungsi variadic
// func main() {
// 	var avg = calculate(2,3,4,5,6,2,4,6,3,5)
// 	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
// 	fmt.Println(msg)
// }

// func calculate(numbers ...int) float64 {
// 	var total int = 0
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))

// 	return avg
// }

// func main() {
// 	var numbers = []int{2,3,4,5,6,4,2,4,5,6,78}
// 	var avg = calculate(numbers ...)
// 	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

// 	fmt.Println(msg)
// }

// func calculate(numbers ...int) float64 {
// 	var total = 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	var avg = float64(total) / float64(len(numbers))

// 	return avg
// }

// fungsi dengan parameter biasa dan variadic
// func main() {
// 	var hobbies = []string{"football", "reading", "badminton"}
// 	var name = "Razzi"
// 	yourHobbies(name, hobbies...)
// }

// func yourHobbies(name string, hobbies ...string) {
// 	var hobbiesString = strings.Join(hobbies, ", ")

// 	fmt.Printf("Hello, my name is : %s \n", name)
// 	fmt.Printf("My hobbies are : %s \n", hobbiesString)
// }

func main() {
	var phones = []string{"xiaomi", "flagship", "mid-range"}
	var name = "Xiaomi Redmi 9A"
	handphone(name, phones...)
}

func handphone(name string, categories ...string) {
	var phones = strings.Join(categories, ", ")

	fmt.Printf("Phone name : %s \n", name)
	fmt.Printf("Categories : %s", phones)
}
