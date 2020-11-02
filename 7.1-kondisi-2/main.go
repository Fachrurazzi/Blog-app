package main

import "fmt"

func main() {
	//switch case
	point := 6

	switch point {
	case 8:
		fmt.Println("Perfect!")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// case untuk banyak kondisi
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//switch case gaya if else
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// penggunaan keyword fallthrough dalam switch
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//if dan switch case bersarang
	if point > 7 {
		switch {
		case point == 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}