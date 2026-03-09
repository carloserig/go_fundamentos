package main

import "fmt"

func main() {
	fmt.Println("Operadores relacionais")
	/*
		==	igual
		!=	diferente
		>	maior
		<	menor
		>=	maior ou igual
		<=	menor ou igual
	*/
	nome1 := "carlos"
	nome2 := "eduardo"

	numero1 := 10
	numero2 := 7

	fmt.Println(nome1 == nome2)
	fmt.Println(numero1 >= numero2)
	fmt.Println(10 > 20)
	fmt.Println(10 < 20)
	fmt.Println(10 <= 20)
	fmt.Println(20 != 10)

	fmt.Println("matheus é maior que lucas?:", len(nome1) >= len(nome2))

}
