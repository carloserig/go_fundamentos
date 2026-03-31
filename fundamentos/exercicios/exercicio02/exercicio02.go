package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite a primeira nota: ")
	nota1Str, _ := reader.ReadString('\n')
	nota1, _ := strconv.ParseFloat(nota1Str, 64)

	fmt.Print("Digite a segunda nota: ")
	nota2Str, _ := reader.ReadString('\n')
	nota2, _ := strconv.ParseFloat(nota2Str, 64)

	media := (nota1 + nota2) / 2
	fmt.Println("media = ", media)
	// se a media for maior ou igual a 7, exibir "Aprovado", caso contrário, exibir "Reprovado".
	if media >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
	fmt.Println("A média é:", media)
}
