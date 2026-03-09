package main

import "fmt"

func main() {
	nota1 := 3.0
	nota2 := 4.0

	media := (nota1 + nota2) / 2

	if media >= 7 {
		fmt.Println("Aprovado!")
	} else if media >= 4 && media < 7 {
		fmt.Println("Recuperação!")
	} else {
		fmt.Println("Reprovado!")
	}
}
