package main

import "fmt"

func main() {
	fmt.Println("Operadores lógicos")
	/*
		&& → E
		|| → OU
		!  → NÃO
	*/
	estoque := true
	vendaLiberada := false
	freteGratis := true

	fmt.Println("E:", estoque && vendaLiberada)
	fmt.Println("OU:", estoque || vendaLiberada)
	fmt.Println("NÃO:", !freteGratis)

}
