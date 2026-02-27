package main

import "fmt"

func main() {
	//Variáveis
	//string
	//int / int8 / int16 / int32 / int64
	// 8 bits → 2ˆ8 = 256 valores possíveis 128 negativos 128 positivos (-128 a 127)
	//uint / uint8 / uint16 / uint32 / uint64
	//float / float32 / float64
	//bool
	//var num int16 = 32767
	//fmt.Println(num)
	fmt.Println("Variáveis")
	fmt.Println("Variáveis do tipo string")
	var nome = "Carlos"
	//var nome string= "Carlos"
	//var nome
	// nome = "Carlos"
	// forma reduzida => precisa inicializar
	//nome := "Carlos"
	var sobrenome = "Erig"
	fmt.Println("Variáveis do tipo float")
	var numero = 40.2045689
	fmt.Println(nome + " " + sobrenome)
	fmt.Printf("Número: %.2f\n", numero)
	nome = "Teste"
	fmt.Println(nome)
}
