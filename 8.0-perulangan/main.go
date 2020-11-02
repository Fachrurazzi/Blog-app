package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("")
	// for dengan argumen keyword hanya kondisi
	j := 0

	for j < 5 {
		fmt.Println("Angka", j)
		j++
	}

	fmt.Println("")

	//for tanpa argumen

	k := 0

	for {
		fmt.Println("Angka", k)
		k++

		if(k == 5) {
			break
		}
	}

	fmt.Println("")
	// break && coutinue
	for l := 1; l < 10; l++ {
		if l % 2 == 1 {
			continue
		}

		if l > 8 {
			break
		}

		fmt.Println("Angka", l)
	}

	// for bersarang
	/*
		m = 0; 0 < 5; 0+0 = 0
		m = 1; 1 < 5; 0+1 = 1
		m = 2; 2 < 5; 1+1 = 2
		m = 3; 1 < 5; 2+1 = 3
		m = 4; 1 < 5; 3+1 = 4
		m = 5; 5 < 5; 4+1 = 5
	*/
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}

	fmt.Println()
	outerLoop:
	for p := 0; p < 5; p++ {
		for q := p; q < 5; q++ {
			if p == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", p, "][", q,"]", "\n")
		}
	}
}