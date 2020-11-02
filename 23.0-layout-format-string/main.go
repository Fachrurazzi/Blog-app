package main

import "fmt"

type student struct {
	name string
	height float64
	age int32
	isGraduated bool
	hobbies []string
}

func main() {
	var data = student {
		name:	"Razzi",
		height:	166.7,
		age:	22,
		isGraduated: true,
		hobbies: []string{"futsal", "badminton", "coding"},
	}

	//layout format %b binary 11010
	fmt.Printf("%b\n", data.age)

	//layout format %c unicode 'n'
	fmt.Printf("%c\n", 1400)

	//layout format %d data numerik
	fmt.Printf("%d\n", data.age)

	//layout format %e atau %E data numerik desimal ke bentuk notasi numerik standar
	fmt.Printf("%e\n", data.height)

	fmt.Printf("%E\n", data.height)

	//layout format %f atau %F desimal
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	//layout format %g atau %G
	fmt.Printf("%g", 0.12121212323222)
	fmt.Printf("%.5g", 0.12121212323222)

	//layout format %o oktal
	fmt.Printf("%o\n", data.age)

	//layout format %p pointer
	fmt.Printf("%p\n", &data.name)

	//layout format %q escape string
	fmt.Printf("%q\n", `" name \ height "`)

	//layout format %s
	fmt.Printf("%s\n", data.name)

	//layout format %t
	fmt.Printf("%t", data.isGraduated)

	//layout format %T
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)

	//layout format %v
	fmt.Printf("%v\n", data)

	//layout format %+v
	fmt.Printf("%+v\n", data)

	//layout format %#v
	fmt.Printf("%#v\n", data)

	//layout format %x
	fmt.Printf("%x\n", data.age)

	//layout format %%
	fmt.Printf("%%\n")
}

