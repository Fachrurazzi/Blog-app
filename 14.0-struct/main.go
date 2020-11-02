package main

import "fmt"

// func main() {
// 	// var st1 student
// 	var st1 = student{}

// 	st1.name = "Razzi"
// 	st1.grade = 2

// 	var st2 = student{"ethan", 2}

// 	var st3 = student{name : "bambang"}

// 	fmt.Println("name :", st1.name)
// 	fmt.Println("grade :", st1.grade)

// 	fmt.Println("student 1 :", st1.name)
// 	fmt.Println("student 2 :", st2.name)
// 	fmt.Println("student 3 :", st3.name)
// }

// type student struct {
// 	name string
// 	grade int
// }

// variabel objek pointer
// func main()  {
// 	var s1 = student{name: "razzi", grade: 2}
// 	var s2 *student = &s1
// 	fmt.Println("student 1, name :", s1.name)
// 	fmt.Println("student 4, name :", s2.name)

// 	s2.name = "tirta"
// 	fmt.Println("student 1, name :", s1.name)
// 	fmt.Println("student 4, name :", s2.name)
// }

// type student struct{
// 	name string
// 	grade int
// }

//embedded struct
// type person struct {
// 	name string
// 	age int
// }

// type student struct {
// 	grade int
// 	person
// }
// func main() {
// 	var s1 = student{}
// 	s1.name = "razzi"
// 	s1.age = 21
// 	s1.grade = 2

// 	fmt.Println("name :", s1.name)
// 	fmt.Println("age :", s1.age)
// 	fmt.Println("age :", s1.person.age)
// 	fmt.Println("grade :", s1.grade)
// }

//embedded struct dengan nama property yang sama

// type person struct {
// 	name string
// 	age int
// }

// type student struct {
// 	person
// 	age int
// 	grade int
// }

func main() {
	// var s1 student

	// s1.name = "razzi"
	// s1.age = 21 // age of student
	// s1.person.age = 22 // age of person

	// fmt.Println("name :", s1.name)
	// fmt.Println("age :", s1.age)
	// fmt.Println("age :", s1.person.age)

	//pengisian nilai sub-struct
	// var p1 = person{name: "razzi", age: 21}
	// var s2 = student{person: p1, grade: 2}

	// fmt.Println("name: ", s2.name)
	// fmt.Println("age : ", s2.person.age)
	// fmt.Println("grade : ", s2.grade)

	// Anonymous struct
	// var s1 = struct {
	// 	person
	// 	grade int
	// }{}
	// 	s1.person = person{name: "Razzi", age: 22}
	// 	s1.grade = 2

	// 	fmt.Println("name :", s1.person.name)
	// 	fmt.Println("age :", s1.person.age)
	// 	fmt.Println("grade :", s1.grade)

	// kombinasi slice dan struct

	// var allStudents = []person{
	// 	{name: "razzi", age: 22},
	// 	{name: "tirta", age: 23},
	// 	{name: "bambang", age: 22},
	// }

	// for _, students := range allStudents{
	// 	fmt.Println(students.name, "age is", students.age)
	// }

	// inisialisasi slice anonymous struct

	// var allStudents = []struct {
	// 	person
	// 	grade int
	// }{
	// 	{person: person{"razzy", 21}, grade: 1},
	// 	{person: person{"tirta", 20}, grade: 2},
	// 	{person: person{"andi", 19}, grade: 3},
	// }

	// for _, students := range allStudents {
	// 	fmt.Println(students)
	// }

	// deklarasi anonymous struct menggunakan keyword var
	// 	var student struct {
	// person
	// grade int
	// }
	// student.person = person{"wick", 21}
	// student.grade = 2
	// hanya deklarasi
	// var student struct {
	// grade int
	// }
	// // deklarasi sekaligus inisialisasi
	// var student = struct {
	// grade int
	// } {
	// 12,
	// }

	//nested struct
	// type student struct {
	// 	type person struct {
	// 		name string
	// 		age int
	// 	}
	// 	grade int
	// 	hobbies []string
	// }

	// type alias
	type Person struct {
		name string
		age int
	}

	type People = Person

	var p1 = Person{name: "razzi", age: 22}
	fmt.Println(p1)

	var p2 = People{name: "tirta", age: 21}
	fmt.Println(p2)

}