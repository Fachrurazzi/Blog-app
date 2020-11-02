package main

import (
	"fmt"
	// "strings"
)

// func main() {
// 	var s1 = student{"razzy tirta", 21}
// 	s1.sayHello()

// 	var name = s1.getNameAt(1)
// 	fmt.Println("nama panggilan", name)
// }

// type student struct {
// 	name string
// 	grade int
// }

// func (s student) sayHello() {
// 	fmt.Println("hallo", s.name)
// }

// func (s student) getNameAt(i int) string {
// 	return strings.Split(s.name, " ")[i-1]
// }


// method pointer
type student struct {
	name string
	grade int
}

func main() {
	var s1 = student{"razzy tirta", 22}
	fmt.Println("before :", s1.name)

	s1.changeName1("bambang sutri")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("catur ridho")
	fmt.Println("s1 after changeName2", s1.name)
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}