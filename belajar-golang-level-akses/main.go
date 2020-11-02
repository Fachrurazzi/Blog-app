package main

import (
	. "belajar-golang-level-akses/library"
	f "fmt"
)

func main() {
	// library.SayHello("razzy")
	// var s1 = Student{"ethan", 22}
	// f.Println("name", s1.Name)
	// f.Println("grade", s1.Grade)

	// sayHello("razzi")
	f.Printf("Name :%s\n", Student.Name)
	f.Printf("Grade : %d\n", Student.Grade)
}