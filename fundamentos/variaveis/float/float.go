package main

import "fmt"

func main() {
	//float / float32 / float64
	fmt.Println("Variáveis do tipo float")
	var numero = 40.2045689
	numero2 := 50.23242
	fmt.Printf("Soma de 2 número float: %.2f\n", numero+numero2)
	fmt.Print("Soma realizada com sucesso.")
	/*
		%  → indica que será inserido um valor
		.2 → mostra 2 casas decimais
		f  → formata como número de ponto flutuante (float)
		\n → Quebra de linha (newline).
	*/

}
