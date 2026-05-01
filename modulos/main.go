package main

// go mod init modules
// executar no terminar na raiz

import (
	"fmt"
	"modules/aluno"
	"modules/utils"
)

func main() {
	fmt.Println("Iniciando o main ...")
	utils.Mensagem("Carlos")
	status := utils.StatusSistema()
	fmt.Println("O sistema está", status)

	a1 := aluno.Aluno{Nome: "Carlos Erig", N1: 10, N2: 6.5}
	aluno.Cadastro(a1)
	fmt.Printf("Média: %.2f\n", aluno.CalcularMedia(a1.N1, a1.N2))

}
