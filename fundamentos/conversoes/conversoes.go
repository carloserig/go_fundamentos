package main

import "fmt"

func main() {
	capital := 2000
	taxa := 0.03
	periodo := 12

	var jurosfinal = float64(capital) * taxa * float64(periodo)
	// obs: Se for float não precisa converter

	fmt.Printf("O juro final é de %.2f\n", jurosfinal)

}
