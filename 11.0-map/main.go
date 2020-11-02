package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40
	chicken["mei"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// cara vertikal
	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara horizontal
	var chicken1 = map[string]int{
		"januari" : 50,
		"februari" : 40,
	}

	fmt.Println(chicken1)

	// variabel map bisa diinisialisasikan tanpa nilai awal dengan 3 cara :
	// var chicken2 = map[string]int{}
	// var chicken3 = make(map[string]int)
	// var chicken4 = *new(map[string]int)

	// iterasi item map menggunakan for - range
	var chicken2 = map[string]int{
		"januari": 50,
		"februari": 40,
		"maret": 34,
		"april": 67,
	}

	for key, val := range chicken2 {
		fmt.Println(key, "\t:", val)
	}

	// menghapus item map

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	// deteksi keberadaan item dengan key tertentu
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	var datum  =  []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "jl. mangga", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chicken := range chickens{
		fmt.Println(chicken["gender"], chicken["name"])
	}

	for _, data := range datum{
		fmt.Println(data)
	}



}