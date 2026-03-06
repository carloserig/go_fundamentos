package main

import "fmt"

func main() {
	v1 := 10
	v2 := 7
	soma := v1 + v2
	fmt.Println("O resultado soma é: ", soma)

	subtracao := v1 - v2
	fmt.Println("O resultado subtracao é: ", subtracao)

	multiplicacao := v1 * v2
	fmt.Println("O resultado da multiplicacao é: ", multiplicacao)

	divisao := v1 / v2
	fmt.Println("O resultado da divisao é: ", divisao)

	resto := v1 % v2
	fmt.Println("O resto é: ", resto)

	// divisão inteiros converter para float
	resultadofloat := float64(v1) / float64(v2)
	fmt.Printf("Resultado: %.2f\n", resultadofloat)

}
