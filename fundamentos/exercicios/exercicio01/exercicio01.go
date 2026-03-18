package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu nome:")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)
	fmt.Print("Digite seu e-mail:")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	fmt.Print("Digite DDD:")
	ddd, _ := reader.ReadString('\n')
	ddd = strings.TrimSpace(ddd)
	fmt.Print("Digite seu telefone:")
	telefone, _ := reader.ReadString('\n')
	telefone = strings.TrimSpace(telefone)

	fmt.Println("\n - - - Dados Informados - - -")
	fmt.Println("Nome:", nome)
	fmt.Println("Email:", email)
	fmt.Println("Telefone: +55", ddd, telefone)
}
