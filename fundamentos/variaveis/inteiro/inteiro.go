package main

import "fmt"

func main() {
	//int / int8 / int16 / int32 / int64
	// 8 bits → 2ˆ8 = 256 valores possíveis 128 negativos 128 positivos (-128 a 127)
	fmt.Println("Variáveis do tipo inteiro(+,-)")
	fmt.Println("A soma é", 1+2)
	var num int8 = 127
	fmt.Println(num)

	//uint / uint8 / uint16 / uint32 / uint64
	fmt.Println("Variáveis do tipo inteiro acima de 0")
	var numPositivo uint8 = 100
	fmt.Println(numPositivo)

	//Incremento/decremento
	fmt.Println("Somando 1")
	numPositivo++
	fmt.Println(numPositivo)

	fmt.Println("Subtraindoo 1")
	numPositivo--
	fmt.Println(numPositivo)

}
