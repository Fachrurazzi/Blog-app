package main

import (
	"fmt"
	"errors"
	"strings"
)

// import (
// 	"fmt"
// 	"strconv"
// )

// error

// func main() {
// 	var input string
// 	fmt.Print("Type some number : ")
// 	fmt.Scanln(&input)

// 	var number int
// 	var err error
// 	number, err = strconv.Atoi(input)

// 	if err == nil {
// 		fmt.Println(number, "is number")
// 	} else {
// 		fmt.Println(input, "is not number")
// 		fmt.Println(err.Error())
// 	}
// }

//custom error

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot empty")
	}

	return true, nil
}

// func main()  {
// 	var name string
// 	fmt.Print("Type your name :")
// 	fmt.Scanln(&name)

// 	if valid, err := validate(name); valid {
// 		fmt.Println("hallo", name)
// 	} else {
// 		fmt.Println(err.Error())
// 	}
// }

//panic()
// func main()  {
// 	var name string
// 	fmt.Print("Type your name :")
// 	fmt.Scanln(&name)

// 	if valid, err := validate(name); valid {
// 		fmt.Println("hallo", name)
// 	} else {
// 		panic(err.Error())
// 		fmt.Println("end")
// 	}
// }

//recover
func catch()  {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("hello", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}