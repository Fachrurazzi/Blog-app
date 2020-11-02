package main

import (
	"fmt"
	"runtime"
)

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var messages = make(chan string)

// 	var sayHelloTo = func (who string) {
// 		var data = fmt.Sprintf("hello %s", who)
// 		messages <- data
// 	}

// 	go sayHelloTo("razzi tirta")
// 	go sayHelloTo("catur ridho")
// 	go sayHelloTo("dimas aditya")

// 	var messages1 = <- messages
// 	fmt.Println(messages1)

// 	var messages2 = <- messages
// 	fmt.Println(messages2)

// 	var messages3 = <- messages
// 	fmt.Println(messages3)
// }


//chanel sebagai parameter tipe data
func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"razzi", "dimas", "catur"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(what chan string)  {
	fmt.Println(<-what)
}