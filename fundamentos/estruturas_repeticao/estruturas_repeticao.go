package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	total := 10
	for i := 1; i <= total; i++ {
		fmt.Println(i, "de total de", total)
	}

	fmt.Println()

	numero := 1
	for numero <= 3 {
		fmt.Println("Oi, passando pela", numero, "vez")
		numero++
	}

	fmt.Println()

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
