package main

import "fmt"

func main()  {
	// penggunaan konstanta
	const firstName string ="Razzy"
	const lastName = "Tirta"

	fmt.Print("Hallo ", firstName, "!\n")
	fmt.Print("nice to meet you ", lastName, "!\n")

	// penggunaan fungsi fmt.print()
	fmt.Println("razzy tirta")
	fmt.Println("razzy", "tirta")

	fmt.Print("razzy tirta\n")
	fmt.Print("razzy", "tirta\n")
	fmt.Print("razzy", " ", "tirta\n")
}