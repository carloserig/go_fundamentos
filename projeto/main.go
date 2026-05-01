package main

import (
	"fmt"
	"projeto/aluno"
	"projeto/utils"
)

// go mod init projeto

func main() {
	fmt.Println("Dentro do main...")
	utils.Mensagem("Carlos")
	status := utils.StatusSistema()
	fmt.Println("Status do sistema:", status)

	resultado := utils.Somar(100, 50)
	fmt.Printf("O resultado da soma é %.2f\n", resultado)

	a := aluno.Aluno{Nome: "Carlos Erig", N1: 7.0, N2: 1.0}
	aluno.Cadastrar(a)

	res := aluno.CalcularMedia(a.N1, a.N2)
	fmt.Println("Aluno:", a.N2)
	fmt.Printf("A média é %.2f\n", res)

	if res >= 7.0 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

}
