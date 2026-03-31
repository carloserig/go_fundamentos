package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// continue => volta logo início for | break => termina na hora
	reader := bufio.NewReader(os.Stdin)

	base := 0.0
	for {
		fmt.Print("Digite a base do triângulo: ")
		baseStr, _ := reader.ReadString('\n')
		baseStr = strings.TrimSpace(baseStr)

		if !strings.Contains(baseStr, ".") {
			fmt.Println("Erro: digite um número decimal (ex: 5.0, 3.25)")
			continue
		}

		valor, err := strconv.ParseFloat(baseStr, 64)
		if err != nil {
			fmt.Println("Erro: valor inválido")
			continue
		}

		base = valor
		break
	}

	var altura float64
	for {
		fmt.Print("Digite a altura do triângulo: ")
		alturaStr, _ := reader.ReadString('\n')
		alturaStr = strings.TrimSpace(alturaStr)

		if !strings.Contains(alturaStr, ".") {
			fmt.Println("Erro: digite um número decimal (ex: 5.0, 3.25)")
			continue
		}

		valor, err := strconv.ParseFloat(alturaStr, 64)
		if err != nil {
			fmt.Println("Erro: valor inválido")
			continue
		}

		altura = valor
		break
	}

	area := (base * altura) / 2

	fmt.Println("\n - - - Resultado - - -")
	fmt.Println("Base:", base)
	fmt.Println("Altura:", altura)
	fmt.Println("Área do triângulo:", area)
}
