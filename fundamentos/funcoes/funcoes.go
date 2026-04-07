package main

import (
	"errors"
	"fmt"
)

//	func boasVindas() {
//		nome := "Pedro"
//		fmt.Println("Olá, seja bem vindo", nome)
//	}
func boasVindas(nome string) {
	fmt.Println("Olá, seja bem vindo(a)", nome)
}

func soma(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("O resultado da soma é", resultado)
}

func somaComRetorno(n1 int, n2 int) int {
	return n1 + n2
}

type Usuario struct {
	Nome  string
	Senha string
}

func autenticar(user Usuario, senha string) (string, error) {
	if user.Senha != senha {
		return "", errors.New("Senha inválida!")
	}
	return user.Nome, nil
}

// função variádica
func somarNumeros(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Iniciando estudo de funções")
	boasVindas("Maria")
	soma(25, 80)
	soma(10, 80)
	res := somaComRetorno(2, 10)
	fmt.Println("O retorno da soma é", res)

	//criar um usuário ref. a struct
	usuarioSalvoNoBanco := Usuario{Nome: "carlos", Senha: "123456"}
	nome, err := autenticar(usuarioSalvoNoBanco, "123456")
	if err != nil {
		fmt.Println("Erro de autenticação =>", err)
	} else {
		fmt.Println("Usuário autenticado =>", nome)
	}

	nums := []int{1, 56, 100, 300, 20, 22, 15, 50}
	total := somarNumeros(nums...)
	fmt.Println("Total ===> ", total)
}
