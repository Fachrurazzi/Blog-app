package main

import "fmt"
// import "strings"
// import (
// 	"fmt"
// 	"math"
// )

// type hitung2d interface {
// 	luas() float64
// 	keliling() float64
// }

// type hitung3d interface {
// 	volume() float64
// }

// type hitung interface {
// 	hitung2d
// 	hitung3d
// }

// type lingkaran struct {
// 	diameter float64
// }

// func (l lingkaran) jariJari() float64 {
// 	return l.diameter / 2
// }

// func (l lingkaran) luas() float64 {
// 	return math.Pi * math.Pow(l.jariJari(), 2)
// }

// func (l lingkaran) keliling() float64 {
// 	return math.Pi * l.diameter
// }

// type persegi struct {
// 	sisi float64
// }

// func (p persegi) luas() float64 {
// 	return math.Pow(p.sisi, 2)
// }

// func (p persegi) keliling() float64 {
// 	return p.sisi * 4
// }

// type kubus struct {
// 	sisi float64
// }

// func (k *kubus) volume() float64 {
// 	return math.Pow(k.sisi, 3)
// }

// func (k *kubus) luas() float64 {
// 	return math.Pow(k.sisi, 2) * 6
// }

// func (k *kubus) keliling() float64 {
// 	return k.sisi * 12
// }

// func main() {
// 	// var bangunDatar hitung

// 	// bangunDatar = persegi{10.0}
// 	// fmt.Println("==== Persegi")
// 	// fmt.Println("luas		:", bangunDatar.luas())
// 	// fmt.Println("keliling 	:", bangunDatar.keliling())

// 	// bangunDatar = lingkaran{14.0}
// 	// fmt.Println("==== Lingkaran")
// 	// fmt.Println("luas 		:", bangunDatar.luas())
// 	// fmt.Println("keliling 	:", bangunDatar.keliling())
// 	// fmt.Println("jari-jari 	:", bangunDatar.(lingkaran).jariJari())

// 	var bangunRuang hitung
// 	bangunRuang	= &kubus{4}
// 	fmt.Println("==== Kubus")
// 	fmt.Println("volume		:", bangunRuang.volume())
// 	fmt.Println("luas		:", bangunRuang.luas())
// 	fmt.Println("keliling	:", bangunRuang.keliling())

// }


// penggunaan interface kosong atau interface{}

// func main() {
	// var secret interface{}

	// secret = "razzy tirta"
	// fmt.Println(secret)

	// secret = []string{"papaya", "apple", "mango"}
	// fmt.Println(secret)

	// secret = 12.5
	// fmt.Println(secret)
	// var data map[string]interface{}

	// data = map[string]interface{}{
	// 	"name": "razzy tirta",
	// 	"grade": 2,
	// 	"breakfast": []string{"apple", "mango", "banana"},
	// }

	// fmt.Println(data["name"])
// }

// casting variabel interface kosong
// func main() {
// 	var secret interface{}

// 	secret = 2
// 	var number = secret.(int) * 10
// 	fmt.Println(secret, "multiplied by 10 is :", number)

// 	secret = []string{"apple", "mango", "lemon"}
// 	var fruits = strings.Join(secret.([]string), ", ")
// 	fmt.Println(fruits, "is my favorite fruits")

// }


// casting variabel interface kosong ke objek pointer
// type person struct {
// 	name string
// 	age int
// }

// func main() {
// 	var secret interface{} = &person{"wadow", 22}
// 	var name = secret.(*person).name
// 	fmt.Println(name)
// }

// Kombinasi slice, map dan interface {}
// func main() {
// 	var person = []map[string]interface{}{
// 	{"name": "razzi", "age": 22},
// 	{"name": "andi", "age": 22},
// 	{"name": "empus", "age": 23},
// 	}

// 	for _, each := range person {
// 		fmt.Println("nama :", each["name"],"\nUmur :", each["age"])
// 	}
// }

func main() {
	var fruits = []interface{}{
	map[string]interface{}{"name": "strawberry", "total" :20},
	[]string{"manggo", "pineapple", "papaya"},
	"orange",
	}

	for _,each := range fruits {
		fmt.Println(each)
	}
}