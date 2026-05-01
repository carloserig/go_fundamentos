package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome        string
	Idade       int
	Peso        float64
	Altura      float64
	Salario     float64
	Dependentes int
}

var reader = bufio.NewReader(os.Stdin)

func lerString(msg string) string {
	for {
		fmt.Print(msg)
		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)

		if texto == "" {
			fmt.Println("Erro: campo não pode ficar vazio.")
			continue
		}
		return texto
	}
}

func lerInt(msg string) int {
	for {
		fmt.Print(msg)
		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)

		valor, err := strconv.Atoi(texto)
		if err != nil {
			fmt.Println("Erro: digite um número inteiro válido.")
			continue
		}
		return valor
	}
}

func lerFloat(msg string) float64 {
	for {
		fmt.Print(msg)
		texto, _ := reader.ReadString('\n')
		texto = strings.TrimSpace(texto)
		texto = strings.ReplaceAll(texto, ",", ".")

		valor, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Erro: digite um número decimal válido.")
			continue
		}
		return valor
	}
}

func cadastrarPessoa() Pessoa {
	fmt.Println("\n=== CADASTRO DE DADOS ===")
	nome := lerString("Nome: ")
	idade := lerInt("Idade: ")
	peso := lerFloat("Peso (kg): ")
	altura := lerFloat("Altura (m): ")
	salario := lerFloat("Salário (R$): ")
	dependentes := lerInt("Quantidade de dependentes: ")

	return Pessoa{
		Nome:        nome,
		Idade:       idade,
		Peso:        peso,
		Altura:      altura,
		Salario:     salario,
		Dependentes: dependentes,
	}
}

func mostrarDados(p Pessoa) {
	fmt.Println("\n=== DADOS CADASTRADOS ===")
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
	fmt.Printf("Peso: %.2f kg\n", p.Peso)
	fmt.Printf("Altura: %.2f m\n", p.Altura)
	fmt.Printf("Salário: R$ %.2f\n", p.Salario)
	fmt.Println("Dependentes:", p.Dependentes)
}

func calcularIMC(p Pessoa) {
	if p.Altura <= 0 {
		fmt.Println("Erro: altura inválida para cálculo de IMC.")
		return
	}

	imc := p.Peso / (p.Altura * p.Altura)

	fmt.Println("\n=== CÁLCULO DE IMC ===")
	fmt.Printf("IMC de %s: %.2f\n", p.Nome, imc)

	if imc < 18.5 {
		fmt.Println("Classificação: Abaixo do peso")
	} else if imc < 25 {
		fmt.Println("Classificação: Peso normal")
	} else if imc < 30 {
		fmt.Println("Classificação: Sobrepeso")
	} else if imc < 35 {
		fmt.Println("Classificação: Obesidade grau I")
	} else if imc < 40 {
		fmt.Println("Classificação: Obesidade grau II")
	} else {
		fmt.Println("Classificação: Obesidade grau III")
	}
}

func calcularImposto(p Pessoa) {
	var aliquota float64

	if p.Salario <= 2000 {
		aliquota = 0.0
	} else if p.Salario <= 3500 {
		aliquota = 0.075
	} else if p.Salario <= 5000 {
		aliquota = 0.15
	} else {
		aliquota = 0.225
	}

	imposto := p.Salario * aliquota
	salarioLiquido := p.Salario - imposto

	fmt.Println("\n=== CÁLCULO DE IMPOSTO ===")
	fmt.Printf("Salário bruto: R$ %.2f\n", p.Salario)
	fmt.Printf("Alíquota aplicada: %.1f%%\n", aliquota*100)
	fmt.Printf("Imposto estimado: R$ %.2f\n", imposto)
	fmt.Printf("Salário líquido estimado: R$ %.2f\n", salarioLiquido)
}

func calcularBonus(p Pessoa) {
	bonusBase := p.Salario * 0.10
	bonusDependentes := float64(p.Dependentes) * 50.0
	totalBonus := bonusBase + bonusDependentes

	fmt.Println("\n=== CÁLCULO DE BÔNUS ===")
	fmt.Printf("Bônus base (10%%): R$ %.2f\n", bonusBase)
	fmt.Printf("Bônus por dependentes: R$ %.2f\n", bonusDependentes)
	fmt.Printf("Bônus total: R$ %.2f\n", totalBonus)
	fmt.Printf("Salário com bônus: R$ %.2f\n", p.Salario+totalBonus)
}

func mostrarResumo(p Pessoa) {
	imc := 0.0
	if p.Altura > 0 {
		imc = p.Peso / (p.Altura * p.Altura)
	}

	var categoriaSalarial string
	if p.Salario < 2000 {
		categoriaSalarial = "Baixa renda"
	} else if p.Salario < 5000 {
		categoriaSalarial = "Média renda"
	} else {
		categoriaSalarial = "Alta renda"
	}

	dados := map[string]interface{}{
		"Nome":               p.Nome,
		"Idade":              p.Idade,
		"Peso":               p.Peso,
		"Altura":             p.Altura,
		"Salário":            p.Salario,
		"Dependentes":        p.Dependentes,
		"IMC":                fmt.Sprintf("%.2f", imc),
		"Categoria salarial": categoriaSalarial,
	}

	fmt.Println("\n=== RESUMO GERAL ===")
	for chave, valor := range dados {
		fmt.Printf("%s: %v\n", chave, valor)
	}
}

func menu() {
	fmt.Println("\n=========== MENU ===========")
	fmt.Println("1 - Cadastrar dados")
	fmt.Println("2 - Mostrar dados")
	fmt.Println("3 - Calcular IMC")
	fmt.Println("4 - Calcular imposto")
	fmt.Println("5 - Calcular bônus")
	fmt.Println("6 - Mostrar resumo geral")
	fmt.Println("0 - Sair")
	fmt.Println("============================")
}

func main() {
	var pessoa Pessoa
	var cadastrado bool

	for {
		menu()
		opcao := lerInt("Escolha uma opção: ")

		switch opcao {
		case 1:
			pessoa = cadastrarPessoa()
			cadastrado = true
			fmt.Println("\nCadastro realizado com sucesso!")

		case 2:
			if !cadastrado {
				fmt.Println("\nErro: cadastre os dados primeiro.")
				continue
			}
			mostrarDados(pessoa)

		case 3:
			if !cadastrado {
				fmt.Println("\nErro: cadastre os dados primeiro.")
				continue
			}
			calcularIMC(pessoa)

		case 4:
			if !cadastrado {
				fmt.Println("\nErro: cadastre os dados primeiro.")
				continue
			}
			calcularImposto(pessoa)

		case 5:
			if !cadastrado {
				fmt.Println("\nErro: cadastre os dados primeiro.")
				continue
			}
			calcularBonus(pessoa)

		case 6:
			if !cadastrado {
				fmt.Println("\nErro: cadastre os dados primeiro.")
				continue
			}
			mostrarResumo(pessoa)

		case 0:
			fmt.Println("\nPrograma encerrado.")
			return

		default:
			fmt.Println("\nOpção inválida. Tente novamente.")
		}
	}
}
