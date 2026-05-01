package aluno

import "fmt"

type Aluno struct {
	Nome string
	N1   float64
	N2   float64
}

// Cadastrar novo aluno
func Cadastro(a Aluno) {
	fmt.Printf("Aluno %s cadastrado com sucesso! \n", a.Nome)
}

func CalcularMedia(n1 float64, n2 float64) float64 {
	return (n1 + n2) / 2
}
