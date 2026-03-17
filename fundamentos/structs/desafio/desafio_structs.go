package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	var pessoas []Pessoa
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("- - - Cadastro de Pessoas - - -")
	for {

		fmt.Print("Digite o nome: ")
		nome, _ := reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		var idade int
		for {
			fmt.Print("Digite a idade: ")
			idadeStr, _ := reader.ReadString('\n')
			idadeStr = strings.TrimSpace(idadeStr)

			var err error

			idade, err = strconv.Atoi(idadeStr)

			if err == nil {
				break
			}

			fmt.Println("Digite apenas números!")
		}

		p := Pessoa{
			Nome:  nome,
			Idade: idade,
		}

		pessoas = append(pessoas, p)

		fmt.Print("Deseja continuar? (s/n): ")
		cont, _ := reader.ReadString('\n')
		cont = strings.TrimSpace(cont)

		if cont == "n" {
			break
		}
	}

	fmt.Println("\nLista de pessoas:")

	for i, p := range pessoas {
		fmt.Printf("%d - Nome: %s | Idade: %d\n", i+1, p.Nome, p.Idade)
	}
}
