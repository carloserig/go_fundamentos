package aluno

import "fmt"

type Aluno struct {
	Nome string
	N1   float64
	N2   float64
}

// Mensagem de Cadastro do aluno com sucesso
func Cadastrar(a Aluno) {
	fmt.Printf("Aluno %s cadastrado com sucesso!", a.Nome)
}

// Calcula a Média do aluno
func CalcularMedia(n1 float64, n2 float64) float64 {
	return (n1 + n2) / 2
}
